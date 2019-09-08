*Data de aula: 9 e 11/9*

*Prof. Wilson ([wsjunior@inmetro.gov.br](mailto:wsjunior@inmetro.gov.br))*

# Hyperledger Fabric

Esse plano de aula é baseado em um experimento que tem como objetivo apresentar noções preliminares sobre como configurar e gerenciar uma rede blockchain usando a plataforma Hyperledger Fabric.  Nós adotamos a versão [Hyperledger Fabric 1.4 LTS](https://hyperledger-fabric.readthedocs.io/en/release-1.4/). Esta é uma versão *long term stable*, o que significa que deve se manter estável por pelo menos 2 anos.

A parte teórica do Hyperledger Fabric foi apresentada na aula sobre [blockchains](../blockchain). Revisite o conteúdo se você precisar relembrar qualquer conceito apresentado previamente.

## Usando uma rede blockchain customizada

Nosso experimento usa uma rede blockchain bastante simples, que possui apenas dois *peers* (um deles funciona como *endorser*) e usa o serviço *solo orderer* para implementar o consenso da rede. Nós também usamos containers [couchdb](https://hyperledger-fabric.readthedocs.io/en/release-1.4/couchdb_tutorial.html) para otimizar o desempenho durante o armazenamento de informações no ledger.

Os arquivos associados à configuração da nossa rede blockchain são os seguintes:

* [configtx.yaml](configtx.yaml): contém o *profile* da nossa rede Fabric.
* [crypto-config-ptb.yaml](crypto-config-ptb.yaml): contém a configuração do (Membership Service Provider). Nós geramos todos os certificados digitais necessários a partir deste arquivo.
* [docker-compose-ptb.yaml](docker-compose-ptb.yaml): contém a configuração dos nossos containers *docker* containers configuration. It extends the file [peer-base.yaml](blockchain/peer-base.yaml) which constitutes a template of standard configuration items.

If you are not used to the Hyperledger Fabric, we strongly recommend this [tutorial](https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html). It teachs in details how to create a basic Fabric network.

The Paillier Experiment network can be started by executing the steps described in the following subsections. All the commands must be executed into the folder blockchain.

### 1. Prepare the host machine

If you are working inside the PTB network, you need to configure your proxy settings for every tool that we will use in the next steps. But... smile! We create a [guide](prerequirements/PTB-Proxy.md) to make your life easier and save the time you would lose by solving this problem. Please check it before starting. You are welcomed!

Now, you need to install the **Hyperledger Fabric 1.4 LTS** basic software and [dependencies](https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html). Again, we try to make things simpler to you by providing a shell script that installs all these stuffs in a clean **Ubuntu 18.04 LTS** system. If you are using this distribution, our script works for you. If you have a different distribution, you can still try the script or you can customize it to work in your system.

Execute the [installation script](prerequirements/installFabric.sh):

```console
./installFabric.sh
```

**OBSERVATION**: you do not need to run the script as *sudo*. The script will automatically ask for your *sudo* password when necessary. That is important to keep the docker containers running with your working user account.

### 2. Generate the MSP artifacts

Before to execute this step, check if the environment variable FABRIC_CFG_PATH is properly defined. If it is not, uncomment the following line in the script [ptbMSP.sh](blockchain/ptbMSB.sh).

```console
export FABRIC_CFG_PATH=$PWD
```

After, in the folder [blockchain](blockchain), execute the script:

```console
./mspPTB.sh
```

This script uses [configtx.yaml](blockchain/configtx.yaml) and [crypto-config-ptb.yaml](blockchain/crypto-config-ptb.yaml) to create the MSP certificates in the folder (blockchain/crypto-config). It also generates the genesis block file *ptb-genesis.block* and the channel configuration file *ptb-channel.tx*. Noticed that this script depends on the tools installed together with Fabric. The script *installFabric.sh* executed previously is expected to modify your $PATH variable and enable the direct execution of the Fabric tools. If this does not happen, try to fix the $PATH manually. The tools usually are located in the folder /$HOME/fabric_samples/bin.

### 3. Manage the docker containers

We use the **docker-compose** tool to manage all the docker containers in our network. It basically reads the docker-compose-ptb.yaml and creates/starts/stops all the containers or a specific group of containers. You can find more details in the [Docker Compose Documents](https://docs.docker.com/compose/).

In the folder [blockchain](blockchain), use the following command to start all the containers:

```console
docker-compose -f docker-compose-ptb.yaml up -d
```

The same tool can be used to stop the containers, just in case you need to stop the blockchain network for any reason. In a similar manner as done before, use the following command to stop all the containers:

```console
docker-compose -f docker-compose-ptb.yaml stop
```

If you need to reset and restart a completly new blockchain network, use the following script to remove containers and clean all dependencies:

```console
./teardown.sh
```

### 4. Create the Fabric channel and join the peers

The last step consists on creating a channel (in practice, the ledger among the peers) and join all the active peer on it. That can be done by executing the following script into the folder blockchain:

```console
./startPTB.sh
```

If you succeed in coming so far, the Hyperledger Fabric shall be running in your machine, with an instance of the PTB network profile. You can see information from the containers by using the following commands:

```console
docker ps
docker stats
```

## The fabmorph chaincode

In this document, we assume you already know how to implement and deploy a chaincode in Fabric. If this is not your case, there is a [nice tutorial](https://hyperledger-fabric.readthedocs.io/en/release-1.4/chaincode4ade.html) that covers a lot of information about this issue. We strongly recomend you to check it before to continue.

If you already know everything you need about developing and deploying a chaincode, we can introduce the **fabmorph** chaincode. It contains basic methods that implement homomorphic computing using Paillier cryptosystem. The chaincode source code is [here](blockchain/fabmorph/fabmorph.go).

If you need to modify, compile and test the **fabmorph** chaincode, be sure that you have the [Golang *shim* interface](https://godoc.org/github.com/hyperledger/fabric/core/chaincode/shim) properly installed in your machine. If you do not have it, you can install it by using the following command:

```console
http_proxy="http://webproxy.berlin.ptb.de:8080" go get -u github.com/hyperledger/fabric/core/chaincode/shim
```

Notice that the provided command assumes you are working in the PTB network and inform the *http_proxy* environment variable. If that is not your case, you can just remove this variable and start your command with *"go get"*.

### Shell commands to deal with a Fabric chaincode

Our blockchain network profile includes the client container *cli0* which is provided only to execute tests with the chaincode. The *cli0* is able to communicate with the blockchain network using the peer *peer0.ptb.de* as an anchor and so execute commands for installing, mantaining and testing the chaincode. These commands documentation can be find [here](https://hyperledger-fabric.readthedocs.io/en/release-1.4/commands/peerchaincode.html). We strongly recommend you read this documentation before continuing.

#### 1. Installing, instantiating and upgrading a chaincode

Use the **install** command to enable the chaincode execution for a given peer. In practice, you are making this peer an __endorser__. You must reexecute the install command every time you change the chaincode version.

```console
docker exec cli0 peer chaincode install -n fabmorph -p github.com/hyperledger/fabric/peer/channel-artifacts/fabmorph -v 1.0
```

Use the **instantiate** command to instantiate the chaincode in a given channel. In practice, you are notifying the blockchain network that the chaincode exists. You also create a entry in the ledger with the chaincode hash.

```console
docker exec cli0 peer chaincode instantiate -o orderer.ptb.de:7050 -C ptb-channel -n fabmorph -v 1.0 -c '{"Args":[]}'
```

Use the **upgrade** command to enable a new version of the chaincode. That is necessary for any chaincode that was already instantiated before. Notice that a upgraded chaincode need to be re-installed in each one of its endorser peers.

```console
docker exec cli0 peer chaincode upgrade -o orderer.ptb.de:7050 -C ptb-channel -n fabmorph -v 1.1 -c '{"Args":[]}'
```

#### 2. Invoking and/or querying a chaincode

We can also use the same structure of commands in *cli0* to test our chaincode. Here we present some examples about how you can do that.

If you want to invoke the *registerMeter* chaincode method, informing a Paillier public key of 512 bits, you can use such command:

```console
docker exec cli0 peer chaincode invoke -o orderer.ptb.de:7050 -C ptb-channel -n fabmorph -c '{"Args":["registerMeter","666","512,9317708529424897702020382930116702407110920461631622943398410342176567591986029379418525445868266642035599843718252107477949966673217711651921453503121487,9317708529424897702020382930116702407110920461631622943398410342176567591986029379418525445868266642035599843718252107477949966673217711651921453503121488,86819692239317489725316183176771020092518150021809920212138582338877238932485744449732551095045029334111747236626667783012791667259632343428969631355566392144138916566553732073336450909522598178337851231139983633513413742668052750450987148230613227511148290981572584270867994369223203735104934252452681091169"]}'
```

After, you can retrieve the measurement consumption in the register made previously, by executing the command:

```console
docker exec cli0 peer chaincode invoke -o orderer.ptb.de:7050 -C ptb-channel -n fabmorph -c '{"Args":["getConsumption","666"]}'
```

## The Client Application

The Client Application is a set of modules that use the chaincode services provide by the blockchain network. All the client modules are writen in Python 3. They make use of the Python PHE Library (which implements the Paillier cryptosystem) and the Fabric Python SDK.

You need to install some dependencies and libraries before geting the clients running properly. We described all the steps necessary to prepare your machine to do that.

### Get pip3

Install the Python PIP3 using the following command:

```console
sudo apt install python3-pip
```

OBSERVATION: About the pip3, aware some packages can require specific versions due to compatibility issues. You can get a specific version using the following sintaxe:

```console
pip3 install --proxy http://webproxy.berlin.ptb.de:8080 -U 'pysha3 == 1.0b1'
```

### Get OPC-UA tools and dependencies

The [OPC Unified Architecture](https://opcfoundation.org/about/opc-technologies/opc-ua/) is a platform independent service-oriented architecture that integrates all the functionality of the individual OPC Classic specifications. OPC is an interoperability standard for the secure and reliable exchange of data in the industrial automation space.

Use pip3 to install the OPC-UA stack and its dependencias:

```console
pip3 install --proxy http://webproxy.berlin.ptb.de:8080 -U PyQt5 cryptography numpy opcua psutil pyqtgraph
```

### Get Partially Homomorphic Encryption (PHE) Python Library

The PHE library implements several homomorphic encryption algorithms, including the Paillier cryptosystem. Use pip3 to install it:

```console
pip3 install --proxy http://webproxy.berlin.ptb.de:8080 phe
```

### Get the Fabric Python SDK

The [Fabric Python SDK](https://github.com/hyperledger/fabric-sdk-py) is not part of the Hyperledger Project. It is maintained by an independent community of users from Fabric. However, this SDK works fine, at least to the basic functionalities we need.

Recently, we have problems with broke interfaces that made our programs stoped of running. So we adopted the 0.8.0 version of the Python SDK (and the respective tag) as our "stable" version.

You need to install the Python SDK dependencies first:

```console
sudo apt-get install python-dev python3-dev libssl-dev
```

Finally, install the Python SDK using *git*. Notice that the repository will be cloned into the current path, so we recommend you install it in your $HOME directory. After cloning the repository, is necessary to checkout the tag associated to the version 0.8.0.

```console
cd $HOME
git clone https://github.com/hyperledger/fabric-sdk-py.git
cd fabric-sdk-py
git checkout tags/v0.8.0
sudo make install
```

### Configure the .json network profile
The Python SDK applications depend on a network profile encoded in a .json format. In this repository, this file the [ptb-network.json](clients/ptb-network.json) file. The network profile keeps the necessary credentials to access the blockchain network. You must configure this file properly every time that you create new digital certificates in the MSP:

* Open the [ptb-network.json](clients/ptb-network.json) in a text editor;
* Check for the entries called "private_key" on each organization. Notice that they points out to a file into the (../blockchain) directory that corresponds to the private key of each organization;
* Check the MSP file structure in your deployment and verify the correct name of the files that contain the private key;
* Modify the .json file with the correct name of the files.

### The Client Application modules

The Client Application includes the following modules:

* [server.py](clients/server.py): It is a basic OPCUA server which emulates a meter (or a set of meters), generating a new measurement each second. Its implementation was taken from other PTB
project and is used as it is;
* [keygen-morph.py](clients/keygen-morph.py): It is a simple Python script that generates a pair of Paillier keys. These keys are necessary to execute all the other   modules, if you want to do homomorphic computing.
* [register-morph.py](clients/register-morph.py): It invokes the *registerMeter* chaincode, that appends a new meter digital asset into the ledger. The respective Paillier public key should be provided,Otherwise, the module understands you will work with plaintext measurements.
* [client-morph.py](clients/client-morph.py): It works as a client that collect measures from the OPCUA server and invoke *insertMeasurement* for encrypted measurements. It does that continually, and the developer can choose between a fix sleep time or a keypress from the user to continue generating measurements.
* [display-morph.py](clients/display-morph.py): It works as a simple client which queries the ledger and retrieves the digital asset relatad to a respective meter. This module requires the Paillier private key to decrypted the encrypted consumption measurement.
* [client-morph-mt.py](clients/client-morph-mt.py): It creates multiple threads that collect measures from the OPCUA server and invoke *insertMeasurement* (for encrypted measurements) or the *insertedPlainTextMeasurement* (for plaintext measurements) chaincode. The choice depends on the existence (or not) of the meter's Paillier public key. The module also logs statistic information related to the performance of the blockchain peers. AT THE MOMENT, THIS MODULE IS NOT FUNCTIONAL DUE TO PROBLEMS IN USING ASSYNCHRONOUS FABRIC FROM PYTHON SDK AND TRANSACTIONS AND MULTITHREADS.
<!--stackedit_data:
eyJoaXN0b3J5IjpbMzEzNzgwNjk3XX0=
-->