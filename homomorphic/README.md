*Data de aula: 21 e 23 de Outubro*

*Prof. Wilson ([wsjunior@inmetro.gov.br](mailto:wsjunior@inmetro.gov.br))*

# Criptografia Homomórfica

A criptografia de chave pública (também chamada de criptografia assimétrica) é sem dúvida um dos recursos mais importantes que temos para se implementar diferentes mecanismos de segurança. Como já vimos em outros estudos, ela permite atender requisitos de confidencialidade, integridade, autenticidade e não-repudio para uma determinada informação. Entretanto há cenários onde a criptografia assimétrica convencional acarreta problemas críticos de desempenho.

Considere, por exemplo, o seguinte cenário. Você necessita manter um banco de dados acessível a clientes de um determinado serviço. Entretanto, o volume de dados é muito grande para você manter o serviço em sua própria infraestrutura. Uma alternativa é contratar um serviço de hospedagem de servidores em nuvem (e.g., um *pool* de servidores da Amazon), e deste modo você não precisa mais se preocupar com detalhes técnicos para manter o serviço disponível ou expandir seus discos à medida em que o volume de dados aumenta.

Todavia, existe outro problema: os dados são sigilosos, de modo que apenas os clientes podem ter acesso aos mesmos. Manter o serviço executando em uma infraestrutura de terceiros incorre no risco de expor esses dados a acesso indevido. Uma alternativa simples é armazenar esses dados criptografados. Você restringe o acesso às chaves criptográficas aos clientes, e assim apenas eles conseguem escrever e recuperar os dados. Os mantenedores do serviço, embora tenham acesso aos dados, não possuem as chaves criptográficas necessárias para decifrar a informação ali contida.

Isso parece resolver o problema. Todavia, ainda há uma questão crítica a se resolver: desempenho. Como os dados estão criptografados, você tem uma dificuldade imensa para acessá-los de forma eficiente. Suponha, por exemplo, que você deseja consultar sua base de dados e selecionar todos os registros de clientes que possuem mais de 40 anos. Isso seria **trivial** em um banco de dados não criptografado: uma simples *query* filtrando os dados de acordo com a idade. Entretanto, em um banco de dados criptografado, a informação correspondente à idade também estará encriptada. O sistema de banco de dados não consegue realizar essa consulta, sem primeiro descriptografar todos os registros de clientes, para só então poder comparar suas idades.

### Questão 1: Reflita sobre o problema apresentado nessa introdução, e tente imaginar uma solução para o mesmo. O que você faria?

## Conceitos fundamentais

A criptografia homomórfica surge como uma solução para o problema discutido acima. A ideia consiste em se desenvolver criptossistemas que possibilitem a realização de computação (operações lógicas e aritméticas) usando dados criptografados, de forma que o resultado obtido, uma vez descriptografado, seja os mesmo daquele que seria obtido em uma operação com os dados não criptografados.

