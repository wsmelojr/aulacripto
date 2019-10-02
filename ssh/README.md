*Data das aulas: 30/9 e 2/10* <br>
*Prof. Sérgio (smcamara@inmetro.gov.br)*


# Secure Shell (SSH)

SSH is an abbreviation for Secure Shell which is a network protocol that is used to establish a secure connection between client and server. SSH can allow users especially system administrators to access computers remotely through a secure channel on the top of an unsecured network.

SSH supply users with powerful encryption and authentication methods to communicate data between computers that are connecting over an unsecured network. SSH is commonly used by system administrators for connecting to remote machines, execute commands/scripts, handle the administrative tasks, securely transfer files from one machine to another and overall manage systems and applications remotely.

Also, you can use SSH to function as or act as a proxy server and redirect your browsing traffic to an encrypted SSH tunnel. This will prevent users on public networks from seeing your browsing history.

# Instruções

- O laboratório está dividido em 2 aulas. Para cada laboratório, o aluno deverá cumprir as missões propostas. Ao final de cada **Missão**, tire um *print* da tela indicando que ela foi cumprida.
- O *print* de cada **Missão** e as respostas das **Questões** deverão ser compiladas em um relatório (*.pdf*).

Sobre a entrega das respostas:
- Envie para o email do professor (smcamara [at] inmetro [ponto] gov [ponto] br).
- Assunto do email: [AulaCripto2019] SSH
- **Deadline: 7/10/2019**
- Incluir a Missão 3 da aula de SSL - Parte 2.

# Laboratório

Instale o pacote do servidor OpenSSH:

`$ sudo apt install openssh-server`

Visualize o status do serviço de ssh:

`$ sudo systemctl status ssh`


Edite o arquivo de configuração sshd_config:

`$ sudo nano /etc/ssh/sshd_config`

Inclua as linhas:

`AllowUsers nome_do_usuario` <br>
`PermitRootLogin No`


Restarte o serviço de SSH:

`$ sudo systemctl restart ssh`




## Missão 1

1. Adicione um novo usuário no sistema para acessar remotamente via ssh.

`$ sudo adduser nome_do_usuario`

2. Acesse o seu servidor SSH da sua outra máquina

`$ ssh nome_do_usuario@endereço_ip`

3. Verifique o diálogo de login, qual o fingerprint da chave pública do servidor de SSH que você recebeu?


4. Execute os seguintes comandos para verificar o seu usuário e o hostname:

`$ whoami` <br>
`$ hostname`

5. Feche a conexão

`$ exit`

6. Pelo terminal do servidor ssh, verifique os logs do serviço de SSH, localize os registros que relatam o login e logout do usuário criado. Dê um print no log.

`$ grep 'sshd' /var/log/auth.log`

## Missão 2

Agora, vamos configurar o acesso do usuário para ser feito através de chave-pública, ao invés de senha.

1. Gere um novo par de chaves RSA

`$ ssh-keygen -b 4096 -t rsa`

Aperte <kbd>Enter</kbd> para utilizar os nomes padrão _id_rsa_ e _id_rsa.pub_ no diretório _/home/seu_usuário/.ssh_

Não entre com o _passphrase_ (deixe em branco).

2. Copie sua chave pública para o servidor

`$ ssh-copy-id nome_do_usuário@endereço_ip`

3. Conecte-se ao servidor ssh.

`$ ssh nome_do_usuario@endereço_ip`

4. Verifique se a pasta .ssh foi criada e se dentro dela há um arquivo chamado _authorized_keys_ que contém a sua chave pública. Dê um print no conteúdo do arquivo _authorized_keys_.

`$ cat .ssh/authorized_keys`


## Missão 3

1. Escolha/baixe uma figura qualquer na internet.

2. Renomeie esse figura com o seu nome e copie essa figura para ao seu servidor ssh.

`$ scp nome.jpg nome_do_usuario@endereço_ip:/home/nome_do_usuario`

3. Conecte-se ao seu servidor ssh, e a partir dele, copie a figura copiada para o servidor ssh do professor.

`$ scp nome.jpg aulacripto@endereço_ip:/home/aulacripto`

4. Conecte-se a esse outro servidor ssh para verificar se a figura foi copiada corretamente. Dê um print no conteúdo do diretório remoto. Depois, feche a conexão.

5. Do terminal da sua máquina local (Ubuntu), copie para a sua máquina local a figura _/home/aulacripto/missao3.jpg_ que está no servidor do professor. Dê um print na figura.

`$ scp aulacripto@:/home/aulacripto/missao3.jpg /home/usuario_local`











### Extras:

Para parar o servidor ssh:

`$ sudo systemctl stop ssh`

Para desativar o serviço SSH para que ele não inicie durante o boot do sistema:

`$ sudo systemctl disable ssh`



## Links interessantes:
https://www.openssh.com/ <br>
https://charlesreid1.com/wiki/Stunnel/SSH <br>
https://www.itzgeek.com/how-tos/linux/ubuntu-how-tos/how-to-enable-ssh-on-ubuntu-18-04-linux-mint-19-debian-9.html <br>
https://www.itzgeek.com/how-tos/linux/ubuntu-how-tos/how-to-enable-ssh-on-ubuntu-18-04-linux-mint-19-debian-9.html <br>
https://linode.com/docs/security/authentication/use-public-key-authentication-with-ssh/ <br>
https://serverfault.com/questions/130482/how-to-check-sshd-log <br>
