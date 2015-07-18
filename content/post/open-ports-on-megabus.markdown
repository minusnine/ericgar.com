---
title: "Open ports on megabus to Boston"
date: "2010-11-02"
---

The following is nmap output over Megabus-provided wifi (en route
to Boston) showing the proxied ports.

<pre>
Starting Nmap 5.00 ( http://nmap.org ) at 2010-11-02 19:39 EDT
Interesting ports on dagny.detechnis.com (173.255.227.97):
Not shown: 987 filtered ports
PORT     STATE  SERVICE
53/tcp   open   domain
80/tcp   closed http
110/tcp  closed pop3
143/tcp  closed imap
443/tcp  closed https
465/tcp  closed smtps
587/tcp  closed submission
993/tcp  closed imaps
995/tcp  closed pop3s
1863/tcp closed msnp
5190/tcp closed aol
5222/tcp closed unknown
8080/tcp closed http-proxy

Nmap done: 1 IP address (1 host up) scanned in 30.28 seconds
</pre>

What's sort of annoying is that:

1. Port 22 (ssh) is blocked.
1. ICMP Time Exceeded messages are blocked (not shown), making `tracepath` unusable.

I will be attending the [Linux Plumbers Conference][lpc] and the [Beer Advocate
Belgian BeerFest][beer] in the coming days. More on those as they happen.

[lpc]: http://www.linuxplumbersconf.org/2010/
[beer]: http://beeradvocate.com/bbf/
