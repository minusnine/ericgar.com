---
wordpress_id: "600"
wordpress_url: "http://ericgar.com/?p=600"
title: "XSLT for OPML to XHTML List"
date: "2009-12-28"
---
The following some XSLT sufficient to transform an OPML file into a list, ready for you to edit and post:

{{< highlight xml >}}
<?xml version="1.0" encoding="UTF-8"?>

<xsl:stylesheet version="1.0"
  xmlns:xhtml="http://www.w3.org/1999/xhtml"
  xmlns="http://www.w3.org/1999/xhtml"
  xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
  xmlns:xs="http://www.w3.org/2001/XMLSchema"
  exclude-result-prefixes="xhtml xsl xs">

    <xsl:template match="body">
        <ol><xsl:text>
</xsl:text>
            <xsl:for-each select="outline">
                <li> <a href="{@htmlUrl}" ><strong><xsl:value-of select="@text" /></strong></a> - your text </li><xsl:text>
</xsl:text>
            </xsl:for-each>
        </ol>
    </xsl:template>
</xsl:stylesheet>
{{< /highlight >}}

It isn't 100% complete, but will get you a list of the form:

{{< highlight html >}}
<ol>
    <li><a href="URL"><strong>Title</strong></a> - your text</li>
</ol>
{{< /highlight >}}

In a sane interface to an operating system, you can run the following to produce transformed output, given the XSL above and an OPML file:

<pre>$ xsltproc extract.xsl google-reader-subscriptions.xml</pre>
