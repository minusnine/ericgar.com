---
wordpress_id: "203"
wordpress_url: "http://www.ericgar.com/2008/01/29/bzipd-tar-file-returns-error/"
title: "bzip'd tar file returns error"
date: "2008-01-29"
---

right. so I got this today when trying to untar all of my academic
work from an archive:

{{< highlight bash >}}
ericgar@babbage extusb$ tar -xjf columbia-2007-10-31.tar.bzip2
You can use the `bzip2recover' program to attempt to recover
data from undamaged sections of corrupted files.
tar: Child returned status 2
tar: Error exit delayed from previous errors
ericgar@babbage extusb$ bzip2recover columbia-2007-10-31.tar.bzip2 
bzip2recover 1.0.4: extracts blocks from damaged .bz2 files.
bzip2recover: searching for block boundaries ...
bzip2recover: I/O error reading `columbia-2007-10-31.tar.bzip2', possible reason follows.
bzip2recover: Input/output error
bzip2recover: warning: output file(s) may be incomplete.
{{< /highlight >}}

I was thinking, "Well, data loss sucks." 

But it turns out the underlying filesystem was mounted read-write
on a read-only mount point. D'oh. I feel like tar and bzip2recover
could have told me that off the bat.
