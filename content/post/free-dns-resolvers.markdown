---
wordpress_id: "163"
wordpress_url: "http://www.ericgar.com/2007/10/17/free-dns-resolvers/"
title: "Free DNS Resolvers"
date: "2007-10-17"
---
I just recently spent a few hours diagnosing a very inconsistent network behavior when browsing between websites. I've known for a while that it's probably a DNS issue. After testing, I found that xo.net's (the bandwidth seller to my ISP) DNS servers are just intermittently dog slow in resolving individual queries, up to 10 seconds a piece.

I demoed using opendns.com's free resolvers, which indeed are fairly fast, but load my browser with advertising for every domain name typo. Not going to fly.

As a temporary fix instead, I started using local (and my former) universities' DNS servers, as they're really quick, reliable, close, and most importantly, wide open to the world. I try to be mindful of using other people's resources, though, prompting me to ask my coworkers, "Is it ethical to use someone else's DNS resolver?" 

Just today on the <a href="http://www.nylug.org/mlist/index.shtml">NYLUG mailing list</a> was a post about <a href="http://nylug.org/pipermail/nylug-talk/2007-October/035776.html">using freely available DNS servers</a> instead of crap ones. Indeed, a <a href="http://www.dnsserverlist.org/">list of open DNS servers</a> exists and even conveniently geolocates the closest DNS server. Problem is, the three it suggested for me are also really slow. Where the uni DNS servers are about 20ms away, all three random DNS servers bottomed out at 140ms. It appears that one of them, even, is some guy on an ADSL line.

Which just leads me to believe I should set up my own <a href="http://www.tinydns.org/">tinyDNS</a> server and achieve ultimate roundtrip times through tuning. Alas, my router does run <a href="http://www.dd-wrt.com">dd-wrt</a>, but does not have the necessary memory for yet another service.

Ah the problems a geek faces when he gets home from his day (and, for the last while, night) job. 
