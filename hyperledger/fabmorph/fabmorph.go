//////////////////////////////////////
//    THE PAILLIER EXPERIMENT     ////
//////////////////////////////////////
/* 
	This is the fabmorph, a chaincode that runs in Hyperledger Fabric 1.4.
	He was created as part of the Paillier Experiment. You can invoke its methods
	to send an encrypted energy measurement record to the ledger, or to retrieve such 
	information. It can also implement accumalative sums with the measurements using
	Paillier cryptosystem. Please check the code to find more information about how
	use it.

	@author: Wilson S. Melo Jr.
	@date: Apr/2019
*/
package main

import (
	//the majority of the imports are trivial...
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"crypto/rand"
	"math/big"
	"errors"
	"io"
	"time"
	"encoding/json"

	//these imports are for Hyperledger Fabric interface
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

////////////////////////////////////////
//////// GOMORPH INSERTION CODE ////////
////////////////////////////////////////

/* All the following functions belong to the Gomorph Library. We insert all the code into the chaincode
main package because it was easier than make Hyperledger Fabric to accept external packages.
Gomorph is available at https://github.com/radicalrafi/gomorph.

	Paillier cryptosystem implementation

	Provides primitives for Public & Private Key Generation /  Encryption / Decryption
	Provides Functions to operate on the Cyphertext according to Paillier algorithm

	@author: radicalrafi
	@license: Apache 2.0

*/

/* The Paillier crypto system picks two keys p & q and denotes n = p*q
Messages have to be in the ring Z/nZ (integers modulo n)
Therefore a Message can't be bigger than n
*/
var ErrLongMessage = errors.New("Gaillier Error #1: Message is too long for The Public-Key Size \n Message should be smaller than Key size you choose")

//constants
var one = big.NewInt(1)

//Key structs
type PubKey struct {
	KeyLen int
	N      *big.Int //n = p*q (where p & q are two primes)
	G      *big.Int //g random integer in Z\*\n^2
	Nsq    *big.Int //N^2
}

type PrivKey struct {
	KeyLen int
	PubKey
	L *big.Int //lcm((p-1)*(q-1))
	U *big.Int //L^-1 modulo n mu = U = (L(g^L mod N^2)^-1)
}

func GenerateKeyPair(random io.Reader, bits int) (*PubKey, *PrivKey, error) {

	p, err := rand.Prime(random, bits/2)

	if err != nil {
		return nil, nil, err
	}

	q, err := rand.Prime(random, bits/2)

	if err != nil {
		return nil, nil, err
	}

	//N = p*q

	n := new(big.Int).Mul(p, q)

	nSq := new(big.Int).Mul(n, n)

	g := new(big.Int).Add(n, one)

	//p-1
	pMin := new(big.Int).Sub(p, one)
	//q-1
	qMin := new(big.Int).Sub(q, one)
	//(p-1)*(q-1)
	l := new(big.Int).Mul(pMin, qMin)
	//l^-1 mod n
	u := new(big.Int).ModInverse(l, n)
	pub := &PubKey{KeyLen: bits, N: n, Nsq: nSq, G: g}
	return pub, &PrivKey{PubKey: *pub, KeyLen: bits, L: l, U: u}, nil
}

/*
	Encrypt :function to encrypt the message into a paillier cipher text
	using the following rule :
	cipher = g^m * r^n mod n^2
	* r is random integer such as 0 <= r <= n
	* m is the message
*/
func Encrypt(pubkey *PubKey, message []byte) ([]byte, error) {

	r, err := rand.Prime(rand.Reader, pubkey.KeyLen)
	if err != nil {
		return nil, err
	}

	m := new(big.Int).SetBytes(message)
	if pubkey.N.Cmp(m) < 1 {
		return nil, ErrLongMessage
	}
	//c = g^m * r^nmod n^2

	//g^m
	gm := new(big.Int).Exp(pubkey.G, m, pubkey.Nsq)
	//r^n
	rn := new(big.Int).Exp(r, pubkey.N, pubkey.Nsq)
	//prod = g^m * r^n
	prod := new(big.Int).Mul(gm, rn)

	c := new(big.Int).Mod(prod, pubkey.Nsq)

	return c.Bytes(), nil
}

/*
	Decrypts a given ciphertext following the rule:
	m = L(c^lambda mod n^2).mu mod n
	* lambda : L
	* mu : U

*/
func Decrypt(privkey *PrivKey, cipher []byte) ([]byte, error) {

	c := new(big.Int).SetBytes(cipher)

	if privkey.Nsq.Cmp(c) < 1 {
		return nil, ErrLongMessage
	}

	//c^l mod n^2
	a := new(big.Int).Exp(c, privkey.L, privkey.Nsq)

	//L(x) = x-1 / n we compute L(a)
	l := new(big.Int).Div(new(big.Int).Sub(a, one), privkey.N)

	//computing m
	m := new(big.Int).Mod(new(big.Int).Mul(l, privkey.U), privkey.N)

	return m.Bytes(), nil

}

/*
	Homomorphic Properties of Paillier Cryptosystem

	* The product of two ciphers decrypts to the sum of the plain text
	* The product of a cipher with a non-cipher raising g will decrypt to their sum
	* A Cipher raised to a non-cipher decrypts to their product
	* Any cipher raised to an integer k will decrypt to the product of the deciphered and k
*/

//Add two ciphers together
func Add(pubkey *PubKey, c1, c2 []byte) []byte {

	a := new(big.Int).SetBytes(c1)
	b := new(big.Int).SetBytes(c2)

	// a * b mod n^²
	res := new(big.Int).Mod(new(big.Int).Mul(a, b), pubkey.Nsq)

	return res.Bytes()
}

//Add a constant & a cipher
func AddConstant(pubkey *PubKey, cipher, constant []byte) []byte {

	c := new(big.Int).SetBytes(cipher)
	k := new(big.Int).SetBytes(constant)

	//result = c * g^k mod n^2
	res := new(big.Int).Mod(
		new(big.Int).Mul(c, new(big.Int).Exp(pubkey.G, k, pubkey.Nsq)), pubkey.Nsq)

	return res.Bytes()

}

//Multiplication by a constant integer
func Mul(pubkey *PubKey, cipher, constant []byte) []byte {

	c := new(big.Int).SetBytes(cipher)
	k := new(big.Int).SetBytes(constant)

	//res = c^k mod n^2
	res := new(big.Int).Exp(c, k, pubkey.Nsq)

	return res.Bytes()
}

////////////////////////////////////////
//////// GOMORPH CODE ENDS HERE ////////
////////////////////////////////////////


//////////////////////////////////////
//////// FABMORPH - CHAINCODE ////////
//////////////////////////////////////

/* All the following functions are used to implemetn fabmorph chaincode. This chaincode
basically works with 2 scenarios:
	1) A measurement instrument sends measures in plain text to the blockchain. The smartcontract
	applies homomorphic cryptography and stores encrypted measures. The client can recover it by
	reading the encrypted value and decrypting locally.
	2) A measurement instrument sends encrypted measures. The smartcontract applies homomorphic 
	operations to obtain the (encrypted) total consumption value and to store it in the ledger.
	The client can it by reading the encrypted value and decrypting locally.

	@author: Wilson S. Melo Jr.
	@date: Apr/2019

*/

/*
 * This is an auxiliar function to convert the Pallier public key struct into a string.
 * It is necessary to lead with the object in Fabric transactions.
 * The key in string format will have the 4 fields in the following order: KeyLen, N, G, Nsq 
 */
func PubKey2String(pubkey *PubKey) (string){
	return strconv.Itoa(pubkey.KeyLen) + "," + pubkey.N.String() + "," + pubkey.G.String() + "," + pubkey.Nsq.String()
}

/*
 * This is an auxiliar function to convert a string key into the Pallier public key struct.
 * It is necessary to lead with the object in Fabric transactions.
 * The key in string format MUST be formated with 4 comma-separate substrings, that correspond to the
 * respective fields: KeyLen, N, G, Nsq.
 * Example: 
 * "512,9317708529424897702020382930116702407110920461631622943398410342176567591986029379418525445868266
	642035599843718252107477949966673217711651921453503121487,9317708529424897702020382930116702407110920
	46163162294339841034217656759198602937941852544586826664203559984371825210747794996667321771165192145
	3503121488,868196922393174897253161831767710200925181500218099202121385823388772389324857444497325510
	95045029334111747236626667783012791667259632343428969631355566392144138916566553732073336450909522598
	17833785123113998363351341374266805275045098714823061322751114829098157258427086799436922320373510493
	4252452681091169"
 * which means:
	KeyLen = 512
	N = 9317708529424897702020382930116702407110920461631622943398410342176567591986029379418525445868266
	642035599843718252107477949966673217711651921453503121487
	G = 9317708529424897702020382930116702407110920461631622943398410342176567591986029379418525445868266
	642035599843718252107477949966673217711651921453503121488
	Nsq = 86819692239317489725316183176771020092518150021809920212138582338877238932485744449732551095045
	02933411174723662666778301279166725963234342896963135556639214413891656655373207333645090952259817833
	78512311399836335134137426680527504509871482306132275111482909815725842708679943692232037351049342524
	52681091169
 */
func String2PubKey(pubkey string) (*PubKey){
	//extract tokens from pubkey string to a vector
	strvec := strings.Split(pubkey, ",")

	if len(strvec) != 4 {
		return nil
	}

	//get the key lenght
	l, err := strconv.Atoi(strvec[0]);

	n, res := new(big.Int).SetString(strvec[1],10)
	g, res := new(big.Int).SetString(strvec[2],10)
	nSq, res := new(big.Int).SetString(strvec[3],10)

	if err == nil && res {
		//compose public key using values informed in the string pubkey
		return &PubKey{KeyLen: l, N: n, G: g, Nsq: nSq}
	} else {
		return nil
	}
}

/*
 * This is an auxiliar function that encapsulates de conversion of a vetor of bytes
 * representing a big.Int number into a literal transcription in string format.
 */
func Bytes2String(bcipher []byte) (string){
	//creates an auxiliar big.Int to manipulate the string
	cipherInt := new(big.Int).SetBytes(bcipher)
	return cipherInt.String()
}

/*
 * This is an auxiliar function that encapsulates de conversion of a string with a
 * literal representation of a big.Int number (base 10) into a vector of bytes in the
 * format required by Gomorph library. 
 */
func String2Bytes(cipher string) ([]byte){
	//creates an auxiliar big.Int to manipulate the string
	cipherInt, err := new(big.Int).SetString(cipher,10)

	if err {
		return cipherInt.Bytes()
	} else {
		return nil
	}
}

/*
 * SmartContract defines the chaincode base structure. All the methods are implemented to
 * return a SmartContrac type.
 */
type SmartContract struct {
}

/*
 Meter constitutes our key|value struct (digital asset) and implements a single 
 record to manage the 
 meter public key and measures. All blockchain transactions operates with this type.
 IMPORTANT: all the field names must start with upper case
 */
type Meter struct {
 	PublicKey		string `json:"publickey"`
	PlainMeasure 	int64  `json:"plainmeasure"`
	EncrypMeasure 	string `json:"encrypmeasure"`
}

/*
 * The Init method is called when the Smart Contract "fabmorph" is instantiated.
 * Best practice is to have any Ledger initialization in separate function.
 * Note that chaincode upgrade also calls this function to reset
 * or to migrate data, so be careful to avoid a scenario where you
 * inadvertently clobber your ledger's data!
 */
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The invoke function is called on each transaction invoking the chaincode. It
 * follows a structure of switching calls, so each valid feature need to
 * have a proper entry-point.
 */
func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	// extract the function name and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	//implements a switch for each acceptable function
	if fn == "registerMeter" {
		//registers a new meter into the ledger
		return s.registerMeter(stub, args)

	} else if fn == "insertMeasurement" {
		//inserts a measurement which increases the meter consumption counter. The measurement
		return s.insertMeasurement(stub, args)
	
	} else if fn == "insertPlainTextMeasurement" {
		//inserts a measurement which increases the meter consumption counter. The measurement
		return s.insertPlainTextMeasurement(stub, args)
	
	} else if fn == "getConsumption" {
		//retrieves the accumulated consumption 
		return s.getConsumption(stub, args)

	} else if fn == "sleepTest" {
		//retrieves the accumulated consumption 
		return s.sleepTest(stub, args)

	} else if fn == "countHistory" {
		//look for a specific fill up record and brings its changing history
		return s.countHistory(stub, args)

	} else if fn == "countLedger" {
		//look for a specific fill up record and brings its changing history
		return s.countLedger(stub)

	} else if fn == "queryLedger" {
		//execute a CouchDB query, args must include query expression
		return s.queryLedger(stub, args)
	}

	//function fn not implemented, notify error
	return shim.Error("Chaincode do not support this function")
}

