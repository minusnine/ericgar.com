---
wordpress_id: "98"
wordpress_url: "http://www.ericgar.com/2007/04/10/sendmail-smtp-auth-config-explanation/"
title: "Sendmail SMTP AUTH Config Explanation"
date: "2007-04-10"
---
Both for my own knowledge and that of the internet, I took the time to finely document my sendmail.mc configuration file for use with SMTP AUTH to an external SMTP server. In this configuration, my local users can send mail to the rest of the internet by authenticating with my ISP's SMTP server and send mail through it.

The following sites were extremely helpful when initially trying to configure this:
	<ul><li><a href="http://www.sendmail.org/~ca/email/auth.html">SMTP AUTH in sendmail 8.10-8.13</a></li>
<li><a href="http://www.madboa.com/geek/sendmail-auth/">Sendmail SMTP AUTH Quick Start</a></li>
<li><a href="http://www.joreybump.com/code/howto/smtpauth.html">Using SMTP AUTH and STARTTLS with sendmail</a></li>
<li><a href="http://www.linuxquestions.org/questions/showthread.php?t=224543">Sendmail SMTP AUTH Howto</a></li>
<li><a href="http://slacksite.com/other/auth.html">Setting up SMTP AUTH with sendmail and Cyrus-SASL</a></li></ul>

