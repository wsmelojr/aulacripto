*Data de aula: 21 e 23 de Outubro*

*Prof. Wilson ([wsjunior@inmetro.gov.br](mailto:wsjunior@inmetro.gov.br))*

# Criptografia Homomórfica

A criptografia de chave pública (também chamada de criptografia assimétrica) é sem dúvida um dos recursos mais importantes que temos para se implementar diferentes mecanismos de segurança. Como já vimos em outros estudos, ela permite atender requisitos de confidencialidade, integridade, autenticidade e não-repudio para um dada informação. Entretanto há cenários onde a criptografia assimétrica convencional acarreta problemas críticos de desempenho.

Considere, por exemplo, o seguinte cenário. Você necessita manter um banco de dados acessível a clientes de um determinado serviço. Entretanto, o volume de dados é muito grande para você manter o serviço em sua própria infraestrutura. Uma alternativa é contratar um serviço de hospedagem de servidores em nuvem (e.g., um *pool* de servidores da Amazon), e deste modo você não precisa mais se preocupar com detalhes técnicos para manter o serviço disponível ou expandir seus discos à medida em que o volume de dados aumenta.

Todavia, existe outro problema: os dados são sigilosos, de modo que apenas os clientes podem ter acesso aos mesmos. Manter o serviço executando em uma infraestrutura de terceiros incorre no risco de expor esses dados a acesso indevido. Uma alternativa simples é armazenar esses dados criptografados. Você restringe o acesso às chaves criptográficas aos clientes, e assim apenas eles conseguem escrever e recuperar os dados. Os mantenedores do serviço, embora tenham acesso aos dados, não possuem as chaves criptográficas necessárias para decifrar a informação ali contida.

Isso parece resolver o problema. Todavia, ainda há uma questão crítica a se resolver: desempenho. Como os dados estão criptografados, você tem uma dificuldade imensa para acessá-los de forma eficiente. Suponha, por exemplo, que você deseja consultar sua base de dados e selecionar todos os registros de clientes que possuem mais de 40 anos. Isso seria **trivial** em um banco de dados não criptografado: uma simples *query* filtrando os dados de acordo com a idade. Entretanto, em um banco de dados criptografado, a informação correspondente a idade também estará encriptada. O sistema de banco de dados não consegue realizar essa consulta, sem primeiro decriptografar todos os registros de clientes, para só então poder comparar suas idades.

### Questão 1: Reflita sobre o problema apresentado nessa introdução, e tente imaginar uma solução para o mesmo. O que você faria?

## Conceitos fundamentais

A criptografia homomórfica surge como uma solução para o problema discutido acima. A ideia consiste em se desenvolver criptossistemas que possibilitem a realização de computação (operações lógicas e aritméticas) usando dados criptografados, de forma que o resultado obtido, uma vez decriptografado, seja os mesmo daquele que seria obtido em uma operação com os dados não criptografados.

