"""
    Votacao Homomorfica - SimpleSum
    ~~~~~~~~~
    Este modulo implementa a soma de dois parametros criptografados, 
    retornando o resultado também de forma criptografada.

    O usuario informa a chave publica a ser utilizada na operacao, e
    os dois numeros criptografados.

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
    if len(sys.argv) != 4:
        print("Informe os parametros corretos:",sys.argv[0],"<chave publica> <dado-A> <dado-B>")
        exit(1)

    #obtem os parametros para facilitar a manipulacao
    pub_key_file = sys.argv[1]
    data_A = sys.argv[2]
    data_B = sys.argv[3]

    #feedback para o usuario...
    print("Vou somar os numeros criptografados informados como A e B usando a chave publica",pub_key_file)

    #le o arquivo de chave publica
    pub_key = pickle.load(open(pub_key_file, "rb"))
    
    #cria um objeto EncryptedNumber para realizar operacoes aritmeticas
    encobj = paillier.EncryptedNumber(pub_key,data_A)

    #usa objeto criado para executar a soma dos dois numeros criptografados
    cript = encobj._raw_add(int(data_A),int(data_B))

    #informa o usuario que, a principio, deu tudo certo
    print("O valor encriptado eh",cript)