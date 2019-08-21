*Data das aulas: 19/8 e 21/8
Prof. Sérgio (smcamara@inmetro.gov.br)*


# Pretty Good Privacy (PGP)
**Pretty Good Privacy** (**PGP**), em português **privacidade muito boa**, é um software de criptografia que fornece autenticação e privacidade criptográfica para comunicação de dados. É frequentemente utilizado para assinar, encriptar e descriptografar textos, e-mails, arquivos, diretórios e partições inteiras de disco e para incrementar a segurança de comunicações via e-mail. Foi desenvolvido por [Phil Zimmermann](https://pt.wikipedia.org/wiki/Phil_Zimmermann "Phil Zimmermann") em 1991.[[1]](https://pt.wikipedia.org/wiki/Pretty_Good_Privacy)

O PGP e softwares similares seguem o padrão [OpenPGP](https://pt.wikipedia.org/wiki/OpenPGP "OpenPGP") ([RFC 4880](https://tools.ietf.org/html/rfc4880)) para encriptar e decriptar dados.

# Instruções

- O laboratório de PGP está dividido em 2 aulas. Para cada laboratório, o aluno deverá cumprir as missões propostas. Ao final de cada **Missão**, tire um *print* da tela indicando que ela foi cumprida.
- O *print* de cada **Missão** e as respostas das **Questões** deverão ser compiladas em um relatório (*.pdf*).

Sobre a entrega das respostas:
- Envie sua chave pública (.asc) e o pdf com as respostas assinado digitalmente para o email do professor (smcamara [at] gmail [ponto] com). Obs: Não encriptar o *.pdf*.
- **Deadline: 26/8/2019**

# Laboratório


Para realizar as missões descritas abaixo, primeiramente instale o software *Seahorse*.

    > sudo apt install seahorse



## Missão 1

 1. Cada aluno deve gerar uma chave pública RSA 2048.     
 2. Revogue essa chave pública criada.

## Missão 2

 1. Cada aluno gera sua chave pública RSA 4096.
    
    
 2. Extrair sua chave pública para um arquivo (.asc).
   
        File -> Export
        Escolher "Armored PGP keys"
    
3. Subir sua chave pública para o seu perfil do *Github*.

## Missão 3
1. Subir a sua chave pública para os servidores de chave.
>Caso a publicação de chave não funcionar pelo Seahorse, acessar os sites abaixo e publicar manualmente:
https://keyserver.ubuntu.com/
https://pgp.key-server.io/

2. Adquirir as chaves públicas de outros 5 alunos em sala. Importe-as para o chaveiro do *Seahorse*. Tire um *print* da tela do *Seahorse*.


## Missão 4
A *Key ID* da chave pública do prof. Sérgio é:

    dreadful midsummer classic fascinate



1. Qual é a *Key ID* em hexadecimal?
2. Ache a chave pública correta do professor no servidor de chaves https://keyserver.ubuntu.com/ e importe ao seu chaveiro. Dê *print*.
3. Cite uma aplicação prática para a PGP *word list*.

## Missão 5
Ache na internet dois sites que utilizam arquivos PGP (.pgp ou .asc).
>Ex: indivíduo disponibilizando sua chave pública em seu site pessoal; área de download de um site que contém a assinatura digital dos arquivos para verificação, etc.
> Não vale tutoriais ou servidores de chave.

1. Cite o título e o endereço das páginas encontradas, e seus *printscreens*.



## Missão 6
Instale o plugin do *Seahorse* para o gerenciador de arquivos *Nautilus*:

    > sudo apt install seahorse-nautilus

Reinicie o *Nautilus* antes de usá-lo:

    > nautilus -q

1. Escolha um arquivo (uma mensagem de texto, uma foto, etc), e siga os próximas passos:

     - Abra o gerenciador de arquivos *Nautilus*.
     - Clique com o botão direito sobre o arquivo escolhido.
    - Clique em "Encrypt...".
    - Escolha o destinatário e assine digitalmente o arquivo com sua própria chave.
    - Envie o arquivo gerado (*.pgp*) para o email do destinatário.

2. Decripte o mesmo arquivo *.pgp*:

    - Clique com o botão direito sobre ele.
    - Clique em "Open With Other Application...".
    - Escolha *Decrypt File*.
    - Salve o arquivo a ser decriptado, sem sobrescrever o arquivo original.

## Missão 7
Baixe o baixe o arquivo **missao7.jpg.pgp**.
1. Decripte (passphrase: "inmetro") e verifique a assinatura desse arquivo com a chave pública do professor (recuperada na Missão 4).

## Missão 8
Instale o plugin *Mailvelope* (https://www.mailvelope.com/en/) no seu navegador.
1. Importe suas chaves pública e privada para o chaveiro do *Mailvelope* (Utilize a funcionalidade do *Seahorse* de exportar a chave secreta já criada).
2. Escreva um novo email para algum outro aluno que você possua a chave pública.
3. Encripte e assine essa mensagem usando o *Mailvelope*, e envie-a.


Tutorial para ajudar: https://www.youtube.com/watch?v=4ba0K-DhoGo  (Dá um like e ativa o sininho).

# Questões

 1. Quais são os cinco principais serviços fornecidos pelo PGP? 
 2. Qual é a utilidade de uma assinatura avulsa? 
 3. Por que o PGP gera uma assinatura antes de aplicar a compactação? 
 4. Como o PGP usa o conceito de confiança?


## Links interessantes:

 - Keybase - https://keybase.io/ 
 - PGP word list - https://en.wikipedia.org/wiki/PGP_word_list 
 - Public Key Fingerprint - https://en.wikipedia.org/wiki/Public_key_fingerprint 
  - Seahorse -   https://wiki.gnome.org/Apps/Seahorse 
   - OpenPGP Best Practices - https://riseup.net/en/security/message-security/openpgp/best-practices
   -  Encrypting and Decrypting Text with PGP - https://www.youtube.com/watch?v=sRmrvrM3y6o

