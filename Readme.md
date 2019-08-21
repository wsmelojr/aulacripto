---


---

<p><em>Data das aulas: 19/8 e 21/8<br>
Prof. Sérgio (<a href="mailto:smcamara@inmetro.gov.br">smcamara@inmetro.gov.br</a>)</em></p>
<h1 id="pretty-good-privacy-pgp">Pretty Good Privacy (PGP)</h1>
<p><strong>Pretty Good Privacy</strong> (<strong>PGP</strong>), em português <strong>privacidade muito boa</strong>, é um software de criptografia que fornece autenticação e privacidade criptográfica para comunicação de dados. É frequentemente utilizado para assinar, encriptar e descriptografar textos, e-mails, arquivos, diretórios e partições inteiras de disco e para incrementar a segurança de comunicações via e-mail. Foi desenvolvido por <a href="https://pt.wikipedia.org/wiki/Phil_Zimmermann" title="Phil Zimmermann">Phil Zimmermann</a> em 1991.<a href="https://pt.wikipedia.org/wiki/Pretty_Good_Privacy">[1]</a></p>
<p>O PGP e softwares similares seguem o padrão <a href="https://pt.wikipedia.org/wiki/OpenPGP" title="OpenPGP">OpenPGP</a> (<a href="https://tools.ietf.org/html/rfc4880">RFC 4880</a>) para encriptar e decriptar dados.</p>
<h1 id="instruções">Instruções</h1>
<ul>
<li>O laboratório de PGP está dividido em 2 aulas. Para cada laboratório, o aluno deverá cumprir as missões propostas. Ao final de cada <strong>Missão</strong>, tire um <em>print</em> da tela indicando que ela foi cumprida.</li>
<li>O <em>print</em> de cada <strong>Missão</strong> e as respostas das <strong>Questões</strong> deverão ser compiladas em um relatório (<em>.pdf</em>).</li>
</ul>
<p>Sobre a entrega das respostas:</p>
<ul>
<li>Envie sua chave pública (.asc) e o pdf com as respostas assinado digitalmente para o email do professor (smcamara [at] gmail [ponto] com). Obs: Não encriptar o <em>.pdf</em>.</li>
<li><strong>Deadline: 26/8/2019</strong></li>
</ul>
<h1 id="laboratório">Laboratório</h1>
<p>Para realizar as missões descritas abaixo, primeiramente instale o software <em>Seahorse</em>.</p>
<pre><code>&gt; sudo apt install seahorse
</code></pre>
<h2 id="missão-1">Missão 1</h2>
<ol>
<li>Cada aluno deve gerar uma chave pública RSA 2048.</li>
<li>Revogue essa chave pública criada.</li>
</ol>
<h2 id="missão-2">Missão 2</h2>
<ol>
<li>
<p>Cada aluno gera sua chave pública RSA 4096.</p>
</li>
<li>
<p>Extrair sua chave pública para um arquivo (.asc).</p>
<pre><code>File -&gt; Export
Escolher "Armored PGP keys"
</code></pre>
</li>
<li>
<p>Subir sua chave pública para o seu perfil do <em>Github</em>.</p>
</li>
</ol>
<h2 id="missão-3">Missão 3</h2>
<ol>
<li>Subir a sua chave pública para os servidores de chave.</li>
</ol>
<blockquote>
<p>Caso a publicação de chave não funcionar pelo Seahorse, acessar os sites abaixo e publicar manualmente:<br>
<a href="https://keyserver.ubuntu.com/">https://keyserver.ubuntu.com/</a><br>
<a href="https://pgp.key-server.io/">https://pgp.key-server.io/</a></p>
</blockquote>
<ol start="2">
<li>Adquirir as chaves públicas de outros 5 alunos em sala. Importe-as para o chaveiro do <em>Seahorse</em>. Tire um <em>print</em> da tela do <em>Seahorse</em>.</li>
</ol>
<h2 id="missão-4">Missão 4</h2>
<p>A <em>Key ID</em> da chave pública do prof. Sérgio é:</p>
<pre><code>dreadful midsummer classic fascinate
</code></pre>
<ol>
<li>Qual é a <em>Key ID</em> em hexadecimal?</li>
<li>Ache a chave pública correta do professor no servidor de chaves <a href="https://keyserver.ubuntu.com/">https://keyserver.ubuntu.com/</a> e importe ao seu chaveiro. Dê <em>print</em>.</li>
<li>Cite uma aplicação prática para a PGP <em>word list</em>.</li>
</ol>
<h2 id="missão-5">Missão 5</h2>
<p>Ache na internet dois sites que utilizam arquivos PGP (.pgp ou .asc).</p>
<blockquote>
<p>Ex: indivíduo disponibilizando sua chave pública em seu site pessoal; área de download de um site que contém a assinatura digital dos arquivos para verificação, etc.<br>
Não vale tutoriais ou servidores de chave.</p>
</blockquote>
<ol>
<li>Cite o título e o endereço das páginas encontradas, e seus <em>printscreens</em>.</li>
</ol>
<h2 id="missão-6">Missão 6</h2>
<p>Instale o plugin do <em>Seahorse</em> para o gerenciador de arquivos <em>Nautilus</em>:</p>
<pre><code>&gt; sudo apt install seahorse-nautilus
</code></pre>
<p>Reinicie o <em>Nautilus</em> antes de usá-lo:</p>
<pre><code>&gt; nautilus -q
</code></pre>
<ol>
<li>
<p>Escolha um arquivo (uma mensagem de texto, uma foto, etc), e siga os próximas passos:</p>
<ul>
<li>Abra o gerenciador de arquivos <em>Nautilus</em>.</li>
<li>Clique com o botão direito sobre o arquivo escolhido.</li>
<li>Clique em “Encrypt…”.</li>
<li>Escolha o destinatário e assine digitalmente o arquivo com sua própria chave.</li>
<li>Envie o arquivo gerado (<em>.pgp</em>) para o email do destinatário.</li>
</ul>
</li>
<li>
<p>Decripte o mesmo arquivo <em>.pgp</em>:</p>
<ul>
<li>Clique com o botão direito sobre ele.</li>
<li>Clique em “Open With Other Application…”.</li>
<li>Escolha <em>Decrypt File</em>.</li>
<li>Salve o arquivo a ser decriptado, sem sobrescrever o arquivo original.</li>
</ul>
</li>
</ol>
<h2 id="missão-7">Missão 7</h2>
<p>Baixe o baixe o arquivo <strong>missao7.jpg.pgp</strong>.</p>
<ol>
<li>Decripte (passphrase: “inmetro”) e verifique a assinatura desse arquivo com a chave pública do professor (recuperada na Missão 4).</li>
</ol>
<h2 id="missão-8">Missão 8</h2>
<p>Instale o plugin <em>Mailvelope</em> (<a href="https://www.mailvelope.com/en/">https://www.mailvelope.com/en/</a>) no seu navegador.</p>
<ol>
<li>Importe suas chaves pública e privada para o chaveiro do <em>Mailvelope</em> (Utilize a funcionalidade do <em>Seahorse</em> de exportar a chave secreta já criada).</li>
<li>Escreva um novo email para algum outro aluno que você possua a chave pública.</li>
<li>Encripte e assine essa mensagem usando o <em>Mailvelope</em>, e envie-a.</li>
</ol>
<p>Tutorial para ajudar: <a href="https://www.youtube.com/watch?v=4ba0K-DhoGo">https://www.youtube.com/watch?v=4ba0K-DhoGo</a>  (Dá um like e ativa o sininho).</p>
<h1 id="questões">Questões</h1>
<ol>
<li>Quais são os cinco principais serviços fornecidos pelo PGP?</li>
<li>Qual é a utilidade de uma assinatura avulsa?</li>
<li>Por que o PGP gera uma assinatura antes de aplicar a compactação?</li>
<li>Como o PGP usa o conceito de confiança?</li>
</ol>
<h2 id="links-interessantes">Links interessantes:</h2>
<ul>
<li>Keybase - <a href="https://keybase.io/">https://keybase.io/</a></li>
<li>PGP word list - <a href="https://en.wikipedia.org/wiki/PGP_word_list">https://en.wikipedia.org/wiki/PGP_word_list</a></li>
<li>Public Key Fingerprint - <a href="https://en.wikipedia.org/wiki/Public_key_fingerprint">https://en.wikipedia.org/wiki/Public_key_fingerprint</a></li>
<li>Seahorse -   <a href="https://wiki.gnome.org/Apps/Seahorse">https://wiki.gnome.org/Apps/Seahorse</a></li>
<li>OpenPGP Best Practices - <a href="https://riseup.net/en/security/message-security/openpgp/best-practices">https://riseup.net/en/security/message-security/openpgp/best-practices</a></li>
<li>Encrypting and Decrypting Text with PGP - <a href="https://www.youtube.com/watch?v=sRmrvrM3y6o">https://www.youtube.com/watch?v=sRmrvrM3y6o</a></li>
</ul>

