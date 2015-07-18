---
wordpress_id: "304"
wordpress_url: "http://www.ericgar.com/?p=304"
title: "New Computer!!!!111"
date: "2008-11-09"
---
This week, I purchased a new computer:

<blockquote>
Shuttle SG33
Intel Core 2 Quad Q6600 2.4GHz
2 x 1GB DDR2 DIMMs
Cheapo ATI Radeon X1650 ( 2 x DVI-D)
2 x Acer x243w 24" LCDs (yes, that's 2!)
</blockquote>

(I already had cheap hard disks, <a href="http://www.logitech.com/index.cfm/mice_pointers/trackballs/devices/166&amp;cl=us,en">a gorgeous mouse</a>, and <a href="http://matias.ca/tactilepro/">a nice keyboard</a>; though, I envision I'll be bringing <a href="http://www.kinesis-ergo.com/advantage.htm">the love of my life</a> home from work more often).

I purchased all of the parts from <a href="http://newegg.com">newegg</a>, as they were offering some good specials for pretty much all the components, including free shipping for the monitors, which arrived the next day.

Assembling and getting the machine running was mostly without issue. The only complaint I have about the hardware is that the graphics card fan is extremely loud, something I will have to remedy in the coming weeks. Since I'm not a hardcore gamer anyway, I tried just unplugging the fan, but the card temporarily died a few minutes later. It <em>does</em> run extremely hot. That's what you get when you skimp on an important component and don't consult <a href="http://www.silentpcreview.com/">slientpcreview.com</a> beforehand.

I've found <a href="http://unetbootin.sourceforge.net/">UNetbootin</a> to be an invaluable tool despite the horrible name. I ran out of CD-Rs a long time ago and don't care to buy any more. UNetbootin has saved a considerable amount of time by transferring an iso+bootloader onto a hard drive. Highly recommended.

Getting X to deal with spanning across two monitors was a bit of a pain: the ATI Catalyst Control Center is broken and segfaults before the configuration is written out to disk but after the settings are applied to the driver. It took some poking around to realize that after setting the monitors up properly in Catalyst, then running the vanilla screen resolution manager, then blindly clicking "Apply", the settings become static.

I also had an issue where the max_cpu_freq for all four cores was static at 1.6 GHz and Intel Speed Step (EIST) did not work. This was resolved by upgrading the bios to the latest version, which is annoying when not running Windows and not owning a floppy drive. I got around this by <a href="http://www.aselabs.com/articles.php?id=243">following these instructions on creating a bootable USB disk with FreeDOS</a>. Admittedly, that is kind of a roundabout way of installing DOS to a disk (requiring installing a DOS emulator), but it works, where the other methods were far more complicated and had mixed results.

Now that I have a sweet rig and some more time (the latter due to life events), I hope to hammer out code more often. I have an idea to write a version of <a href="http://rstatd.sourceforge.net/">rup</a> that has the same features as <a href="http://fping.sourceforge.net/">fping</a>, or at least times out in a reasonable period.

And, now that Ubuntu Ibex is out, I need to try getting my PS3 running Linux again. Sadly,<a href="https://bugs.launchpad.net/ubuntu-ps3-port/+bug/289982"> it still doesn't support WPA at the moment</a>, meaning I need to <a href="http://www.monoprice.com/products/product.asp?c_id=102&amp;cp_id=10208&amp;cs_id=1020801&amp;p_id=2158&amp;seq=1&amp;format=2">get</a>/fetch from home a really long ethernet cable. 

<span><strong>Update</strong>: Despite this new rig, <a href="http://www.nycresistor.com/2008/11/09/because-newer-isnt-always-better/">I'm still kind of jealous of NYCR</a>.</span>