[![VideoAula](https://img.youtube.com/vi/NBO7t_NVvCc/0.jpg)](https://www.youtube.com/watch?v=NBO7t_NVvCc "Clique para assistir um video divertido sobre o conceito elementar de criptografia homomórfica!")

Em um primeiro momento, podemos ser levados a pensar que isso acontece para qualquer algoritmo criptografado. Mas tal suposição não é verdadeira. Lembre-se que, quando criptografamos uma informação, o dado criptografado consiste em uma transformação complexa da informação, de modo que as operações lógicas e matemáticas que serão aplicadas sobre dos dados criptografados não serão mais idênticas àquelas aplicadas sobre os dados em não criptografados. Para isso, é necessário se descobrir uma **equivalência** entre propriedades matemáticas e lógicas do dominío da informação e do domínio criptográfico. Uma analogia simples é observar o que acontece quando você olha seu reflexo em um espelho. Todos os movimentos passam a ser invertidos. Quando você levanta sua mãe direita, a imagem parece mostrar que você está levantando sua mão esquerda, e vice-versa. Da mesma forma, a soma de dois números inteiros no domínio dos números naturais não será equivalente à soma de dois números criptográficos. Essa operação certamente será diferente (por exemplo, uma multiplicação ou potenciação).

A figura a seguir dá uma ideia de como isso acontece. Ela supõe que a transformação criptográfica consiste em multiplicar o número por 2 (uma criptografia que seria ridiculamente fácil de se "quebrar"). Entretanto, na prática, tal transformação é o suficiente para afetar as operações aritméticas que são feitas com esses números.

![Exemplo de um criptossistema homomórfico extremamente simples](exemplo1.jpg)


Esse plano de aula é baseado em um experimento que tem como objetivo apresentar noções preliminares sobre como configurar e gerenciar uma rede blockchain usando a plataforma Hyperledger Fabric.  Nós adotamos a versão [Hyperledger Fabric 1.4 LTS](https://hyperledger-fabric.readthedocs.io/en/release-1.4/). Esta é uma versão *long term stable*, o que significa que deve se manter estável por pelo menos 2 anos.

A parte teórica do Hyperledger Fabric foi apresentada na aula sobre [blockchains](../blockchain). Revisite o conteúdo se você precisar relembrar qualquer conceito apresentado previamente.

## Usando uma rede blockchain customizada

Nosso experimento usa uma rede blockchain bastante simples, que foi desenvolvida em um projeto em parceria com o PTB, o NMI da Alemanha. Essa rede possui apenas dois *peers* (um deles funciona como *endorser*) e usa o serviço *solo orderer* para implementar o consenso da rede. Você vai notar durante o experimento que todos os *peers* tem seu nome associado ao domínio *ptb.de*. Nós também usamos containers [couchdb](https://hyperledger-fabric.readthedocs.io/en/release-1.4/couchdb_tutorial.html) para otimizar o desempenho durante o armazenamento de informações no ledger.

Os arquivos associados à configuração da nossa rede blockchain são os seguintes:

* [configtx.yaml](configtx.yaml): contém o *profile* da nossa rede Fabric.
* [crypto-config-ptb.yaml](crypto-config-ptb.yaml): contém a configuração do (Membership Service Provider). Nós geramos todos os certificados digitais necessários a partir deste arquivo.
* [docker-compose-ptb.yaml](docker-compose-ptb.yaml): contém a configuração dos nossos containers *docker*. Este arquivo usa como base as definições contidas no arquivo [peer-base.yaml](peer-base.yaml), que constitui um modelo de configuração de containers padrão.

### Questão 1: Examine o arquivo docker-compose-ptb.yaml e acompanhe a explanação do professor sobre a arquitetura da rede blockchain representada neste arquivo. Liste cada um dos 7 containers que integram a rede e explique de forma sucinta a função de cada um deles.

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

### Questão 2: Examine o diretório de trabalho e veja se o bloco genesis e a configuração do channel foram criados. Agora, explique com suas palavras o que é o bloco genesis e o que é um channel.

### 3. Gerencie os containers docker

Nós usamos a ferramenta **docker-compose** para gerenciar todos os containers docker da nossa rede. Essa ferramenta basicamente lê o arquivo [docker-compose-ptb.yaml](docker-compose-ptb.yaml) e cria/inicia/pára todos os containers ou um grupo especifico de containers. Você pode aprender mais sobre o docker-compose [aqui](https://docs.docker.com/compose/).

### Questão 3: Explique com suas palavras o que é um container, e qual a diferença básica entre um container e uma máquina virtual.

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
### Questão 4: Examine cuidadosamente o script startPTB.sh. Note que há uma diferença sutil na inicialização dos containers  peer0 e peer1. Você sabe dizer que diferença é essa? Explique porque ela ocorre.

Se você conseguiu chegar até aqui, nossa rede blockchain usando o Hyperledger Fabric estará executando em sua máquina, com uma instância do *profile* de rede utilizado. Você pode confirmar isso buscando por informações da rede. Para tanto, use os seguintes comandos:

```console
docker ps
docker stats
```

Um comando muito útil é o *logs*, que permite consultar o registro de log de um container específico. Você pode por exemplo verificar o log do container *peer0.ptb.de* executando o comando a seguir:
```console
docker logs peer0.ptb.de
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

Se você quer invocar o método *registerMeter* informando uma chave criptográfica de 512 bits, você pode usar o seguinte comando:

```console
docker exec cli0 peer chaincode invoke -o orderer.ptb.de:7050 -C ptb-channel -n fabmorph -c '{"Args":["registerMeter","666","512,9317708529424897702020382930116702407110920461631622943398410342176567591986029379418525445868266642035599843718252107477949966673217711651921453503121487,9317708529424897702020382930116702407110920461631622943398410342176567591986029379418525445868266642035599843718252107477949966673217711651921453503121488,86819692239317489725316183176771020092518150021809920212138582338877238932485744449732551095045029334111747236626667783012791667259632343428969631355566392144138916566553732073336450909522598178337851231139983633513413742668052750450987148230613227511148290981572584270867994369223203735104934252452681091169"]}'
```

De forma semelhante, você pode consultar o consumo do medidor registrado anteriormente, executando o comando a seguir:

```console
docker exec cli0 peer chaincode invoke -o orderer.ptb.de:7050 -C ptb-channel -n fabmorph -c '{"Args":["getConsumption","666"]}'
```

### Questão 5: Suponha que você esteja administrando uma rede blockchain de 10 *peers*, sendo que 3 deles são endorsers de um mesmo chaincode. O desenvolvedor do chaincode então comunica a você que uma nova versão do chaincode foi escrita. O que você precisará fazer?
