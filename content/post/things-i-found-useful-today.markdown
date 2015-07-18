---
wordpress_id: "232"
wordpress_url: "http://www.ericgar.com/?p=232"
title: "Things I found useful today"
date: "2008-06-24"
---
Really, I found most of these useful last week. But then I got my GPT and my partition table hopelessly out of sync and had to spend my weekend rebuilding my laptop, instead of getting my DB2 administration foo on. I've revisited almost all of them and found some new ones.

<ul>	<li><a href="https://wiki.ubuntu.com/MacBookPro">The authoritative reference on Ubuntu on a MacBook Pro</a> is just a life-saver. My ubuntu experience would be so much more hell without its existence.</li>
	<li><a href="http://shreevatsa.wordpress.com/2007/07/31/using-gmail-with-mutt-the-minimal-way/"> Using Mutt with Gmail (the minimal way)</a> gives a copy-and-pasteable few lines on how to use Mutt with Gmail over IMAP. Easy to remember, but easier to look up and forget. Most of the other references are with fetchmail. The only thing I would add is to include other IMAP folders as mailboxes: 

<code>mailboxes ="INBOX"
mailboxes ="[Gmail]/Sent Mail"
mailboxes ="[Gmail]/Drafts"
mailboxes ="[Gmail]/Spam"
mailboxes ="[Gmail]/All Mail"
</code>
</li>
	<li>The official <a href="http://www.mutt.org/doc/PGP-Notes.txt">Mutt and GPG Notes</a>.</li>
	<li><a href="https://www.myopenid.com/">myopenid</a> gives free domain-customized <a href="http://en.wikipedia.org/wiki/OpenID">openids</a>. My openid is now http://eric.ericgar.com. Watch out.</li>
	<li><a href="http://www.arsgeek.com/?p=3344">This post on CPU scaling</a> reminded me how to root-enable some CPU tools to allow me to easily scale my processor speeds</li>
	<li>The <a href="http://lifehacker.com/">Lifehacker</a> post detailing <a href="http://lifehacker.com/396741/functional-firefox-user-styles">how to minimize Firefox 3's toolbar</a>. Do it. You'll feel less constrained. It's a new world.</li>
	<li><a href="http://www.ibm.com/developerworks/db2/library/techarticle/0202kline/0202kline.html">Recovering from a failed LOAD operation in DB2 LUA</a>. Enough said. </li>
	<li><a href="http://www.ibm.com/developerworks/linux/library/l-cfs/?ca=dgr-lnxw04CFC4Linux">An overview of the Completely Fair Scheduler in Linux</a> and why I probably won't use it ever.</li>
	<li>Doug, who pointed out there exists an <a href="http://dougsland.livejournal.com/57614.html">Ubuntu package for mutt with some really useful patches</a> that vanilla mutt doesn't roll with. It saved me probably a good 45 minutes of patching mutt myself out of the source.</li>

</ul>

Remember kids: 64-bit web browsing in Linux currently sucks. Avoid it. Since I doubt you have over <a href="http://www.spack.org/wiki/LinuxRamLimits">4GB of RAM</a>: Do the right thing. Go 32-bit.

Update: Here's a screenshot of my ff menubar:

<a href='/uploads/2008/06/ffmenubar.png'><img src="/uploads/2008/06/ffmenubar.png" alt="" title="ffmenubar" class="alignnone size-medium wp-image-233" /></a>

