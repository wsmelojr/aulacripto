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

Como qualquer chaincode, o fabnotas precisa ser instalado em um *peer* (endorser) e instanciado no channel. Para isso, utilize a mesma rede (PTB) disponível na pasta [hyperledger](..). Você deve utilizar os mesmos comandos executados na aula anterior, mas substituindo o fabmorph (chaincode usado anteriormente) pelo fabnotas (novo chaincode).

### Questão 3: Instale o fabnotas no *peer0.ptb.de* e depois o instancie ele no channel *ptb-channel*. Copie as saidas de tela dos comandos como resposta dessa questão.

Uma vez que o chaincode esteja instalado e instanciado, você pode executar qualquer uma das funções disponíveis nele. O comando de execução se baseia no uso do container *cli0*, tal como fizemos na [aula anterior](..). O comando a seguir serve como exemplo de como invocar a função *inserirAvaliacao*, criando o ID de avaliação *1234* e inserindo a questão *\"Quem é o hacker mais famoso do Inmetro?\"*.

    docker exec cli0 peer chaincode invoke -o orderer.ptb.de:7050 -C ptb-channel -n fabnotas -c '{"Args":["inserirAvaliacao","1234","Quem é o hacker mais famoso do Inmetro?"]}'

Note que os parametros em *Args* devem ser informados de acordo com a função que se deseja executar. Os parâmetros devem seguir a ordem estrita determinada na implementação da função. A sintaxe (uso de colchetes e aspas) também deve ser observada rigorosamente.

Um outro exemplo dado a seguir consiste em invocar a função *responderAvaliacao* para atribuir uma resposta à questão inserida anteriormente. Repetimos como argumento o mesmo ID de avaliação *1234* e inserimmos uma resposta à questão: "O nome dele é Ewerton Madruga".

    docker exec cli0 peer chaincode invoke -o orderer.ptb.de:7050 -C ptb-channel -n fabnotas -c '{"Args":["responderAvaliacao","1234","O nome dele é Ewerton Madruga"]}'

### Questão 4: Execute as funções *corrigirAvaliacao* e *countLedger* disponíveis no chaincode, seguindo a mesma estrutura de comandos. Copie as saídas de tela dos comandos como resposta desta questão.

Para fechar esta etapa da atividade, insira pelo menos mais 5 avaliações no blockchain, com suas respectivas respostas e notas. Verifique se o número de avaliações coincide com o total de chaves (keys) informado pela função *countLedger*.

## Conhecendo melhor o Couchdb
Uma das vantagens de se utilizar o CouchDB é fato do mesmo se tratar de um sistema de gerenciamento de banco de dados (SGDB) e possuir uma interface pŕopria para acesso às suas informações.

Se você tem a rede blockchain executando em sua máquina, você pode acessar a interface web do projeto **Fauxton** associada ao CouchDB. Para acessar a base de dados do container *peer0.ptb.de* clique [aqui](http://localhost:5984/_utils/). Acesse o link para os registros do chaincode fabnotas clicando no link ***ptb-channel_fabnotas***. Você poderá consultar todos os ativos digitais associados a este chaincode.

**IMPORTANTE**: A interface web de acesso ao CouchDB **NÃO** utiliza os mecanismos de acesso do blockchain. Ela dá uma visão correspondente à cópia local do ledger contida no container *peer0.ptb.de*. Lembre-se que qualquer alteração realizada diretamente na cópia local pode comprometer essa réplica do ledger.

O CouchDB disponibiliza a ferramenta **Mango** para consulta livre (chamamos de *queries*) em suas bases de dados. Essas *queries* são escritas em formato JSON. Sua formatação é a princípio um pouco complexa, mas ainda assim facilita bastante a consulta das informações armazenadas no CouchDB. Você pode aprender um pouco sobre o Mango [aqui](https://blog.4linux.com.br/introducao-ao-couchdb/).

O código abaixo ilustra como executar uma *query* usando o Mango. Clique na opção ***Run a Query with Mango***, copie o código abaixo para a janela de comando, clique no botão ***Run Query*** e observe o resultado. Você vai ver que essa query retorna todas as avaliações onde o campo *resposta* contém a palavra *Madruga*.

    {"selector":{
      "resposta": {
        "$regex": "Madruga"
      }
    }}

Tente praticar outras *queries* informando diferentes parâmetros e operadores. Uma lista de comandos está disponível no tutorial sobre o CouchDB informado no link anterior.

Agora que conhecemos melhor o CouchDB, é possível testar também a função *queryLedger* no chaincode fabnotas. Esta função permite executar diretamente uma consulta (*query*) no ledger, todavia usando a rede blockchain, e não a cópia local de um determinado *peer*. Assim como no Mango, a *query* precisa ser formatada como uma expressão JSON. O comando a seguir exemplifica como isso pode ser feito:

    docker exec cli0 peer chaincode invoke -o orderer.ptb.de:7050 -C ptb-channel -n fabnotas -c '{"Args":["queryLedger","{\"selector\":{\"resposta\":{\"$regex\":\"Madruga\"}}}"]}'

Observe que a consulta é a mesma realizada diretamente no Mango. O único problema é que precisamos usar um caracter "\\" antes de cada aspas interna à estrutura JSON. Isso é necessario para que o terminal entenda que essas aspas são símbolos literais e não comandos de formatação do shell. Esse formato é comumente descrito como "escape" dos caracteres.

## Customizando o chaincodes

O chaincode fabnotas pode ser customizado para acrescentar funções adicionais ou mesmo modificar funções existentes. Com isso, você pode escrever seu próprio código fonte e executá-lo como parte do chaincode. Para fazer isso, siga os seguintes passos:

1. Planeje a sua função. Pense no nome, e nos argumentos que serão necessários;
2. Escreva o código da sua função no arquivo .go. Você pode se basear na estrutura das outras funções, respeitando o uso da interface *shim*;
3. Crie o *entry point* de sua função dentro do método *Invoke()*;
4. Compile o chaincode modificado com o comando *go build*. Se houver erros de compilação, corrija os problemas e só prossiga quando não houver mais erros;
5. Instale e instancie o chaincode em sua rede blockchain. Se o chaincode já foi instalado previamente, faça a instalação incrementando o número da versão e execute o comando *upgrade* ao invés do comando *instantiate* para instanciar o chaincode.

### Questão 5: É hora de você escrever e executar sua propria função dentro *chaincode*. Pense em uma função que seja útil para o funcionamento deste chaincode, escreva o código da mesma e teste sua execução no blockchain. Copie o código fonte de sua função e as saídas de tela dos comandos como resposta desta questão.
