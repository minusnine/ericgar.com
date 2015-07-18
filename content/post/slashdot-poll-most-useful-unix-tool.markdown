---
wordpress_id: "516"
wordpress_url: "http://ericgar.com/?p=516"
title: "Slashdot poll: Most Useful UNIX Tool"
date: "2009-09-19"
---
In response to the current <a href="http://slashdot.org">/.</a> poll, <a href="http://slashdot.org/pollBooth.pl?qid=1855&amp;aid=-1">Most Useful UNIX Tool</a>: I have the following from my ~/.zhistory at work:

<blockquote><pre>
$ for cmd in sed grep cat find telnet init exit ; do echo -n $cmd= ; grep -c "$cmd " ~/.zhistory; done
sed=85
grep=875
cat=762
find=126
telnet=15
init=32
exit=33
</pre></blockquote>

Which is interesting to me. I can explain some of the counts:

* I run `cat $file | (grep|sed|awk) | ...` too often out of a good/bad habit.
* `exit` is small because I've bound `x` to exit (which itself is a bad habit; I should have just started using `Control-D`.
* `ssh` > `telnet`
* I don't have root, so `init` isn't really used, and my regex isn't careful enough to eliminate `disk_init`.
* `sed` is mega-useful, but is usually the second command in a pipeline. I want to start using `perl -pe` more.
