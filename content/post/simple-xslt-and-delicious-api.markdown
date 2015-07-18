---
wordpress_id: "382"
wordpress_url: "http://ericgar.com/?p=382"
title: "Simple XSLT and Delicious API"
date: "2009-04-07"
---
I decided I wanted to take five minutes today to implement a simple idea I've had for a long time. When I get home, I typically eat dinner while reading news in my RSS reader and I've found less and less time to write my own posts. I simply wanted to be able to bookmark the things I found interesting and have them automatically appear as a weblog entry.

Part one was getting the data from <a href="http://delicious.com/minusnine">Delicious</a>, which is the obvious choice for bookmarking. They make it really easy to get all bookmarks made on the last day of activity with a given tag with their <em><a href="http://delicious.com/help/api#posts_get">posts/get</a></em> function. To get all bookmarks I made today with a tag of "viewed", I simply ran:

<code>$ curl 'https://minusnine:password@api.del.icio.us/v1/posts/get?dt=2009-04-07&amp;tag=viewed' > todays.xml</code>

which saves an XML form of the posts.

Transforming the XML into a list is a trivial task but I saw two approaches: <a href="http://www.gnu.org/software/sed/">sed</a> would do nicely for the crufty old sysadmin in me, but <a href="http://en.wikipedia.org/wiki/XSL_Transformations">XSLT</a> satisfies the young whippersnapper in me. 

The XSLT is [very simple and elegant][xml].

[xml]: /uploads/2009/04/xslt.xml

Combining these with <a href="http://xmlsoft.org/XSLT/xsltproc2.html">xsltproc</a>:

`$ xsltproc posts.xsl todays.xml`

yields a cut-and-pastable post. The next step is to pipeline these (also trivial) and upload it to my Wordpress installation.