[![VideoAula](https://img.youtube.com/vi/NBO7t_NVvCc/0.jpg)](https://www.youtube.com/watch?v=NBO7t_NVvCc "Clique para assistir um video divertido sobre o conceito elementar de criptografia homomórfica!")

Em um primeiro momento, podemos ser levados a pensar que isso acontece para qualquer algoritmo criptografado. Mas tal suposição não é verdadeira. Lembre-se que, quando criptografamos uma informação, o dado criptografado consiste em uma transformação complexa da informação, de modo que as operações lógicas e matemáticas que serão aplicadas sobre dos dados criptografados não serão mais idênticas àquelas aplicadas sobre os dados em não criptografados. Para isso, é necessário se descobrir uma **equivalência** entre propriedades matemáticas e lógicas do dominío da informação e do domínio criptográfico. Uma analogia simples é observar o que acontece quando você olha seu reflexo em um espelho. Todos os movimentos passam a ser invertidos. Quando você levanta sua mãe direita, a imagem parece mostrar que você está levantando sua mão esquerda, e vice-versa. Da mesma forma, a soma de dois números inteiros no domínio dos números naturais não será equivalente à soma de dois números criptográficos. Essa operação certamente será diferente (por exemplo, uma multiplicação ou potenciação).

A figura a seguir dá uma ideia de como isso acontece. Ela supõe que a transformação criptográfica consiste em multiplicar o número por 2 (uma criptografia que seria ridiculamente fácil de se "quebrar"). Entretanto, na prática, tal transformação é o suficiente para afetar as operações aritméticas que são feitas com esses números. Deste modo, você pode ter uma ideia do quanto um algoritmo criptográfico eficiente deve acrescentar em termos de complexidade no tratamento de diretivas computacionais no domínio criptografado. Na grande maioria dos casos, a implementação dessas diretivas computacionais é impraticável.

![Exemplo de um criptossistema homomórfico extremamente simples](exemplo1.jpg)

Assim, criptossistemas que conservam propriedades que permitem a realização de operações lógicas e aritméticas no domínio criptográfico são verdadeiros "achados" no campo da criptografia. Essa se tornou uma área muito interessante de pesquisa, e muitos criptográfos se dedicam atualmente a descobrir novos criptossistemas que possuem essas propriedades ou então aprimorar os criptossistemas já descobertos. Esses criptossistemas são chamados **homomórficos**, dando a ideia que eles mantem propriedades equivalentes em ambos os "lados do espelho".

### Questão 2: Reflita um pouco e discuta com o colega ao lado sobre o conceito aprendido de criptografia homomórfica. Tente imaginar uma possível aplicação, além daquela descrita no início dessa aula (referente à proteção do banco de dados). Descreva seu exemplo de aplicação, e como poderia ser implementado.

## Aplicações Pŕaticas

As aplicações envolvendo criptografia homomórfica em geral estão associadas a problemas de privacidade. Trata-se de cenários onde um conjunto de informações precisa ser computado para um fim específico, mas sem revelar partes individuais da informação que dizem respeito a uma pessoa, ou a uma empresa, ou a uma entidade isoladamente.

Estes são alguns exemplos de aplicações práticas envolvendo criptografia homomórfica:
* Um hospital ou um centro de pesquisa na área de saúde pode coletar o DNA de pacientes/voluntários para fins de estudos científicos de natureza estatística. Todavia, o DNA de cada indivíduo é criptografado, de modo a proteger sua privacidade.
* Um sistema de distribuição de energia inteligente (smart grid) precisa de informações de consumo dos usuários em tempo real, para fins de controle e alocação de carga do sistema. No entanto, o consumo individual de cada residência é criptografado, de modo a evitar que os dados sejam vazados e permitam a alguém inferir o padrão de consumo naquela residência.
* Um sistema de voto eletrônico é um exemplo clássico de aplicação da criptografia homomórfica. A computação no domínio criptográfico permite ao governo apurar o resultado final de uma eleição, mas o voto de cada cidadão não é revelado, preservando o direito de voto secreto de cada um.

A seguir, temos mais um vídeo que constroi uma analogia entre a criptografia homomórfica e o mundo real, para firmar os conhecimentos.

[![VideoAula](https://img.youtube.com/vi/BylWT5gsgfM/0.jpg)](https://www.youtube.com/watch?v=BylWT5gsgfM "Clique para assistir um video da NSF sobre o tema!")

**IMPORTANTE**: se você prestou atenção ao vídeo, notou que ele usa o termo *fully homomorphic encryptation*, ou criptografia completamente homomórfica. Esse termo é usado para definir um criptossistema capaz de construir todas as operações computacionais elementares (no mínimo, adição e multiplicação, ou uma porta lógica NAND). Por sua vez, o termo *semi-homomorphic encryptation*, ou criptografia parcialmente homomórfica, refere-se a criptossistemas que disponibilizam apenas uma dessas operações elementares (e.g., apenas a adição). 

Um dos grandes desafios à criptografia completamente homomórfica é o custo computacional. Os criptossistemas atuais pertencentes a essa classe são extremamente demorados quando executados em um computador, o que reduz significativamente sua aplicação prática. Embora vários avanços tenham sido feitos nos últimos anos neste campo, o mesmo constitui ainda um dos principais desafios para os criptólogos. Por outro lado, a criptografia parcialmente homomórfica possui um custo computacional bastante razoável, sendo assim mais aplicável a problemas práticos no mundo atual.

## Atividade prática em Python - Criptossistema Paillier

Passaremos agora a algumas atividades práticas, usando uma implementação do criptossistema [Paillier](https://en.wikipedia.org/wiki/Paillier_cryptosystem). O Paillier é um criptossistema parcialmente homomórfico, que possibilita operações de soma no domínio criptográfico. Sua implementação é simples, o que faz com que o mesmo esteja disponível em diferentes bibliotecas de programas que disponibilizam algoritmos criptográficos. É o caso do Python [PHE](https://python-paillier.readthedocs.io/en/stable/phe.html), que utilizaremos nessa prática de laboratório.

### 1 - Crie um par de chaves criptográficas

Para trabalhar com o Python PHE, precisamos criar um par de chaves criptográficas. Isso pode ser feito usando o programa [keygen.py](keygen.py). Abra o programa em uma outra aba do seu navegador e verifique como ele trabalha. Observe os parâmetros que devem ser informados. Em seu funcionamento, esse programa basicamente faz uso de um objeto *paillier* para criar o par de chaves, que serão salvos em disco em arquivos de extensão *.pub* e *.priv*.

Você pode executar esse programa conforme o comando abaixo:

```console
$ python3 keygen.py minhachave 512
```
Verifique se os arquivos de chaves foram gerados corretamente. Você pode inclusive usar o comando *hexdump* para verificar o conteúdo, uma vez que o arquivo é salvo em modo binário.

### 2 - Criptografe alguns números usando sua chave pública

Agora, use o programa [simple-crypt.py](simple-crypt.py) para encriptar alguns números que usaremos depois para realizar a computação no domínio criptográfico. Os números devem ser encriptados usando a chave pública gerada anterioremente. Abra o programa em uma outra aba do seu navegador e verifique como ele trabalha. Observe os parâmetros que devem ser informados. Você poderá executar o programa conforme sintaxe a seguir:

```console
$ python3 simple-crypt.py minhachave.pub 123
```
Note que o valor encriptado é um número muito maior do que o valor em claro. Na prática, o tamanho do número encriptado é definido pelo tamanho de chave escolhido (no nosso exemplo, 512 bits). Você deve criptografar pelo menos 3 números de modo a gerar entradas para o próximo passo desse tutorial.

### 3 - Some dois números encriptados

Usando as saidas criptografadas geradas no passo anterior, iremos chamar o programa [simple-sum.py](simple-sum.py) para realizar a soma de dois números encriptados. Assim como fizemos anteriormente, abra o programa em uma outra aba do seu navegador e verifique como ele trabalha. Observe os parâmetros que devem ser informados. Após entender seu funcionamento, execute o mesmo com a sintaxe a seguir, substituindo os parâmetros *dado A* e *dado B* por números criptografados gerados no passo anterior.

```console
$ python3 simple-sum.py minhachave.pub <dado-A> <dado-B>
```

### 4 - Descriptografe o resultado da soma

Usando a saída criptografada gerada no passo anterior, descriptografe a soma realizada usando o programa [simple-decrypt.py](simple-decrypt.py). Desta vez, é necessário se utilizar a **chave privada** para obter o valor referente à soma. Novamente, abra o programa em uma outra aba do seu navegador e verifique como ele funciona. Observe os parâmetros que devem ser informados. Em seguida, execute o comando conforme indicado a seguir:

```console
$ python3 simple-decrypt.py minhachave.priv <dado criptografado>
```
Se você executou todos os procedimentos de forma correta, vai observar que a soma obtida no domínio criptografo é exatamente a mesma daquela que seria obtida somando-se os números em claro.

### Questão 3: Repita os procedimentos práticos de criação de chaves, criptografia, soma e descriptografia realizado anteriormente, mas agora com chaves de 1024 e 2048 bits. Você pode dizer o que acontece gradualmente à medida que se aumenta o tamanho da chave? Acrescente os logs de tela de seu experimento para sustentar suas observações.

## Atividade prática em Python - Votação eletrônica

Como vimos anteriormente, o uso de criptografia homomórfica para proteger a privacidade de individuos em uma votação eletrônica constitui uma das aplicações mais interessantes para essa tecnologia. 

O uso de sistemas de voto eletrônico (*e-voting*) é um tema bastante controverso. Mesmo especialistas divergem sobre a forma mais segura e eficiente de se implementar esse sistema. Para ter uma ideia dos problemas envolvidos, veja esse vídeo, que entrevista um pesquisador da Unicamp a respeito do tema.

[![VideoAula](https://img.youtube.com/vi/xATaNCsre9Q/0.jpg)](https://www.youtube.com/watch?v=xATaNCsre9Q "Clique para assistir essa entrevista com o Prof. Diego Aranha")

Agora que sabemos como o tema é controverso, vamos executar uma prática de laboratório em que simulamos um sistema de voto eletrônico baseado em criptografia homomórfica. Novamente, utilizaremos o criptossistema Paillier para implementar as diretivas criptográficas necessárias.

O sistema de voto eletrônico implementado funciona um pouco diferente de uma eleição convencional. Ao invés de escolher um único candidato de uma lista, você deve atribuir uma nota de 0 a 10 para cada um dos candidatos que constam na lista. Ao final, o sistema de voto eletrônico soma as notas atribuídas. O vencedor da votação é, consequentemente, aquele que obtiver a maior soma das notas. Assim, vamos à prática!

### 1. Preencha sua cédula de voto eletrônico
O primeiro passo em nossa prática é preencher a cédula eleitoral com os votos (notas) atribuidos a cada candidato. Para isso, é preciso ter um par de chaves criptográficas. Para isso, você pode usar o mesmo par de chaves criado para o exercício anterior, ou usar o programa [keygen.py](keygen.py) para criar novas chaves.

De posse do seu par de chaves, dê uma analisada no programa [vote-ballot.py](vote-ballot.py). Esse programa é mais complexo do que os programas usados anteriormente, mas ele basicamente utiliza uma estrutura de dicionário do Python (*dict*) para criar uma cédula eleitoral. Em seguida, um *loop* é criado para se preencher o *dict* com os votos de cada candidato. Por fim, a estrutura completa é salva em arquivo, usando-se a classe *pickle*.

Para invocar o programa, execute:

```console
$ python3 vote-ballot.py minhachave.pub meuvoto.vot
```

Esse comando vai usar a chave pública *minhachave.pub* para criptografar seu voto e salvará a célula encriptada no arquivo *meuvoto.vot*.

Feito esse comando, repita o processo para pelo menos 3 outros votos, salvando-os em arquivos diferentes. Ao final, você deverá ter 3 arquivos, cada um contendo uma cédula eleitoral com diferentes votos.

### 2. Some os votos criptografados

Agora que temos pelo menos 3 arquivos contendo cédulas de votos criptografados, vamos usar o programa [vote-sum.py](vote-sum.py) para somar os votos individuais em cada arquivo, e salvar em uma cédula que totaliza esses votos. Examine o programa [vote-sum.py](vote-sum.py) antes de executá-lo, pra entender como ele funciona. Note que o programa é capaz de processar uma lista de arquivos de cédula, que são informados como argumentos ao final do programa. Assim, para realizar a soma dos votos, utilize o programa com a sintaxe a seguir:

```console
$ python3 vote-sum.py minhachave.pub total.vot voto1.vot voto2.vot voto3.vot
```

Esse comando vai usar a chave pública *minhachave.pub* para salvar no arquivo *total.vot* a soma dos votos contidos nos arquivos *voto1.vot*, *voto2.vot* e *voto3.vot*. Modifique a sintaxe do programa de acordo com o nome que você deu para os arquivos de voto que você criou no passo anterior.

### 3. Revele o voto contido numa cédula

Agora que temos uma cédula que contém a soma dos votos individuais, podemos revelar o resultado da votação descriptografando essa cédula. Para isso usamos o programa [vote-reveal.py](vote-reveal.py). Abra o programa e, tal como fizemos com os outros, examine seu código. Você verá que ele é bastante simples comparado com os demais, uma vez que simplesmente abre a cédula criptografada e descriptografa cada uma de suas posições, usando a chave privada respectiva. Para executar esse comando, utilize a seguinte sintaxe:

```console
$ python3 vote-reveal.py minhachave.priv total.vot
```

Note que esse programa também pode ser usado para descriptografar os votos individuais, uma vez que dispomos da chave privada. Para preservar os votos individuais, portanto, seria necessário um protocolo de segurança que garanta que quem tem os votos individuais não possui a chave privada, e vice-versa.

### Questão 4: Copie os logs de tela gerados na execução dos passos 1, 2 e 3 como resposta a essa questão, evidenciando que você executou as atividades.

## E finalmente, a escolha do seu Professor Favorito (Serginho)!
Por fim, faremos a prática esperada por todos! Usando os programas usados nas etapas anteriores, cada aluno vai gerar uma cédula de votação criptografada. Os votos individuais serão somados também de forma criptografada, de modo a preservar o sigilo dos mesmos. Por fim, o arquivo contendo a soma dos votos será disponibilizado ao professor, para que ele revele o resultado da votação.

Para manter esse procedimento de forma segura, usaremos o seguinte protocolo:

1. O professor gera um par de chaves e disponibiliza a chave pública aos alunos.
2. Os alunos escolhem dois representantes que ficam responsáveis por coletar os votos individuais e e somar esses votos.
3. Cada aluno gera um arquivo de voto criptografado, usando o programa [vote-ballot.py](vote-ballot.py).
4. Os alunos representantes coletam os votos individuais e os somam, usando o programa [vote-sum.py](vote-sum.py). O arquivo de totalização é entregue ao professor.
5. O professor revela o resultado da votação usando o programa [vote-reveal.py](vote-reveal.py).

### Questão 5: Existem diversas falhas de segurança no nosso protocolo de votação. Embora ele garanta a privacidade, os alunos podem facilmente burlar o resultado da eleição. Aponte pelo menos uma possível forma de fazer isso, e apresente em seguida uma contramedida que poderia ser usada para prevenir esse ataque.