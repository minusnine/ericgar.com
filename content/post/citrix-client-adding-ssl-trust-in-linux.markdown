---
wordpress_id: "716"
wordpress_url: "http://ericgar.com/?p=716"
title: "Citrix Client: Adding SSL trust in Linux"
date: "2010-06-19"
---

The company I work for somehow decided to get their SSL certificate
for their employee portal from a CA not trusted by one of the main
applications on the portal, the Citrix Client. (There may be an
addition of "Xen" in the actual product name, who knows).

I received the error:

> "You have not chosen to trust "UTN-USERFirst-Hardware", the issuer
> of the server's security certificate (SSL error 61).

The answer is fairly clear: we need to add this particular SSL cert,
or the CA's root cert into our trust anchor. To do this, we could
save this cert and put it in the right place, or we can trust that
OpenSSL has the appropriate root certs.

I chose the latter. To do this, we simply move the certificates
trusted by the Citrix installation and symlink to the OpenSSL-based
certs. Not rocket science.

<pre>
$ sudo mv /usr/lib/ICAClient/keystore/cacerts /usr/lib/ICAClient/keystore/cacerts.old
$ sudo ln -sf /usr/lib/ssl/certs /usr/lib/ICAClient/keystore/cacerts
</pre>

Voila. Retrying to connect to the Citrix (?:Xen)?(?:App)? Server works.
