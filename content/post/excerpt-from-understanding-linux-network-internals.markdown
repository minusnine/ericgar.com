---
wordpress_id: "640"
wordpress_url: "http://ericgar.com/?p=640"
title: "Excerpt from \"Understanding Linux Network Internals\""
date: "2009-12-30"
---
I'm currently reading <a href="http://www.oreillynet.com/pub/au/2398">Christian Benvenuti</a>'s excellently written <em><a href="http://oreilly.com/catalog/9780596002558/">Understanding Linux Network Internals</a></em> from O'Reilly which is helping to shore up my knowledge about how the networking stack is implemented in Linux. It's a fantastic read so far, on course to match <em><a href="http://www.amazon.com/Linux-Kernel-Development-Robert-Love/dp/0672325128">Linux Kernel Development</a></em> by <a href="http://rlove.org/">Robert Love</a>, one of my all-time favorite books.

The following paragraph from Benvenuti's book really made me step back and take a look at the bigger picture:

<blockquote>
A device driver can also disable the egress queue before a transmission (to prevent the kernel from generating another transmission request on the device), and re-enable it only if there is enough free memory on the NIC; if not, the device asks for an interrupt that allows it to resume transmission at a later time. Here is an example of this logic, taken from the <em>el3_start_xmit</em> routine, which the <em>drivers/net/3c509.c</em> driver installs as its <em>hard_start_xmit</em> function in its <em>net_device</em> structure...
</blockquote>

That passage is just brilliant. <strong>Count the occurrences of jargon there!</strong>

I guess i should read more <a href="http://arxiv.org/">arXiv</a> papers to better understand how much larger the world actually is.
