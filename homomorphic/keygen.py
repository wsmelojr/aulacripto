"""
    Votacao Homomorfica - Keygen
    ~~~~~~~~~
    Este modulo cria um par de chaves homomorficas para implementar
    o sistema-aula de votacao usando criptografia homomorfica.
    O criptosistema adotado eh o Paillier. A biblioteca phe do 
    python implementa esse criptosistema.

    O usuario executa esse comando informando o nome da chave (prefixo
    dos arquivos de chaves) e o tamanho da chave, que pode ser de 512, 
    1024 ou 2048 bits. O programa deve gerar os arquivos com as chaves
    publica e privada respectivas (extensoes pub e priv).
        
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
        print("Informe os parametros corretos:",sys.argv[0],"<nome> <tamanho>")
        exit(1)

    #obtem os parametros para facilitar a manipulacao
    key_name = sys.argv[1]
    key_size = sys.argv[2]

    #feedback para o usuario...
    print("Gerando um novo par de chaves...")

    #instancia um par de chaves invocando as funcoes da classe paillier
    pub_key, priv_key = paillier.generate_paillier_keypair(None,int(key_size))

    #formata os nomes de arquivos de chave, conforme parametro do usuario
    pub_key_file = key_name + ".pub"
    priv_key_file = key_name + ".priv"

    #escreve cada chave em seu respectivo arquivo
    pickle.dump(pub_key,open(pub_key_file, "wb"))
    pickle.dump(priv_key,open(priv_key_file, "wb"))

    #informa o usuario que, a principio, deu tudo certo
    print("As chaves foram salvas nos arquivos",pub_key_file,"e",priv_key_file)
