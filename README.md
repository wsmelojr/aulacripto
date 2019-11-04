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
- 7 e 14/Out: [Criaçao de chaincodes - Continuação](hyperledger/fabnotas)
- 9/Out: [Criptografia de volumes](veracrypt)
- 16/Out: Acompanhamento individual de trabalhos
- 21 e 23/Out: [Criptografia homomórfica](homomorphic)
- 30/Out: Acompanhamento individual de trabalhos
- 4/Nov: [Código de Autenticação de Mensagem](hmac)
- 6/Nov: Acompanhamento individual de trabalhos
- 8/Nov: Palestra sobre Forense Computacional

## Notas

 Acessa a planilha de avaliações [aqui](https://docs.google.com/spreadsheets/d/e/2PACX-1vSM_QDfPngjaZl_-zKzCqY3Q7y8xJYSfIeSJbVkeUWu3qHRTCNK6LFNlFoho0vhvILZkKAUph-Xspad/pubhtml).

## Trabalho final

Os dias das apresentações e os temas foram definidos como segue:

- 18/Nov (8h)
  - Hyperledger Composer (Hellen e Pedro)
  - Token de Banco (Ícaro)
  - Esteganografia (Mileny e Hugo)

- 18/Nov (10h)
  - Site seguro com Certificado Digital (Diego)
  - Anonimato (Charlles e Hitallo)

- 27/Nov (10h)
  - Relógio Eletrônico de Ponto (Marcus e Marlon)
  - Serviço orderer usando Kafka (Humberto e Marcos)


### Entrega

O _deadline_ para entrega dos trabalhos escritos é dia **27/11/2019**.

Devem seguir o [Template da SBC](https://www.sbc.org.br/documentos-da-sbc/summary/169-templates-para-artigos-e-capitulos-de-livros/878-modelosparapublicaodeartigos). 

Entregar em formato _.PDF_.

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

- Hyperledger Composer
  - Explicar o funcionamento da ferramenta Hyperledger Composer.
  - Descrever a aplicação implementada.
  - Mostrar a funcionalidade da aplicação em um exemplo prático.

- Serviço orderer usando Kafka
  - Explicar o que é o Kafka, e que funcionalidades ele disponibiliza.
  - Descrever um modelo de rede blockchain usando Kafka.
  - Implementar um tutorial usando uma rede blockchain com Kafka como orderer.
