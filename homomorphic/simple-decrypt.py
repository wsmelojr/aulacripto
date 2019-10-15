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
        print("Informe os parametros corretos:",sys.argv[0],"<chave privada> <dado criptografado>")
        exit(1)

    #obtem os parametros para facilitar a manipulacao
    priv_key_file = sys.argv[1]
    data_cript = sys.argv[2]

    #feedback para o usuario...
    print("Vou decriptografar o numero",data_cript,"usando a chave privada",priv_key_file)

    #le o arquivo de chave publica
    priv_key = pickle.load(open(priv_key_file, "rb"))
    
    #executa a criptografia 
    decrypt = priv_key.raw_decrypt(int(data_cript))

    #informa o usuario que, a principio, deu tudo certo
    print("O valor decriptado eh",decrypt)