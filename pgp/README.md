<p><em>Data das aulas: 19/8 e 21/8<br>
Prof. Sérgio (<a href="mailto:smcamara@inmetro.gov.br">smcamara@inmetro.gov.br</a>)</em></p>
<h1 id="pretty-good-privacy-pgp">Pretty Good Privacy (PGP)</h1>
<p><strong>Pretty Good Privacy</strong> (<strong>PGP</strong>), em português <strong>privacidade muito boa</strong>, é um software de criptografia que fornece autenticação e privacidade criptográfica para comunicação de dados. É frequentemente utilizado para assinar, encriptar e descriptografar textos, e-mails, arquivos, diretórios e partições inteiras de disco e para incrementar a segurança de comunicações via e-mail. Foi desenvolvido por <a href="https://pt.wikipedia.org/wiki/Phil_Zimmermann" title="Phil Zimmermann">Phil Zimmermann</a> em 1991.<a href="https://pt.wikipedia.org/wiki/Pretty_Good_Privacy">[1]</a></p>
<p>O PGP e softwares similares seguem o padrão <a href="https://pt.wikipedia.org/wiki/OpenPGP" title="OpenPGP">OpenPGP</a> (<a href="https://tools.ietf.org/html/rfc4880">RFC 4880</a>) para encriptar e decriptar dados.</p>
<h1 id="instruções">Instruções</h1>
<p>O laboratório de PGP está dividido em 2 aulas. Para cada laboratório, o aluno deverá cumprir as missões propostas. Ao final de cada <strong>Missão</strong>, tire um <em>print</em> da tela indicando que ela foi cumprida.<br>
O <em>print</em> de cada <strong>Missão</strong> e as respostas das <strong>Questões</strong> deverão ser compiladas em um relatório (<em>.pdf</em>) para futura entrega. Mais detalhes, em breve.</p>
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
<li>Adquirir as chaves públicas de outros 5 alunos em sala.</li>
</ol>
<h2 id="missão-4">Missão 4</h2>
<p>O <em>fingerprint</em> da chave pública do prof. Sérgio é:</p>
<ol>
<li>Qual é o <em>fingerprint</em> em hexadecimal?</li>
<li>Ache a chave pública correta do professor no servidor de chaves <a href="https://keyserver.ubuntu.com/">https://keyserver.ubuntu.com/</a> e importe ao seu chaveiro.</li>
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
<h2 id="missão-6-...-coming-soon---2108">Missão 6 … [coming soon - 21/08]</h2>
<h2 id="missão-7-...-coming-soon---2108">Missão 7 … [coming soon - 21/08]</h2>
<h2 id="section">…</h2>
<h1 id="questões">Questões</h1>
<ol>
<li>Quais são os cinco principais serviços fornecidos pelo PGP?</li>
<li>Qual é a utilidade de uma assinatura avulsa?</li>
<li>Por que o PGP gera uma assinatura antes de aplicar a compactação?</li>
<li>Como o PGP usa o conceito de confiança?</li>
</ol>
<h2 id="links-interessantes">Links interessantes:</h2>
<p><a href="https://keybase.io/">https://keybase.io/</a><br>
<a href="https://en.wikipedia.org/wiki/PGP_word_list">https://en.wikipedia.org/wiki/PGP_word_list</a><br>
<a href="https://en.wikipedia.org/wiki/Public_key_fingerprint">https://en.wikipedia.org/wiki/Public_key_fingerprint</a><br>
<a href="https://wiki.gnome.org/Apps/Seahorse">https://wiki.gnome.org/Apps/Seahorse</a><br>
<a href="https://riseup.net/en/security/message-security/openpgp/best-practices">https://riseup.net/en/security/message-security/openpgp/best-practices</a></p>

