"""
    Votacao Homomorfica - VoteBallot
    ~~~~~~~~~
    Este modulo cria uma cedula de votacao com os itens ponderados defindos
    no corpo do codigo.

    Uma vez preenchida a cedula, cada voto ponderado eh criptografado
    e o resultado eh armazenado em um arquivo com os respectivos valores.

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
        print("Informe os parametros corretos:",sys.argv[0],"<chave publica> <arquivo voto>")
        exit(1)

    #obtem os parametros para facilitar a manipulacao
    pub_key_file = sys.argv[1]
    ballot_file = sys.argv[2]

    #le o arquivo de chave publica
    pub_key = pickle.load(open(pub_key_file, "rb"))

    #feedback para o usuario...
    print("Vou criptografar seu voto e salvar no arquivo",ballot_file)

    #monta a cedula eleitoral com os strings padrao
    meuvoto = {
        "Flavinha": 0,
        "Lucila": 0,
        "Joao": 0,
        "Madruga": 0,
        "Rodrigo": 0,
        "Serginho": 0,
        "Wilson": 0,
        "Wladmir": 0
    }

    #percorre as entradas da cedula, solicitando o valor de voto para cada um dos items
    for x, y in meuvoto.items():
        #gera mensagem para pedir entrada de dados
        msg = "Atribua seu voto para " + x + " (0-10): "

        #faz o controle do limite de valores informados
        voto = -1
        while voto < 0 or voto > 10:
            #pede que o usuario informe o voto
            voto = int(input(msg))

        #criptografa o valor informado no voto
        meuvoto[x] = pub_key.raw_encrypt(voto)

    #escreve a celula com voto criptografado em arquivo
    pickle.dump(meuvoto,open(ballot_file, "wb"))

    #informa o usuario que, a principio, deu tudo certo
    print("A cedula criptografada foi gerada:\n",meuvoto)