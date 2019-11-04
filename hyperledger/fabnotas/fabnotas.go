/////////////////////////////////////////////////
//    CONTROLE SEGURO DE NOTAS DE ALUNOS     ////
/////////////////////////////////////////////////
/* 
	Devido a desconfiança dos alunos de que suas notas em avaliacoes estavam
	sendo modificadas indevidamente, foi criada uma aplicacao de blockchain
	para manter o registro das avaliacoe seguro, com a replicacao do mesmo
	entre todos os alunos e professores.

	O chaincode fabnotas encapsula essa aplicacao e será mantido pelos alunos
	da disciplina de Criptografia Aplicada.

	@author: Wilson S. Melo Jr.
	@date: Sep/2019
*/
package main

import (
	//Esses "imports" são biobliotecas diversas do Golang que precisamos para fazer 
	//nosso chaincode funcionar
	"bytes"
	"fmt"
	"strconv"
	"encoding/json"

	//Esses imports já são específicos do Hyperledger Fabric
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

/*
 SmartContract define a estrutura base do chaincode. Todos os metodos
 que escreveremos depois retornam um objeto SmartContract para pela
 interface shim.
 */
type SmartContract struct {
}

/*
 Avaliacao constitui nosso "digital asset", ou seja, ela é a estrutura tipo
 key/value que sera armazenada no blockchain. 
 */
type Avaliacao struct {
 	Questao			string 	`json:"questao"`
	Resposta 		string  `json:"resposta"`
	Nota 			int 	`json:"nota"`
}

/*
 O metodo Init() é chamado sempre que o chaincode eh instanciado em um channel.
 Aqui nao estamos usando ele pra nada em especial, mas voce podera customiza-lo
 pra inserir valores iniciais no ledger, por exemplo.
 */
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 O metodo Invoke() funciona como um broker para as funcoes especificas que o chaincode
 implementa. Ele basicamente mantem uma estrutura if... else ... que redireciona
 a execucao do chaincode de acordo com a funcao desejada.
 */
func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	//extrai o nome da funcao e os argumentos associados a uma transacao
	fn, args := stub.GetFunctionAndParameters()

	//implementa o tratamento para cada funcao reconhecida pelo chaincode
	if fn == "inserirAvaliacao" {
		//permite a um professor cadastrar uma avaliacao especifica
		return s.inserirAvaliacao(stub, args)

	} else if fn == "responderAvaliacao" {
		//permite ao aluno responder a avaliacao
		return s.responderAvaliacao(stub, args)

	} else if fn == "colarAvaliacao" {
		//permite ao aluno responder a avaliacao (colando)
		return s.colarAvaliacao(stub, args)
		
	} else if fn == "corrigirAvaliacao" {
		//permite ao professor corrigir e atribuir nota a avaliacao
		return s.corrigirAvaliacao(stub, args)
	
	} else if fn == "countLedger" {
		//conta o numero de registros no ledger
		return s.countLedger(stub)

	} else if fn == "queryLedger" {
		//executa uma query livre no CouchDB
		return s.queryLedger(stub, args)
	}

	//function fn not implemented, notify error
	return shim.Error("Ce tah de brincadera??? Essa funcao nao existe no chaincode!")
}

/*
	SmartContract::inserirAvaliacao(...)
	Um professor usa essa funcao para registrar uma avaliacao no blockchain. O primeiro argumento informado
	deve ser o ID da avaliacao, determinado de tal forma a identificar também o aluno.
	Alem do ID, o professor deve informar também a questão da avaliacao.

	O vetor args[] deve conter dois parametros:
	- args[0] - ID da avaliacao
	- args[1] - A questao de avaliacao
 */
func (s *SmartContract) inserirAvaliacao(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//valida o tamanho do vetor args
	if len(args) != 2 {
		return shim.Error("Eram esperados dois parametros: <avaliacao id> <questao>")
	}

	//obtem os parametros informados
	avalid := args[0]
	questao := args[1]

	//cria um registro de avaliacao (não respondida, nota igual a zero)
	var avaliacao = Avaliacao{Questao: questao, Resposta: "", Nota: 0}

	//encapsula o tipo avaliacao em um vetor de bytes (marshall)
	avaliacaoAsBytes, _ := json.Marshal(avaliacao)

	//insere a avaliacao no ledger
	stub.PutState(avalid, avaliacaoAsBytes)

	//Gera registro de log...
	fmt.Println("Inserindo avaliacao: ", avaliacao)

	//informa que fez tudo com sucesso
	return shim.Success(nil)
}