/*
	SmartContract::registerMeter(...)
   	Does the register of a new meter into the ledger. The meter is the base of the key|value structure.
   	The key constitutes the meter ID. If a cryptographic public key is informed, it indicates that
   	the meter accumulates its measurements using homomorphic cryptography.
   	The vector args[] must contain two parameters:
	- args[0] - meter ID
	- args[1] - the Paillier public key associated with the meter
	- args[2] - (optional) initial encrypted measurement value.
 */
func (s *SmartContract) registerMeter(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//validate args vector lenght
	if !(len(args) == 2 || len(args) == 3) {
		return shim.Error("It was expected the parameters: <meter id> <public key> [encrypted inital consumption]")
	}

	//gets the parameters associated with the meter ID and the public key (in string format)
	meterid := args[0]
	strpubkey := args[1]

	//create a var to store the encrypted initial consumption
	encrypMeasureString := ""

	//tests if the encrypted inital consumption was provided
	if len(args) == 3 {
		//initiates the value of the encrypted measurement
		encrypMeasureString = args[2]
	}

	fmt.Println("Value of encrypted measurement: (", encrypMeasureString, ")")

	//defines the initial consumption as zero
	plainMeasure := new(big.Int).SetInt64(0)

	//test if strpubkey is a empty string
	if len(strpubkey) > 0 {
		//there is a public key, so we need to encrypt the measurement
		pubkey := String2PubKey(strpubkey)

		if pubkey != nil && encrypMeasureString == "" {
			//encrypts the initial consumption using initial value in plainMeasure
			encrypMeasure, err := Encrypt(pubkey, plainMeasure.Bytes())

			//test if encryptation had success
			if err != nil {
				return shim.Error("Error on encrypting with the informed public key")
			}			
			encrypMeasureString = Bytes2String(encrypMeasure)

			fmt.Println("I encrypted the measurement: (", encrypMeasureString, ")")
		}
	}

	//creates the meter record
	var meter = Meter{PublicKey: strpubkey, PlainMeasure: plainMeasure.Int64(), EncrypMeasure: encrypMeasureString}

	//encapsulates meter in a JSON structure
	meterAsBytes, _ := json.Marshal(meter)

	//registers meter in the ledger
	stub.PutState(meterid, meterAsBytes)

	//loging...
	fmt.Println("Registering meter: ", meter)

	//notify procedure success
	return shim.Success(nil)
}

