---
wordpress_id: "31"
wordpress_url: "http://www.ericgar.com/2006/10/15/subversion-add-all-with-spaces-in-filenames/"
title: "Subversion: Add all, with spaces in filenames"
date: "2006-10-15"
---

I found a one-liner few days ago that exactly accomplished what
I needed to do: add a lot of files with spaces in each filename
that exist in various parts of a source tree to the subversion
respository. I reproduce the line here, to ensure use for posterity:

{{< highlight bash >}}
svn status | grep "^?" | sed -e 's/? *//' | sed -e 's/ /\ /g' | xargs svn add</code>
{{< /highlight>}}

This one-liner is shamelessly ripped from [Britt Selvitelle's blog][blog].

[blog]: http://lukewarmtapioca.com/articles/2005/06/01/adding-all-new-files-with-svn.html