/* 
 Este metodo eh usado pelo aluno para responder a uma avaliacao. Ele procura por um
 ID de avaliacao especifico e permite modificar o valor do campo Resposta.

 O vetor args[] deve conter dois parametros:
 - args[0] - ID da avaliacao
 - args[1] - A resposta da avaliacao
*/
func (s *SmartContract) responderAvaliacao(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//valida o tamanho do vetor args
	if len(args) != 2 {
		return shim.Error("Eram esperados dois parametros: <avaliacao id> <questao>")
	}

	//obtem os parametros informados
	avalid := args[0]
	resposta := args[1]
	
	//tenta recuperar o registro da avaliacao no ledger
	avaliacaoAsBytes, err := stub.GetState(avalid)

	//testa se encontrou a avaliacao
	if err != nil || avaliacaoAsBytes == nil {
		return shim.Error("Avaliacao nao encontrada no ledger.")
	}

	//cria uma struct propria pra manipular a avaliacao
	avaliacao := Avaliacao{}

	//converte os bytes lidos do ledger em um struct avaliacao
	json.Unmarshal(avaliacaoAsBytes, &avaliacao)

	//preenche a resposta da avaliacao
	avaliacao.Resposta = resposta

	//encapsula o tipo avaliacao novamente no vetor de bytes (marshall)
	avaliacaoAsBytes, _ = json.Marshal(avaliacao)

	//modifica a avaliacao no ledger (adicionando a resposta)
	stub.PutState(avalid, avaliacaoAsBytes)

	//Gera registro de log...
	fmt.Println("A avaliacao foi respondida: ", avaliacao)

	//informa que fez tudo com sucesso
	return shim.Success(nil)
}

/* 
 Esta funcao do chaincode eh para ensinar os alunos como colar
 usando o blockchain.

 O vetor args[] deve conter dois parametros:
 - args[0] - ID da avaliacao que eu quero responder
 - args[1] - Palavra chave pra buscar uma resposta
*/
func (s *SmartContract) colarAvaliacao(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//valida o tamanho do vetor args
	if len(args) != 2 {
		return shim.Error("Eram esperados dois parametros: <avaliacao id> <palavra chave>")
	}

	//obtem os parametros informados
	avalid := args[0]
	palavrachave := args[1]
	
	//monta uma string contendo a query formatada em JSON
	minhaquery := "{\"selector\":{\"questao\":{\"$regex\":\"" + palavrachave + "\"}}}"

	//executa uma query procurando pela palavra chave
	meuiterator, err := stub.GetQueryResult(minhaquery)

	//testa se a busca no ledger retornou algo valido
	if err != nil {
		//o retorno foi invalido, gera mensagem de erro e finaliza
		return shim.Error(err.Error())
	}
	//defer iterator closes at the end of the function
	defer meuiterator.Close()

	//variaveis auxiliares para descobrir a maior nota
	//maiorID := "" //por que????
	maiorNota := 0 //nenhuma nota de valor inteiro positio eh maior que zero
	melhorResposta := ""

	//manipular questoes obtidas na "cola"
	for meuiterator.HasNext() {	
		//pega o registro corrente da minha lista
		registro, err := meuiterator.Next()

		//testa se a busca no ledger retornou algo valido
		if err != nil {
			//o retorno foi invalido, gera mensagem de erro e finaliza
			return shim.Error(err.Error())
		}

		//faz uma copia para outra variavel e resolver problema de cast
		minhacola := Avaliacao{}
		json.Unmarshal(registro.Value, &minhacola)

		//testo se a nota do registro corrente eh maior
		if minhacola.Nota > maiorNota {
			//encontrei uma nota maior ainda!
			maiorNota = minhacola.Nota

			//se essa eh a melhor nota, deve ser tambem a melhor resposta
			melhorResposta = minhacola.Resposta

			//guardo o ID da nota
			//maiorID := registro.Key
		} 
	}

	//tenta recuperar o registro da avaliacao no ledger
	avaliacaoAsBytes, err := stub.GetState(avalid)

	//testa se encontrou a avaliacao
	if err != nil || avaliacaoAsBytes == nil {
		return shim.Error("Avaliacao nao encontrada no ledger.")
	}

	//cria uma struct propria pra manipular a avaliacao
	avaliacao := Avaliacao{}

	//converte os bytes lidos do ledger em um struct avaliacao
	json.Unmarshal(avaliacaoAsBytes, &avaliacao)

	//preenche a resposta da avaliacao
	avaliacao.Resposta = melhorResposta

	//encapsula o tipo avaliacao novamente no vetor de bytes (marshall)
	avaliacaoAsBytes, _ = json.Marshal(avaliacao)

	//modifica a avaliacao no ledger (adicionando a resposta)
	stub.PutState(avalid, avaliacaoAsBytes)

	//Gera registro de log...
	fmt.Println("A avaliacao foi respondida (com a cola :-D ", avaliacao)

	//informa que fez tudo com sucesso
	return shim.Success(nil)
}

