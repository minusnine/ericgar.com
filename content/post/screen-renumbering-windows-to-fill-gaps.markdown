---
wordpress_id: "206"
wordpress_url: "http://www.ericgar.com/2008/02/03/screen-renumbering-windows-to-fill-gaps/"
title: "Screen: Renumbering Windows to Fill Gaps"
date: "2008-02-03"
---
After running a single session of <a href="http://www.gnu.org/software/screen/">screen</a> for a long time, I often find that I have several gaps in the numerical ordering of windows. Using <a href="http://www.hmug.org/man/1/screen.php">:number</a> is definitely feasible, but it takes a bit more effort than I'd care to contribute every time I want to make my windows contiguously numbered.

I've created <a href="http://ericgar.com/uploads/screen-renumber.patch">a patch</a> against <a href="http://savannah.gnu.org/cvs/?group=screen">CVS HEAD</a> to fill in the holes of the window numbering. It simply moves windows to lower positions until there are no holes left. Any (constructive) comments are welcome.

<span>The patch can be found <a href="http://ericgar.com/uploads/screen-renumber.patch">here</a>. It was also sent to the <a href="http://lists.gnu.org/archive/html/screen-devel/2008-02/msg00000.html">screen-devel mailing list</a>.</span>
