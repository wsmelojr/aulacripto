*Data de aula: 23 e 25 de Setembro*

*Prof. Wilson (wsjunior@inmetro.gov.br)*

# Criação de Chaincodes para o Hyperledger Fabric

A aula dessa semana consiste em analisar e customizar um *chaincode*, que nada mais é do que um *smart contract* escrito para a plataforma [Hyperledger Fabric 1.4 LTS](https://hyperledger-fabric.readthedocs.io/en/release-1.4/).

Se você precisar revisar os conceitos prévios, você pode acessar o conteúdo disponibilizado para as aulas sobre [blockchains](../../blockchain) e também sobre o [Hyperledger Fabric](..).

## Visão geral do chaincode

O *chaincode* que vamos utilizar nessa aula se chama [fabnotas](fabnotas.go). O fabnotas foi pensado como uma aplicação simples para que os professores do curso de Segurança Cibernética registrem suas avaliações em um blockchain. As avaliações são representadas por um identificador e expressas em questões. Os alunos devem responder as questões, e a resposta também ficará registrada no blockchain. Por fim, os professores corrigem as questões e registram a nota também no blockchain. Deste modo, todas as etapas da avaliação ficam registradas no ledger. Isso significa que as informações:
* não podem ser excluidas;
* são replicadas entre todos os participantes da rede blockchain.

## Estrutura do fabnotas

O fabnotas utiliza um ativo digital (*digital asset*) para armazenar as informações de uma avaliação. Ele segue a seguinte estrutura:

    type Avaliacao struct {
      Questao string `json:"questao"`
      Resposta string `json:"resposta"`
      Nota int `json:"nota"`
    }

Note que o ativo digital é basicamente uma estrutura [JSON](https://www.json.org/json-pt.html). O JSON (JavaScript Object Notation) é um formato padrão de apresentação da informação, que é suportado nativamente por diversas ferramentas de desenvolvimento e sistemas de bancos de dados.

O fabnotas segue a estrutura de um *chaincode* padrão. Ele tem dois métodos principais que implementam a interface [Golang *shim*](https://godoc.org/github.com/hyperledger/fabric/core/chaincode/shim). Essa interface funciona como um *plug* para que permite que nosso código seja conectado a uma rede blockchain.

### Questão 1: Acesse a documentação da interface *shim* no link informado acima. Com auxílio do professor, escolha duas funções da interface que você acho interessante e explique o que elas fazem, usando suas próprias palavras.

Os métodos que o fabnotas implementa são o *Init()* e o *Invoke()*. O primeiro é executado sempre que o chaincode é **instanciado**.

    func (s SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
    	return shim.Success(nil)
    }

O *Invoke()*, por sua vez, funciona como um *distribuidor* das funções do chaincode criadas pelo programador. Ele é o método invocado sempre que o chaincode é utilizado (por isso o chamamos de *entry point*). Dentro dele, a função especifica solicitada é selecionada, e o fluxo de código é transferido para essa função.

    func (s SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
    	//extrai o nome da funcao e os argumentos associados a uma transacao
    	fn, args := stub.GetFunctionAndParameters()

    	//implementa o tratamento para cada funcao reconhecida pelo chaincode
    	if fn == "inserirAvaliacao" {
    		//permite a um professor cadastrar uma avaliacao especifica
    		return s.inserirAvaliacao(stub, args)

    	} else if fn == "responderAvaliacao" {
    		//permite ao aluno responder a avaliacao
    		return s.responderAvaliacao(stub, args)

    	} else if fn == "corrigirAvaliacao" {
    		//permite ao professor corrigir e atribuir nota a avaliacao
    		return s.corrigirAvaliacao(stub, args)

### Questão 2: Localize as linhas onde os métodos *Init()* e o *Invoke()* estão no fabnotas. Leia a implementação e os comentários deixados sobre os mesmos. Em seguida, liste quais funções são tratadas pelo método *Invoke()*, e explique o que cada uma delas faz.

## Compilando o fabnotas

Para compilar o fabnotas, você precisa ter o compilador da linguagem **Golang** instalado na sua máquina. Isso a princípio foi feito junto com a instação do Hyperledger Fabric, na aula anterior. Para ter certeza que o **Golang** está instalado, digite o seguinte comando no terminal:

    $ go version
    go version go1.7.1 linux/amd64

Se o **Golang** não foi encontrado, será necessário repetir o procedimento de instalação disponível na aula sobre o [Hyperledger Fabric 1.4 LTS](..).

O próximo passo é instalar a interface *shim*. Para isso, execute o seguinte comando:

    $ go get -u github.com/hyperledger/fabric/core/chaincode/shim

**OBSERVAÇÃO**: a execução desse comando demora um pouco, e a interface não oferece *feedback* de progresso. Mesmo assim, é só aguardar um pouco que a instalação deve ocorrer sem problemas.

Com a interface *shim* instalada, você pode compilar o fabnotas acessando o diretorio onde o chaincode se encontra, e executando o comando *go build*, conforme apresentado a seguir:

    $ cd fabnotas
    $ go build

## Executando o fabnotas

Como qualquer chaincode, o fabnotas precisa ser instalado em um *peer* (endorser) e instanciado no channel. Para isso, utilize a mesma rede (PTB) disponível na pasta [hyperledger](..). Você deve utilizar os mesmos comandos executados na aula anterior, mas substituindo o fabmorph (chaincode usando anteriormente) pelo fabnotas (novo chaincode).

### Questão 3: Instale o fabnotas no *peer0.ptb.de* e depois o instancie ele no channel *ptb-channel*. Copie as saidas de log dos comandos como resposta dessa questão.

*O conteúdo dessa aula ainda não está completo. Mais atividades serão adicionadas aqui nos próximos dias!*
