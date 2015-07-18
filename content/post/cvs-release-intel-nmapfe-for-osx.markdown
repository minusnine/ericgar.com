---
wordpress_id: "99"
wordpress_url: "http://www.ericgar.com/2007/02/27/cvs-release-intel-nmapfe-for-osx/"
title: "CVS Release: Intel NmapFE for OSX"
date: "2007-02-27"
---
The developer of <a href="http://faktory.org/m/software/nmap/">NmapFE for OS X</a> hasn't made a release since late 2004 and consequently is using an outdated version of Nmap. Outdated to the point it doesn't work:

sendto in send_ip_packet: sendto(5, packet, 28, 0, 127.0.0.1, 16) => Invalid argument
Sleeping 60 seconds then retrying

I've created an Intel build from CVS using nmap 4.20, which seems to
work for me. It is available for download here. (link removed)

<span><strong>Update:</strong> Actually, this doesn't work very well. I'll compile something soon that does.</span>
