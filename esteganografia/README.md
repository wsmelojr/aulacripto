# Esteganografia

Esteganografia é a ciência de ocultar uma informação dentro de outra informação. Ela pode ser descrita como uma forma de **segurança por obscuridade**. Neste caso, a segurança não está essencialmente baseada na dificuldade de "quebra" de um algoritmo de criptografia, e sim no fato de que apenas o correto destinatário da mensagem sabe que aquela informação esconde algum segredo.

O vídeo a seguir apresenta de forma sucinta os principais conceitos associados à esteganografia.

[![VideoAula](https://img.youtube.com/vi/8FO3iqmLFN8/0.jpg)](https://www.youtube.com/watch?v=8FO3iqmLFN8 "Clique para assistir um video introdutório!")

### Questão 1: Explique com suas palavras a diferença entre a esteganografia e a forma como usualmente a criptografia é utilizada.

## Aplicações para a esteganografia

Basicamente, as aplicações de estaganografia envolvem casos onde uma informação é transmitida de uma parte para outra de maneira furtiva em um canal amplamente acessível. Para tanto, diferentes tipos de mídia podem ser utilizados (imagens, áudios e vídeos, por exemplo).

Ferramentas que fazem uso de imagens ou vídeos para incluir mensagens ocultas geralmente chamam mais a atenção das pessoas, pelo impacto quase cinematográfico que isso sobre elas. Entretanto, qualquer arquivo digital pode ser usado em aplicações de esteganografia.

Veja o vídeo abaixo, que é parte de uma série de ficção. Ele ilustra um caso onde criminosos enviam mensagens por meio de vídeos esteganografados.

[![VideoAula](https://img.youtube.com/vi/T4tG8_MFBsQ/0.jpg)](https://www.youtube.com/watch?v=T4tG8_MFBsQ "Clique para assistir um video que ilustra um possível uso da esteganografia!")

### Questão 2: Descreva um exemplo onde uma esteganografia poderia ser usada para ocultar um crime. Explique como você acredita que um investigador forense deveria atuar nessa situação.

A aplicação da esteganografia, todavia, não está restrita a atividades criminosas. Esta técnica pode ser sim explorada como mecanismo de proteção de sistemas. Uma aplicação muito interessante é a inserção de **marca d'agua** em um software usando esteganografia. Com isso, uma espécie de "registro de propriedade secreto" é inserido dentro do software, e pode depois ser utilizado para expor uma tentativa de pirataria de um produto.

**DICA DO DIA**: A tese de doutorado da Profa. Lucila Bento aborda a inserção de marca d'agua em produtos de software. Conversem com ela sobre o tema!

## Como utilizar a esteganografia

Existem diferentes algoritmos para se obter a ocultação de uma informação usando esteganografia. Em alguns casos, esses algoritmos podem inclusive combinar outras técnicas de criptografia convencionais, como por exemplo um algorítmo de chave simétrica.

Nessa aula de laboratório, faremos uso de um aplicativo simples, todavia bastante funcional, chamado [*steghide*](http://steghide.sourceforge.net/documentation/manpage.php).

### Instalando  e testando o *steghide*
O primeiro passo é atualizar o repositório de programas do Ubuntu com o seguinte comando:

    $ sudo apt-get update	

Em seguida instale o programa *steghide* com o comando:

    $ sudo apt-get install steghide

Teste se a instalação foi bem sucedida, executando o comando com um parâmetro de *help*:

    $ steghide --help

Se você consegue ver as informações sobre o comando (versão, parâmetros, etc), então ele está instalado e pronto pra ser usado.

### Adicionando uma mensagem secreta a uma imagem

O *steghide* funciona com arquivos de imagem e áudio. Atualmente o *steghide* suporta apenas imagens digitais em formato JPEG e BMP e arquivos de áudio em WAV ou AU. 

Para verificar seu funcionamento, vamos escolher um arquivo qualquer chamado *foto.jpg* (pode ser uma foto sua, ou uma paisagem, ou qualquer outra imagem que você goste, desde que seja em formato JPEG). Para realizar a prática deste laboratório de forma correta, crie uma cópia do arquivo de *foto.jpg* com o seguinte nome: *foto-original.jpg*. 

**IMPORTANTE**: se você está realizando esse tutorial com uma cópia do reposítorio em sua máquina, ele já provê arquivos exemplo para *foto.jpg* e "foto-original.jpg". Você pode sobreescrever esses dois arquivos, sem problemas!

Com o arquivo de imagem em mãos, é necessário criar a mensagem secreta em um arquivo texto. Para isso, crie um arquivo chamado *mensagem.txt* e digite nele sua mensagem secreta (uma frase curta, por exemplo).

Em seguida, procedemos executando o seguinte comando:

    $ steghide embed -cf foto.jpg -ef mensagem.txt

O *steghide* soicitará uma *passphrase* (correspondente a uma senha ou código de acesso) para proteção da esteganografia. Você deve digitar a *passphrase* e confirmá-la em seguida, conforme o seguinte prompt:

    Enter passphrase: 
    Re-Enter passphrase:
    embedding "mensagem.txt" in "foto.jpg"... done

O processo reverso, ou seja, a extração da mensagem secreta, pode ser feito usando-se o seguinte comando:

    $ steghide extract -sf foto.jpg

Observe que, ao executar esse comando, o *steghide* solicitará a *passphrase* usada no momento da criação da esteganografia. Sem a *passphrase* correta, a extração da esteganografia não vai funcionar!

Faça uma comparação visual entre os dois arquivos. Abra ambas as imagens lado a lado, e veja se você consegue perceber alguma diferença entre elas.

Em seguida, faça a verificação dos *hashes* de ambos os arquivos usando o algoritmo SHA1:

    $ sha1sum foto.jpg
    $ sha1sum foto-original.jpg

Os codigos *hash* mostram que os arquivos são diferentes. A esteganografia está lá, mas ela não é perceptível na imagem visual.
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTE0MjYzNjAxMjcsMTg2OTg2OTg1OSwyOT
cwNDY1MSwxODY5ODY5ODU5LDEwNzA0ODI5NDgsLTk5NjU1MTcz
MCwzNDMzNjAzODAsLTUxNTU0MTIwMCwxNzIxODk2MzYxLDExMz
EwMjI1MTMsLTE1OTY3NzA0MjUsMTU4MjYwODAyNSwtMTg3MDQ0
NTU1LC0zNTUzMjI1NjMsLTExMzg4MDgxNSw1ODc0MjA1NjEsNT
c3MjcxNTY0LDM1Nzc0OTYsMTk4MDM3MjI5NiwzMzAwNjk2NzJd
fQ==
-->