/* 
	This method implements the insertion of encrypted measurements in the blockchain.
	The encryptation must uses the same public key configured to the meter.
	Notice that the informed measurement will be added (accumulated) to the the previous
	encrypted measurement consumption information.
	The vector args[] must contain two parameters:
	- args[0] - meter ID
	- args[1] - the encrypted measurement, in a string representing a big int number.
*/
func (s *SmartContract) insertMeasurement(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//validate args vector lenght
	if len(args) != 2 {
		return shim.Error("It was expected 2 parameter: <meter ID> <measurement>")
	}

	//gets the parameter associated with the meter ID and the incremental measurement
	meterid := args[0]

	//try to convert the informed measurement into the format []byte, required by Gomorph
	measurement := String2Bytes(args[1])

	//check if we have success 
	if measurement == nil {
		//measurement is not a proper number
		return shim.Error("Error on veryfing measurement, it is not a proper input")
	}
	
	//retrive meter record
	meterAsBytes, err := stub.GetState(meterid)

	//test if we receive a valid meter ID
	if err != nil || meterAsBytes == nil {
		return shim.Error("Error on retrieving meter ID register")
	}

	//creates Meter struct to manipulate returned bytes
	MyMeter := Meter{}

	//convert bytes into a Meter object
	json.Unmarshal(meterAsBytes, &MyMeter)

	//convert meter public key to the format expected by Gomorph
	pubkey := String2PubKey(MyMeter.PublicKey)

	//compute the new measurement value by using homomorphic adding constant property
	newEncrypMeasure := Add(pubkey, String2Bytes(MyMeter.EncrypMeasure), measurement)  

	//update encrypted measure cipher
	MyMeter.EncrypMeasure = Bytes2String(newEncrypMeasure);

	//encapsulates meter back into the JSON structure
	newMeterAsBytes, _ := json.Marshal(MyMeter)

	//update meter state in the ledger
	stub.PutState(meterid, newMeterAsBytes)

	//loging...
	fmt.Println("Updating meter consumption:", MyMeter)

	//notify procedure success
	return shim.Success(nil)
}

