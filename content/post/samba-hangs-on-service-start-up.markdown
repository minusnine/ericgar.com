---
wordpress_id: "21"
wordpress_url: "http://www.ericgar.com/blog/2006/10/14/samba-hangs-on-service-start-up/"
title: "Samba hangs on service start up"
date: "2006-10-14"
---
I just encountered a problem where Samba failed to start and just hung before daemonizing.

<img id="image18" alt="Samba fails to start." src="http://www.ericgar.com/uploads/2006/10/picture-3.png" />
The last few messages left in /var/log/samba/log.smbd referred to trying to find CUPS: "cups server left to default localhost" and "Unable to connect to CUPS server localhost - Connection timed out". I don't use CUPS since I don't have a printer and the service was not active.

<img id="image19" alt="Samba error log" src="http://www.ericgar.com/uploads/2006/10/picture-4.png" />
Why would it hang on connecting to a service that doesn't exist? For kicks, I started up CUPS without any defined printers to see what would happen to SMB. No luck, smbd still hung at exactly the same place.

Then I realized that I'm a big dummy. Network interface lo was not loaded. I'm not entirely sure how it was removed from my startup configuration. After doing an 'ifconfig lo up', I started samba:

<img id="image20" alt="Samba starts." src="http://www.ericgar.com/uploads/2006/10/picture-5.png" />

Moral of the story: services need lo to use other services on the same box.
