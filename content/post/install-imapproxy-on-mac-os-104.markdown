---
wordpress_id: "63"
wordpress_url: "http://www.ericgar.com/2006/12/13/install-imapproxy-on-mac-os-104/"
title: "Install imapproxy on Mac OS 10.4"
date: "2006-12-13"
---
<ol>
<li><strong>Optional prerequisite, proctools</strong>: If you're going to use the init script to start imapproxy, you need to install <a href="http://proctools.sourceforge.net/">proctools</a> to let the script kill the process on shutdown. Installation is a little different than normal, but is explained nicely in the README.</li>
<li><strong>Do the normal thing</strong>
<ol>
<li><code>wget  up-imapproxy-1.2.5rc2.tar.gz</code></li>
<li><code>tar -xzf  up-imapproxy-1.2.5rc2.tar.gz</code></li>
<li><code>cd  up-imapproxy-1.2.5rc2</code></li>
<li><code>./configure</code></li>
<li><code>make</code></li>
</ol>
</li>
<li><strong>Install executable files</strong>: `make install` is broken as the user and group is hardcoded as <code>bin</code>. You can either edit the makefile or copy the necessary files manually, as below:
	<ol><li><code>sudo cp bin/in.imapproxyd /usr/local/sbin/</code></li>
		<li><code>sudo cp bin/pimpstat /usr/local/sbin/</code></li>
	</ol></li>
<li><strong>Install configuration files</strong>: <code>sudo make install-conf</code></li>
<li><strong>Install init files</strong>: <code>sudo mkdir /etc/init.d/ &amp;amp;&amp;amp; sudo mkdir /etc/rc.2 &amp;amp;&amp;amp; sudo make install-init</code></li>
<li><strong>Install Certification Authority certificates</strong>: The included documentation is unclear what to do if you intend to use imapproxy to connect to an IMAP server with TLS/SSL. First, imapproxy cannot connect to IMAP servers using port 993. Instead, imapproxy will only connect to servers using regular IMAP and invoking TLS though the <code>starttls</code> IMAP command (RFC 2595). imapproxy does not currently support a client connecting to it through TLS; it is intended to be run on the same host as the client. Second, you must install the appropriate Certification Authority (CA) certificates that will validate the TLS certificate offered by the server. I grappled with this for a while, even using a working imapproxy installation as a reference to no avail. In the end, I packaged all of the root CA certs from OS X's keychain into a .pem (which I've <a href="http://www.ericgar.com/uploads/files/cacerts.pem">made available here</a>). It is not necessary to create your own CA or certs as implied in the <code>README.ssl</code> included with the source.</li>
<li><strong>Edit the configuration file</strong>: <code>sudo nano imapproxy.conf</code> . I changed the <code>server_hostname</code> string to the IMAP server I want to connect to, the <code>listen_port</code> to allow other IMAP server connection attempts to succeed, the <code>listen_address</code> to <code>127.0.0.1</code> to only allow clients on the localhost to connect, the <code>force_tls</code> to <code>yes</code> to force TLS (although the IMAP server I'm connecting to uses LOGINDISABLED). Finally, I changed <code>tls_ca_file</code> to the file path of the CA certs. My working imapproxy configuration file is available <a href="http://www.ericgar.com/uploads/files/imapproxy.conf">here</a>.</li>
<li><strong>Start the proxy</strong>: <code>/etc/init.d/imapproxy start</code></li>
</ol>

That should be it for getting it to work. You should configure your IMAP client to connect to 127.0.0.1 on whatever port you specified above.

That's not the whole story, however. Mac OS X does not use init scripts but uses a much cooler launching mechanism called <code>launchd</code>. Launchd takes care of launching, maintaining, and closing daemons, meaning it eliminates a lot of hassle for the daemon programmer. 

I've created a <a href="http://www.ericgar.com/uploads/files/imapproxy.plist">launchd plist</a> that can be used with launchd to start imapproxy the right way in Mac OS X. If you use this, you do not have to use #5 or #8 above. Also, you must set <code>foreground_mode</code> to <code>yes</code>. Copy the plist to <code>/Library/LaunchDaemons</code> and use <code>sudo launchctl [load|unload] /Library/LaunchDaemons/imapproxy.plist</code> to start and stop the proxy. The proxy will also start whenever the machine is started (to change this, move the plist out of that directory).

There is more Apple information on <a href="http://developer.apple.com/documentation/Darwin/Reference/ManPages/man5/launchd.plist.5.html">launchd plists</a>, <a href="http://developer.apple.com/documentation/Darwin/Reference/ManPages/man1/launchctl.1.html#//apple_ref/doc/man/1/launchctl">launchctl</a>, and <a href="http://developer.apple.com/technotes/tn2005/tn2083.html">writing</a> <a href="http://developer.apple.com/macosx/launchd.html">daemons</a> that use launchd. 
