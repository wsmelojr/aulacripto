*Data das aula: 9/10* <br>
*Prof. Sérgio (smcamara@inmetro.gov.br)*


# Criptografia de volumes


# Laboratório

Instale o Veracrypt, utilize este arquivo de instalação:
https://launchpad.net/veracrypt/trunk/1.24/+download/veracrypt-1.24-Ubuntu-18.04-amd64.deb

`$ sudo dpkg -i veracrypt-1.24-Ubuntu-18.04-amd64.deb`

## Missão 1

1. Crie um volume encriptado de 600 MB (algoritmo AES + Twofish) por senha.
2. Monte o volume e copie algum arquivo para dentro dele.

## Missão 2

1. Gere e salve um arquivo de chave (Tools --> Keyfile Generator).
2. Crie um volume encriptado que utilize a keyfile em vez de senha.

## Missao 3

1. Crie um volume escondido (hidden volume) por senha.
2. Coloque arquivos no Volume externo (outer) e no escondido (hidden).
3. Explique a propriedade de _Negação plausível_ (_plausible deniability_) que o volume escondido adiciona ao Veracrypt.


## Missão 4

1. Instale o gerenciador de senhas KeePassXC (https://keepassxc.org/download/#linux).
2. Crie um novo database de senhas (.kdbx) com um senha forte.
3. Crie um novo registro de senha para salvar uma das senhas de um volume encriptado criado anteriormente.







## Links interessantes:
https://www.osradar.com/install-veracrypt-on-ubuntu-18-04/ <br>
