"""
    Votacao Homomorfica - SimpleDecript
    ~~~~~~~~~
    Este modulo decriptografa um valor (usualmente um inteiro longo) encriptado
    previamente com uma chave publica, usando a chave privada respectiva.

    :copyright: © 2019 by Wilson Melo Jr.
"""
import sys
sys.path.insert(0, "..")
import math
import phe.encoding
from phe import paillier
import pickle

if __name__ == "__main__":
    #testa se os argumentos foram informados corretamente
    if len(sys.argv) != 3:
        print("Informe os parametros corretos:",sys.argv[0],"<chave privada> <cedula criptografada>")
        exit(1)

    #obtem os parametros para facilitar a manipulacao
    priv_key_file = sys.argv[1]
    ballot_file = sys.argv[2]

    #feedback para o usuario...
    print("Vou decriptografar a cedula",ballot_file,"usando a chave privada",priv_key_file)

    #le o arquivo de chave publica
    priv_key = pickle.load(open(priv_key_file, "rb"))

    #le o arquivo referente a cedula de votos
    meuvoto = pickle.load(open(ballot_file, "rb"))

    #percorre os items em meuvoto, decriptografando cada um deles
    for x, y in meuvoto.items():
        #decriptografa o voto do item x
        meuvoto[x] = priv_key.raw_decrypt(int(y))

    #informa o usuario que, a principio, deu tudo certo
    print("O resultado da votacao eh:\n",meuvoto)