#/bin/bash

# Esse script precisa de dois argumentos, um com nome
# do arquivo para gerar o MAC, e outro com a senha MAC
echo "$(cat $1)"$2 | sha256sum
