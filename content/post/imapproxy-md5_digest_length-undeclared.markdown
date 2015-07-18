---
wordpress_id: "97"
wordpress_url: "http://www.ericgar.com/2007/02/24/imapproxy-md5_digest_length-undeclared/"
title: "IMAPProxy: MD5_DIGEST_LENGTH undeclared"
date: "2007-02-24"
---
When I attempted to compile <a href="http://www.imapproxy.org/">IMAPProxy</a> on Linux 2.6.18 with GCC 4.1.1, I received this error:

<pre>
src/imapcommon.c: In function 'Get_Server_conn':
src/imapcommon.c:380: error: 'MD5_DIGEST_LENGTH' undeclared (first use in this function)
</pre>

The <a href="http://lists.andrew.cmu.edu/pipermail/imapproxy-info/2006-March/000535.html">post by Jakob Hirsch</a> on the <a href="https://lists.andrew.cmu.edu/mailman/listinfo/imapproxy-info">IMAPProxy mailing list</a> provided the answer:

> Problem is that md5.h is not longer included by evp.h (which is included
> by src/imapcommon.c), so MD5_DIGEST_LENGTH is not defined.
> `#include <openssl/md5.h>` in src/imapcommon.c fixed it for me.
