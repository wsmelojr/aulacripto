"""
    Votacao Homomorfica - VoteSum
    ~~~~~~~~~
    Este modulo faz a soma dos votos encriptados no dominio
    homomorfico e armazena o total em uma nova cedula, igualmente
    encriptada.

    O usuario deve informar a chave publica, o nome do arquivo novo
    que sera criado, e a lista de arquivos de votos a serem somados
    (pelo menos 2).

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
    if len(sys.argv) < 5:
        print("Informe os parametros corretos:",sys.argv[0],"<chave publica> <arquivo apuracao> <cedula 1> ... <cedula n>")
        exit(1)

    #obtem os parametros para facilitar a manipulacao
    pub_key_file = sys.argv[1]
    result_file = sys.argv[2]
    cedula = sys.argv[3]

    #le o arquivo de chave publica
    pub_key = pickle.load(open(pub_key_file, "rb"))

    #cria um objeto EncryptedNumber para realizar operacoes de soma
    encobj = paillier.EncryptedNumber(pub_key,0)

    #le o arquivo referente a primeira cedula
    total = pickle.load(open(cedula, "rb"))

    #percorre cada um dos arquivos de cedula informados
    #ele comeca em 4 porque eh a posicao esperada para a 
    #segunda cedula, e temos que ter pelo menos 2 cedulas
    for n in range(4, len(sys.argv)):
        #le a proxima cedula
        cedula = sys.argv[n]

        #le o arquivo referente a primeira cedula
        prox = pickle.load(open(cedula, "rb"))

        #verifica se as cedulas tem o mesmo tamanho
        if len(total) != len(prox):
            #tamanhos diferentes, precisa parar a apuracao
            print("A cedula",cedula,"tem tamanho incompativel com as demais")

            #sai fora
            exit

        #cria um segundo loop para somar os votos de cada posicao da cedula
        for x, y in total.items():
            #faz a soma dos votos referentes ao item x, no dominio criptografico
            total[x] = encobj._raw_add(int(y),int(prox[x]))

    #feedback para o usuario...
    print("Vou criptografar a apuracao de votos no arquivo",result_file)

    #escreve a celula com voto criptografado em arquivo
    pickle.dump(total,open(result_file, "wb"))

    #informa o usuario que, a principio,c deu tudo certo
    print("A cedula criptografada foi gerada:\n",total)