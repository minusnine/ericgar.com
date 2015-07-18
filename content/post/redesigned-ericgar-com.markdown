---
title: "Redesigned ericgar.com"
date: "2010-09-26"
---

I decided that managing Wordpress on my web host was not interesting
any longer. I wanted to host my simple blog in git, to compose using
`vim`, and to not have to manage a database or the insecurities of a
web application. Not to mention that the RSS feed was slightly broken
and the pages were not standards compliant.

Consequently, I redesigned this site using [Jekyll], a Ruby-based
"blog-aware" publishing tool that is the backend behind [GitHub]
pages. Instead of dynamically generating content, it transforms all
the content to static HTML files. It uses [Markdown] for its markup
language, so composing directly in cumbersome HTML continues to not
be necessary. I now have both a suite of unit tests that are run
prior to deployment, as well as a QA site to test changes before
going production with them.

[jekyll]: http://github.com/mojombo/jekyll 
[GitHub]: http://github.com/mojombo/jekyll
[Markdown]: http://daringfireball.net/projects/markdown/

The unfortunate side effect was that your feed reader probably went
a little crazy and marked-as-new a bunch of old posts. Apologies
for this: it wasn't worth keeping the same timestamps. Also, there
was a late-breaking bug in the RSS feed (but [the recommended Atom
feed][feed] was fine).

[feed]: http://ericgar.com/2008/04/05/rss-20-vs-rss-93-vs-atom-03/

The source for this blog is [available on GitHub][repo]. Let me know
if anything is broken.[^old] 

[^old]: Admittedly, I haven't taken the time to make
all posts prior to those on the front page XHTML Strict compliant as
it would take too much time.

[repo]: http://github.com/minusnine/ericgar.com