/* 
 Este metodo eh usado pelo professor (ou por um aluno hacker) para corrigir e atribuir
 uma nota a avaliacao. Ele procura por um ID de avaliacao especifico e permite modificar 
 o valor do campo Nota.

 O vetor args[] deve conter dois parametros:
 - args[0] - ID da avaliacao
 - args[1] - A nota da avaliacao
*/
func (s *SmartContract) corrigirAvaliacao(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//valida o tamanho do vetor args
	if len(args) != 2 {
		return shim.Error("Eram esperados dois parametros: <avaliacao id> <nota>")
	}

	//obtem os parametros informados
	avalid := args[0]
	nota, err := strconv.Atoi(args[1])

	//testa se o valor da nota é menor ou igual a 10
	if err!= nil || nota > 10 {
		return shim.Error("Nota nao eh numero ou eh maior do que 10. Informe um valor correto.")
	}
	
	//tenta recuperar o registro da avaliacao no ledger
	avaliacaoAsBytes, err := stub.GetState(avalid)

	//testa se encontrou a avaliacao
	if err != nil || avaliacaoAsBytes == nil {
		return shim.Error("Avaliacao nao encontrada no ledger.")
	}

	//cria uma struct propria pra manipular a avaliacao
	avaliacao := Avaliacao{}

	//converte os bytes lidos do ledger em um struct avaliacao
	json.Unmarshal(avaliacaoAsBytes, &avaliacao)

	//preenche a nota da avaliacao
	avaliacao.Nota = nota

	//encapsula o tipo avaliacao novamente no vetor de bytes (marshall)
	avaliacaoAsBytes, _ = json.Marshal(avaliacao)

	//modifica a avaliacao no ledger (adicionando a nota)
	stub.PutState(avalid, avaliacaoAsBytes)

	//Gera registro de log...
	fmt.Println("A avaliacao agora tem nota: ", avaliacao)

	//informa que fez tudo com sucesso
	return shim.Success(nil)
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
 This method brings the changing history of a specific data asset. It can be useful to
 query all the changes that happened with a id value.
 - args[0] - asset key (or ID)
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
   This method brings the number of times that a digital asset was modified in the ledger.
   It performs faster than queryHistory() method once it does not retrive any information,
   it only counts the changes.
   - args[0] - asset key (or ID)
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
 * The main function starts up the chaincode in the container during instantiate
 */
func main() {

	/////////////////////////////////////////////////////////////
	// Use esse bloco para compilar seu codigo como um chaincode
	if err := shim.Start(new(SmartContract)); err != nil {
	    fmt.Printf("Error starting SmartContract chaincode: %s\n", err)
	}
	/////////////////////////////////////////////////////////////

	/////////////////////////////////////////////////////////////
	// Use esse bloco para realizar quaisquer testes com o Go
	// ...
	/////////////////////////////////////////////////////////////

}