/* 
	This method enables the insertion of measurements as plaintext information.
	Its main intent is to provide a comparing in terms of performance with the 
	insertMeasurement() method.
	The vector args[] must contain two parameters:
	- args[0] - meter ID
	- args[1] - the measurement value (as integer, float type is not supported)
*/
func (s *SmartContract) insertPlainTextMeasurement(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//validate args vector lenght
	if len(args) != 2 {
		return shim.Error("It was expected 2 parameter: <meter ID> <measurement>")
	}

	//gets the parameter associated with the meter ID and the incremental measurement
	meterid := args[0]

	//try to convert the informed measurement into the format []byte, required by Gomorph
	measurement, err := strconv.ParseInt(args[1], 10, 64)

	//check if we have success 
	if err != nil {
		//measurement is not a proper number
		return shim.Error("Error on veryfing measurement, it is not a proper int64 input")
	}
	
	//retrive meter record
	meterAsBytes, err := stub.GetState(meterid)

	//test if we receive a valid meter ID
	if err != nil || meterAsBytes == nil {
		return shim.Error("Error on retrieving meter ID register")
	}

	//creates Meter struct to manipulate returned bytes
	MyMeter := Meter{}

	//convert bytes into a Meter object
	json.Unmarshal(meterAsBytes, &MyMeter)

	//update the plaintext measurement
	MyMeter.PlainMeasure += measurement;

	//encapsulates meter back into the JSON structure
	newMeterAsBytes, _ := json.Marshal(MyMeter)

	//update meter state in the ledger
	stub.PutState(meterid, newMeterAsBytes)

	//loging...
	fmt.Println("Updating meter plaintext consumption:", MyMeter)

	//notify procedure success
	return shim.Success(nil)
}

