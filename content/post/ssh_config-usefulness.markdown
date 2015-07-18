---
wordpress_id: "95"
wordpress_url: "http://www.ericgar.com/2007/02/22/ssh_config-usefulness/"
title: "ssh_config usefulness"
date: "2007-02-22"
---
I'm sad to say it, but I didn't know about ssh_config until last week. Sometimes you think you're not that bad at what you do, but then there's just something that makes you realize you don't know as much as you think.

A long time ago, I obtained a copy of <a href="http://www.oreilly.com/catalog/wirelesshks/">wireless hacks</a> from a talk that was given on campus. Most of it were things I'd already thought of, the rest uninteresting. A good book for anyone who doesn't know anything about using wireless networking to their advantage. There was mention of using bash shortcuts to ssh into remote machines easily. 

So, instead of typing <code>ssh -X -Y -A ekg2002@canberra.cc.columbia.edu</code> each time I wanted to remotely log into a Columbia machine, I put that exact line into a executable file somewhere in my path. That way, I could just type <code>$ canberra</code> at my prompt and be logged in.

The problem with this scheme is that I can't easily, say, scp to the same machine. I always found this to be annoying.

Then I found ssh_config. It allows you to define host shortcuts in a single file including any ssh options. <code>man ssh_config</code> gives lots of information on it. So, to replicate the above ssh command, I add this to ~/.ssh/config:
<code>
Host canberra
HostName ekg2002@canberra.cc.columbia.edu
User ekg2002
ForwardAgent yes
ForwardX11 yes
ForwardX11Trusted yes
</code>
 
I can do this for all machines to which I ssh regularly. This is beneficial for two reasons: I can ssh and scp equally easily and I can manage this file centrally (and install it on all of my hosts). The downside is that I have four more characters to type to connect to a given host (which can be reduced to two through aliasing).
