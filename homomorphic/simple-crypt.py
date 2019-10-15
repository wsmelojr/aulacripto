"""
    Votacao Homomorfica - SimpleCript
    ~~~~~~~~~
    Este modulo implementa uma criptografia simples, usando a chave
    publica informada e um numero como argumento. O programa criptografa
    o numero informado usando a chave respectiva e exibe o resultado.

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
        print("Informe os parametros corretos:",sys.argv[0],"<chave publica> <dado>")
        exit(1)

    #obtem os parametros para facilitar a manipulacao
    pub_key_file = sys.argv[1]
    data_num = sys.argv[2]

    #feedback para o usuario...
    print("Vou criptografar o numero",data_num,"usando a chave publica",pub_key_file)

    #le o arquivo de chave publica
    pub_key = pickle.load(open(pub_key_file, "rb"))
    
    #executa a criptografia 
    cript = pub_key.raw_encrypt(int(data_num))

    #informa o usuario que, a principio, deu tudo certo
    print("O valor encriptado eh",cript)