/* 
	This method enables retrieves the consumption information from the ledger.
	It returns a vector of bytes with the meter asset in JSON format.
	The vector args[] must contain two parameters:
	- args[0] - meter ID
*/
func (s *SmartContract) getConsumption(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	//validate args vector lenght
	if len(args) != 1 {
		return shim.Error("It was expected 1 parameter: <meter ID>")
	}

	//gets the parameter associated with the meter ID and the incremental measurement
	meterid := args[0]

	//retrive meter record
	meterAsBytes, err := stub.GetState(meterid)

	//test if we receive a valid meter ID
	if err != nil || meterAsBytes == nil {
		return shim.Error("Error on retrieving meter ID register")
	}

	//return payload with bytes related to the meter state
	return shim.Success(meterAsBytes)
}

/* 
	This method is a dummy test that makes the endorser "sleep" for some seconds.
	It is usefull to check either the sleeptime affects the performance of concurrent
	transactions.
	- args[0] - sleeptime (in seconds)
*/
func (s *SmartContract) sleepTest(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	//validate args vector lenght
	if len(args) != 1 {
		return shim.Error("It was expected 1 parameter: <sleeptime>")
	}

	//gets the parameter associated with the meter ID and the incremental measurement
	sleeptime, err := strconv.Atoi(args[0])

	//test if we receive a valid meter ID
	if err != nil {
		return shim.Error("Error on retrieving sleep time")
	}

	//tests if sleeptime is a valid value
	if sleeptime > 0 {
		//stops during sleeptime seconds
		time.Sleep(time.Duration(sleeptime) * time.Second)
	}

	//return payload with bytes related to the meter state
	return shim.Success(nil)
}

