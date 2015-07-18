---
wordpress_id: "313"
wordpress_url: "http://www.ericgar.com/?p=313"
title: "Workaround: Blank Google Reader"
date: "2008-11-09"
---
There have been sporadic reports that <a href="http://groups.google.com/group/google-reader-troubleshoot/browse_thread/thread/bed8a9e18092e450?pli=1">Google Reader does not display content</a> sometimes. For me, this is chronic, but I haven't isolated the cause yet. There is a work around that I haven't seen discussed: the page displays properly with the sidebar expanded. When the page renders empty, there is not a lot of hint that it can be expanded.

Here is how Reader renders initially for me:

<a href="/uploads/2008/11/blank_reader.png"><img src="/uploads/2008/11/blank_reader.png" alt="" title="blank_reader" width="400" height="324" class="alignnone size-full wp-image-314" /></a>

Hovering over the left edge of the page reveals the sidebar expansion tab:

<a href="/uploads/2008/11/blank_reader_bar.png"><img src="/uploads/2008/11/blank_reader_bar.png" alt="" title="blank_reader_bar" width="500" height="405" class="alignnone size-full wp-image-315" /></a>

And expanding the sidebar immediately reveals the content of the page:

<a href="/uploads/2008/11/nonblank_reader.png"><img src="/uploads/2008/11/nonblank_reader.png" alt="" title="nonblank_reader" width="500" height="405" class="alignnone size-full wp-image-316" /></a>

I haven't dug into Firebug to see why this is the case yet, but it should be fixable with a Stylish or Greasemonkey script, I'd imagine. 

<span><strong>Update</strong>: Mihai from the Google Reader team provides this work around in the comments: "The 'blank screen' problem seems to be caused by a negative page zoom value (while the screen is blank, you can try going to to the View menu, and in the Zoom submenu choose Reset)." Thanks Mihai!</span>
