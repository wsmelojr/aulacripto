*Data de aula: 9 e 11/9*

*Prof. Wilson ([wsjunior@inmetro.gov.br](mailto:wsjunior@inmetro.gov.br))*

# Hyperledger Fabric

Esse plano de aula é baseado em um experimento que tem como objetivo apresentar noções preliminares sobre como configurar e gerenciar uma rede blockchain usando a plataforma Hyperledger Fabric.  Nós adotamos a versão [Hyperledger Fabric 1.4 LTS](https://hyperledger-fabric.readthedocs.io/en/release-1.4/). Esta é uma versão *long term stable*, o que significa que deve se manter estável por pelo menos 2 anos.

A parte teórica do Hyperledger Fabric foi apresentada na aula sobre [blockchains](../blockchain). Revisite o conteúdo se você precisar relembrar qualquer conceito apresentado previamente.

## Usando uma rede blockchain customizada

Nosso experimento usa uma rede blockchain bastante simples, que foi desenvolvida em um projeto em parceria com o PTB, o NMI da Alemanha. Essa rede possui apenas dois *peers* (um deles funciona como *endorser*) e usa o serviço *solo orderer* para implementar o consenso da rede. Você vai notar durante o experimento que todos os *peers* tem seu nome associado ao domínio *ptb.de*. Nós também usamos containers [couchdb](https://hyperledger-fabric.readthedocs.io/en/release-1.4/couchdb_tutorial.html) para otimizar o desempenho durante o armazenamento de informações no ledger.

Os arquivos associados à configuração da nossa rede blockchain são os seguintes:

* [configtx.yaml](configtx.yaml): contém o *profile* da nossa rede Fabric.
* [crypto-config-ptb.yaml](crypto-config-ptb.yaml): contém a configuração do (Membership Service Provider). Nós geramos todos os certificados digitais necessários a partir deste arquivo.
* [docker-compose-ptb.yaml](docker-compose-ptb.yaml): contém a configuração dos nossos containers *docker*. Este arquivo usa como base as definições contidas no arquivo [peer-base.yaml](peer-base.yaml), que constitui um modelo de configuração de containers padrão.

O experimento deve ser executado conforme os passos descritos a seguir.

### 1. Prepare sua máquina para executar o Fabric

Você precisa instalar o software básico do **Hyperledger Fabric 1.4 LTS** e também suas [dependências](https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html). Para simplificar esta etapa, nós utilizamos um *script* que instala todo o software necessário a partir de uma instalação "limpa" do sistema **Ubuntu 18.04 LTS**.

Execute o [script de instalação](installFabric.sh):

```console
./installFabric.sh
```
### 2. Crie os artefatos do MSP

Antes de executar este passo, verifique se a variável de ambiente FABRIC_CFG_PATH foi propriamente configurada. Se não, descomente a linha respectiva no script [mspPTB.sh](mspPTB.sh).

```console
export FABRIC_CFG_PATH=$PWD
```

Em seguida, execute o script:

```console
./mspPTB.sh
```

Esse script usa os arquivos [configtx.yaml](configtx.yaml) e [crypto-config-ptb.yaml](crypto-config-ptb.yaml) para criar os certificados MSP em uma pasta chamada *crypto-config*. Ele também gera o **bloco genesis** em um arquivo chamado *ptb-genesis.block* e a configuração do *channel* no arquivo *ptb-channel.tx*.

### 3. Gerencie os containers docker

Nós usamos a ferramenta **docker-compose** para gerenciar todos os containers docker da nossa rede. Essa ferramenta basicamente lê o arquivo [docker-compose-ptb.yaml](docker-compose-ptb.yaml) e cria/inicia/pára todos os containers ou um grupo especifico de containers. Você pode aprender mais sobre o docker-compose [aqui](https://docs.docker.com/compose/).

Vamos prosseguir com nosso experimento. Execute o seguinte comando para iniciar todos os containers da nossa rede blockchain:

```console
docker-compose -f docker-compose-ptb.yaml up -d
```

O mesmo comando pode ser usado para parar todos os containers. Isso é útil se você precisar parar a rede blockchain por qualquer motivo. Da mesma forma como fizemos antes, use o comando a seguir para parar todos os containers:

```console
docker-compose -f docker-compose-ptb.yaml stop
```

Se você precisa destruir e recriar os containers da rede blockchain, utiize o seguinte script: 

```console
./teardown.sh
```

Esse comando limpa toda a área de containers criados previamente, e elimina dependências associadas a outros containers temporários.

### 4. Crie um *channel* e conecte os *peers*

O último passo na criação da rede blockchain consiste em criar um *channel* (na prática, uma instância do ledger compartilhada entre os *peers*) e conectar todos os *peers* ativos ao channel. Para isso execute o script [startPTB.sh](startPTB.sh):

```console
./startPTB.sh
```

Se você conseguiu chegar até aqui, nossa rede blockchain usando o Hyperledger Fabric estará executando em sua máquina, com uma instância do *profile* de rede utilizado. Você pode confirmar isso buscando por informações da rede. Para tanto, use os seguintes comandos:

```console
docker ps
docker stats
```

## Trabalhando com *chaincodes*

Apenas relembrando, *chaincode* é o nome dado a um *smart contract* na plataforma Hyperledger Fabric. Por enquanto, vamos utilizar um *chaincode* pronto. Em laboratórios futuros iremos escrever nosso próprio *chaincode*!

Nós usaremos então um chaincode chamado **fabmorph**. Ele contém funções básicas para implementar cripografia homomórfica (outro tema que estudaremos no futuro). O código fonte do chaincode está disponível [aqui](fabmorph/fabmorph.go).

Nossa rede blockchain inclui um container especifico para simular clientes, que é o container *cli0*. Nós vamos usar esse container para testar o chaincode. O container *cli0* se comunica com *peer0.ptb.de*, simulando a execução de transações.

### 1. Instalando, instanciando e atualizando um chaincode

Use o comando **install** para habilitar a execução de um chaincode por um determinado *peer*. Na pratica, você transformará esse *peer* em um  **endorser**. Você precisa reexecutar o comando *install* sempre que fizer qualquer modificação no chaincode.

```console
docker exec cli0 peer chaincode install -n fabmorph -p github.com/hyperledger/fabric/peer/channel-artifacts/fabmorph -v 1.0
```

Use o comando **instantiate** para instanciar o chaincode no respectivo *channel*. Quando fazemos isso, estamos notifciando a rede blockchain a respeito da existência do chaincode. Esse comando também criará um registro no ledger com o hash do chaincode.

```console
docker exec cli0 peer chaincode instantiate -o orderer.ptb.de:7050 -C ptb-channel -n fabmorph -v 1.0 -c '{"Args":[]}'
```

Use  o comando **upgrade** para habilitar uma nova versão do chaincode. Isso é necessário para qualquer chaincode que já tenha sido instanciado anteriormente. Note que ao realizar o upgrade de um chaincode, você precisa reinstalar o mesmo em todos os *endorsers*.

```console
docker exec cli0 peer chaincode upgrade -o orderer.ptb.de:7050 -C ptb-channel -n fabmorph -v 1.1 -c '{"Args":[]}'
```

### 2. Invocando e/ou consultando  um chaincode

Nós usaremos a mesma estrutura de comandos do container *cli0* para testar o chaincode instalado. As linhas de comando a seguir são exemplos de como isso pode ser feito.

Se você quer invocar o método *registerMeter* informandochaincode method, informing a Paillier public key of 512 bits, you can use such command:

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
eyJoaXN0b3J5IjpbLTE3MTg0NjMzODksMTcwMzkxMDAxNSwxND
I2OTE5NzQxLC0xOTcxMjYwNzVdfQ==
-->