/*
   This method brings the changing history of a specific meter asset. It can be useful to
   query all the changes that happened with a meter value.
   - args[0] - asset key (or meter ID)
*/
func (s *SmartContract) queryHistory(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//validate args vector lenght
	if len(args) != 1 {
		return shim.Error("It was expected 1 parameter: <key>")
	}

	historyIer, err := stub.GetHistoryForKey(args[0])

	//verifies if the history exists
	if err != nil {
		//fmt.Println(errMsg)
		return shim.Error("Fail on getting ledger history")
	}

	// buffer is a JSON array containing records
	var buffer bytes.Buffer
	var counter = 0
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for historyIer.HasNext() {
		//increments iterator
		queryResponse, err := historyIer.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		//generates a formated result
		buffer.WriteString("{\"Value\":")
		buffer.WriteString("\"")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("\"")
		buffer.WriteString(", \"Counter\":")
		buffer.WriteString(strconv.Itoa(counter))
		//buffer.WriteString(queryResponse.Timestamp)
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true

		//increases counter
		counter++
	}
	buffer.WriteString("]")
	historyIer.Close()

	//loging...
	fmt.Printf("Consulting ledger history, found %d\n records", counter)

	//notify procedure success
	return shim.Success(buffer.Bytes())
}

/*
   This method brings the number of times that a meter asset was modified in the ledger.
   It performs faster than queryHistory() method once it does not retrive any information,
   it only counts the changes.
   - args[0] - asset key (or meter ID)
*/
func (s *SmartContract) countHistory(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//validate args vector lenght
	if len(args) != 1 {
		return shim.Error("It was expected 1 parameter: <key>")
	}

	historyIer, err := stub.GetHistoryForKey(args[0])

	//verifies if the history exists
	if err != nil {
		//fmt.Println(errMsg)
		return shim.Error("Fail on getting ledger history")
	}

	//creates a counter
	var counter int64
	counter = 0

	for historyIer.HasNext() {
		//increments iterator
		_, err := historyIer.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//increases counter
		counter++

		fmt.Printf("Consulting ledger history, found %d\n records", counter)
	}
	// buffer is a JSON array containing records
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString("\"Counter\":")
	buffer.WriteString(strconv.FormatInt(counter, 10))
	buffer.WriteString("]")

	historyIer.Close()

	//loging...
	fmt.Printf("Consulting ledger history, found %d\n records", counter)

	//notify procedure success
	return shim.Success(buffer.Bytes())
}

/*
   This method counts the total of well succeeded transactions in the ledger.
*/
func (s *SmartContract) countLedger(stub shim.ChaincodeStubInterface) sc.Response {

	//use a range of keys, assuming that the max key value is 999999,
	resultsIterator, err := stub.GetStateByRange("0", "999999")
	if err != nil {
		return shim.Error(err.Error())
	}

	//defer iterator closes at the end of the function
	defer resultsIterator.Close()

	//creates a counter
	var counter int64
	var keys int64
	counter = 0
	keys = 0

	//the interator checks all the valid keys
	for resultsIterator.HasNext() {

		//increments iterator
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//busca historico da proxima key
		historyIer, err := stub.GetHistoryForKey(queryResponse.Key)

		//verifies if the history exists
		if err != nil {
			//fmt.Println(errMsg)
			return shim.Error(err.Error())
		}

		defer historyIer.Close()

		for historyIer.HasNext() {
			//increments iterator
			_, err := historyIer.Next()
			if err != nil {
				return shim.Error(err.Error())
			}

			//increases counter
			counter++
		}
		fmt.Printf("Consulting ledger history, found key %d\n", queryResponse.Key)

		keys++
	}
	// buffer is a JSON array containing records
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString("\"Counter\":")
	buffer.WriteString(strconv.FormatInt(counter, 10))
	buffer.WriteString("\"Keys\":")
	buffer.WriteString(strconv.FormatInt(keys, 10))
	buffer.WriteString("]")

	//loging...
	fmt.Printf("Consulting ledger history, found %d transactions in %d keys\n", counter, keys)

	//notify procedure success
	return shim.Success(buffer.Bytes())
}

