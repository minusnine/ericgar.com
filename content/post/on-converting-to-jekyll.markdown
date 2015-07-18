---
title: "On Converting to Jekyll"
date: "2010-09-28"
---

I [wrote] that I spent some time recently to convert my Wordpress-based
weblog to [Jekyll]. As partially mentioned in that post, Jekyll is a
"blog-aware" static content generator written in Ruby that uses a
lot of existing components, like [Maruku] for Markdown formatting and
[Liquid] for templating.

[wrote]: http://ericgar.com/2010/09/26/redesigned-ericgar-com/
[Jekyll]: http://github.com/mojombo/jekyll
[Maruku]: http://github.com/nex3/maruku
[Liquid]: http://www.liquidmarkup.org/
[Markdown]: http://daringfireball.net/projects/markdown/

The documentation for each of these components is at times a bit
sparse and strangely arranged. The following pages I just found but
I should have read them earlier:

* [Overview of Jekyll]
* [Bloging like a Hacker]
* [Jekyll]

[Overview of Jekyll]: http://www.oiledmachine.com/posts/2008/12/27/overview-of-jekyll--a-static-site-generator-written-in-ruby.html 
[Bloging like a Hacker]: http://tom.preston-werner.com/2008/11/17/blogging-like-a-hacker.html
[Jekyll]: http://henrik.nyh.se/2009/04/jekyll

Between those three, they do a good job of discussing what Jekyll is
and where it succeeds and fails. I'll just comment on my process.

Looking over and stealing the structure of [existing sites], including
[mine] is recommended. Of course, be sure to respect the original
author's copyright for content and code (mine is [Creative Commons]
and [BSD], respectively). I borrowed [Tate Johnson]'s, who recently
[redesigned his site][redesigned] to explicitly mimic Latex output
with CSS3 and HTML5.

[existing sites]: http://github.com/mojombo/jekyll/wiki/Sites
[mine]: http://github.com/minusnine/ericgar.com
[Creative Commons]: http://creativecommons.org/licenses/by-nc-sa/3.0/
[BSD]: http://www.opensource.org/licenses/bsd-license.php
[Tate Johnson]: http://tatey.com/
[redesigned]: http://tatey.com/2010/09/24/a-latex-default-formatting-inspired-design/

The conversion process was both easy and painful: it was easy to get
started and get all of the information in place, but was painful to
import all 220 existing posts. *Exporting* the posts were painless:
the provided [Wordpress migration tool] worked like a charm. I did
have to tweak it a bit to get topic name &amp;rarr; file name conversion
just right, and to pull down the few drafts of unfinished content in
the database.

[Wordpress migration tool]: http://github.com/mojombo/jekyll/wiki/Blog-Migrations

Converting the posts to valid Markdown pages would have been a
Herculean task, but I decided that I at least wanted to convert them
to a format that Maruku would parse and export to valid HTML. Maruku
is particular about where raw HTML can appear, namely approximately
in its own block. A lot of time was spent wrapping things in
&amp;lt;span&amp;gt;s to get Maruku to stop complaining. I wrote some shell
aliases to help me with the edit-debug-commit cycle. For a large
portion of the attention-required posts, I just converted to valid
Markdown.

Wordpress has a nice feature that I started using liberally: if a
media URL appears in the body of a post and is not contained in a tag,
it uses the [oEmbed] standard to convert the media to an HTML
representation. This meant that I had to revise all of my YouTube
and Flickr references to be valid HTML or Markdown. I was half way done
writing a Jekyll plugin that would support oEmbed, but then decided
this defeats the purpose of having a simple, static conversion process.

[oEmbed]: http://www.oembed.com/

The rest of the stuff was easy. I wrote an RSS 2.0 feed and heavily
modified someone's existing Atom one (both of which you should steal
from me). I used [Yahoo's YUI 2 Reset-Base-Font and Grids CSS][css]
as I hate dealing with browser-specific CSS differences. I used
[W3C's validator][validator] constantly. I wrote a [Perl-based unit test][unit]
to make it easy and integrated in [my Rakefile].

[css]: http://developer.yahoo.com/yui/base/
[validator]: http://validator.w3.org/unicorn/
[unit]: http://github.com/minusnine/ericgar.com/blob/master/_t/validate.t
[my Rakefile]: http://github.com/minusnine/ericgar.com/blob/master/Rakefile

I've so far decided against having comments, but I do have a trivial
[topic branch] that integrates [Disqus] comments. I'll probably merge
it some day soon.

[topic branch]: http://github.com/minusnine/ericgar.com/tree/topic%2Fcomments
[Disqus]: http://disqus.com/

My next rainy day task will be categorizing or tagging posts,
which Jekyll supports strangely, so a million different people have
reimplemented it in different ways.

This work was worth it to uncomplicate managing this simple
weblog. I can compose in Vim, write in a decent markup language,
have Makefile-like automation for testing and deployment, maintain
revision history and have sane branching with Git, and not have to
deal with security upgrades and unnecessary slowness.
