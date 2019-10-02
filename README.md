# Criptografia Aplicada
Este repositório foi criado para manter tutoriais com exercícios e práticas de laboratório associados à disciplina de Criptografia Aplicada - Curso Técnico de Segurança da Informação - Inmetro 2019.
Ele é mantido pelos seguintes professores:
* Sérgio Câmara (smcamara@inmetro.gov.br);
* Wilson Melo Jr. (wsjunior@inmetro.gov.br).

## Aulas

- 5 e 7/Ago: Introdução ao curso e Procedimento GPG
- 12/Ago: [Esteganografia](https://github.com/wsmelojr/aulacripto/tree/master/esteganografia)
- 19 e 21/Ago: [Pretty Good Privacy (PGP)](https://github.com/wsmelojr/aulacripto/tree/master/pgp)
- 26 e 28/Ago: [Blockchain](https://github.com/wsmelojr/aulacripto/tree/master/blockchain)
- 2 e 4/Set: [Secure Socker Layer (SSL) - Parte 1](https://github.com/wsmelojr/aulacripto/tree/master/ssl)
- 9 e 11/Set: [Hyperledger Fabric](hyperledger)
- 16 e 18/Set: [Secure Socker Layer (SSL) - Parte 2](https://github.com/wsmelojr/aulacripto/tree/master/ssl2)
- 24 e 25/Set: [Criaçao de chaincodes](hyperledger/fabnotas)
- 30/Set e 2/Out: [Secure Shell (SSH)](ssh)

## Trabalho final

Os dias das apresentações e os temas foram definidos como segue:

- 11/Nov   
  - Blockchain (Hellen e Pedro)  
  - Token de Banco (Ícaro)
  - Esteganografia (Hugo)

- 13/Nov
  - Geração de números aleatórios (Mileny e Marlon)
  - Relógio Eletrônico de Ponto (Marcus)
  - Blockchain (Humberto e Marcos)

- 18/Nov
  - Site seguro com Certificado Digital (Diego)
  - _A definir_ (Hitallo)
  - Anonimato (Charlles)

### Entrega

O _deadline_ para entrega dos trabalhos escritos é dia **22/11/2019**.

Devem seguir o template da SBC:

Entregar em formato _.pdf_.

### Apresentação

As apresentações devem ter duração de 15 a 20 minutos, com 10 minutos reservados para perguntas.

### Requisitos dos trabalhos:

- Token de Banco
  - Implementar 2 programas: um que simula o servidor e o outro que simula o token (linguagem livre)
  - Implementar um dos algoritmos de One-Time Password (OTP), baseado em evento ou baseado em tempo.
  - Ao inserir o valor mostrado no token no programa do servidor, este deve conferir se o valor está correto.

- Geração de números aleatórios
  - Implementar um programa embarcado em Arduíno para gerar números aleatórios (utilizar dados de sensores, etc)
  - O programa deve dar uma saída de 256 bits (64 caacteres em hexa) de tempos em tempo (ex: 1 em 1 minuto)
  - Realizar testes de aletoriedade na saída (ex: testes do NIST)

- Relógio Eletrônico de Ponto
  - Implementar um programa (linguagem livre) que registra o ponto de usuários aleatórios.
  - O programa deve estar sincronizado com algum servidor NTP conhecido.
  - Para cada "registro de ponto simulado", o programa deve assinar digitalmente as informações do usuário a partir de um carimbo de tempo, _timestampping_, de uma Autoridade Certificadora de tempo (falsa).

- Site seguro com Certificado Digital
  - Criar um website considerado seguro.
  - Este site deve conter: Cadastro de Usuário, Tela de Login, funcionalidade para Recuperação de Senha.
  - O site deverá utilizar um certificado digital válido, reconhecido pelos navegadores (utilize o Let's Encrypt para geração do certificado).

- Anonimato
  - Explicar o funcionamento do Tor, Privoxy e Stunnel. Como seus mecanismos contribuem para o anonimato de um usuário da internet?
  - Acessar a DeepWeb, visitar sites que não sejam ilegais, documentar.
  - Implementar uma conexão via Stunnel (ex: https://charlesreid1.com/wiki/Stunnel/SSH)