/*
   This method executes a free query on the ledger, returning a vector of meter assets.
   The query string must be a query expression supported by CouchDB servers.
   - args[0] - query string.
*/
func (s *SmartContract) queryLedger(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//validate args vector lenght
	if len(args) != 1 {
		return shim.Error("It was expected 1 parameter: <query string>")
	}

	//using auxiliar variable
	queryString := args[0]

	//loging...
	fmt.Printf("Executing the following query: %s\n", queryString)

	//try to execute query and obtain records iterator
	resultsIterator, err := stub.GetQueryResult(queryString)
	//test if iterator is valid
	if err != nil {
		return shim.Error(err.Error())
	}
	//defer iterator closes at the end of the function
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		//increments iterator
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		//generates a formated result
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	//loging...
	fmt.Printf("Obtained the following fill up records: %s\n", buffer.String())

	//notify procedure success
	return shim.Success(buffer.Bytes())
}

/*
   This method enables simple tests with the Paillier cryptosystem. Make yourserf
   confortable to modify it as much as necessary. 
*/
func TestPaillier() {
	//set two big.Int number that will be the operands for homomorphic computing
	case1 := new(big.Int).SetInt64(667) //you can assign any number
	case2 := new(big.Int).SetInt64(971) //you can assign any number

	//generate a Paillier key pair (remember of informing the correct size, e.g., 2048)
	pub, priv, err := GenerateKeyPair(rand.Reader, 2048)

	//assert the method execution
	if err != nil {
		fmt.Printf("Error Generating Keypair")
	} else {
		fmt.Printf("Key pair generate with success...\n")
	}

	//encrypt the operands
	fmt.Printf("Starting encryption:\n")
	encCase1, err1 := Encrypt(pub, case1.Bytes())
	encCase2, err2 := Encrypt(pub, case2.Bytes())

	//assert the method execution
	if err1 != nil || err2 != nil {
		fmt.Printf("Error Encrypting Integers")
	} else {
		fmt.Printf("%v -> %v\n", case1, encCase1)
		fmt.Printf("%v -> %v\n", case2, encCase2)
	}

	//test Paillier addition
	fmt.Printf("Testing the addition...\n")
	res := Add(pub, encCase1, encCase2)

	//log result
	fmt.Printf("We got: %v\n",res)

	//set the expected sum value
	corr := case1.Add(case1,case2)

	//decripts the sum result
	fmt.Printf("Testing decryption...\n")
	decRes, err := Decrypt(priv, res)

	//assert the result
	if err != nil {
		fmt.Printf("Failed to Decrypt Result got %v want %v with Error : %v", decRes, corr, err)
	}

	//decryption succeeds, check the result
	resB := new(big.Int).SetBytes(decRes)

	//assert the result
	if resB.Cmp(corr) != 0 {
		fmt.Printf("Failed to Add two ciphers got %v want %v", resB, corr)
	} else {
		fmt.Printf("Decryption is correct, all the tests performed with success!\n")
	}
}

/*
 * The main function starts up the chaincode in the container during instantiate
 */
func main() {

	////////////////////////////////////////////////////////
	// USE THIS BLOCK TO COMPILE THE CHAINCODE 
	if err := shim.Start(new(SmartContract)); err != nil {
	    fmt.Printf("Error starting SmartContract chaincode: %s\n", err)
	}
	////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////
	// USE THIS BLOCK TO PERFORM ANY TEST WITH PAILLIER
	//TestPaillier()
	////////////////////////////////////////////////////////

}
