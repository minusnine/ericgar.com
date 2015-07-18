---
wordpress_id: "69"
wordpress_url: "http://www.ericgar.com/2006/12/14/thecus-n2050-on-gentoo-linux/"
title: "Thecus N2050 on Gentoo Linux"
date: "2006-12-14"
---
I just received my two <a href="http://www.seagate.com/products/personal/barracuda.html">Seagate 7200.9 250GB drives</a> and have already installed them into my <a href="http://www.thecus.com/products_over.php?cid=1&amp;pid=3">Thecus N2050 RAID DAS device</a> (both <a href="http://www.newegg.com/Product/Product.asp?Item=N82E16822148111">hd</a> and <a href="http://www.newegg.com/Product/Product.asp?Item=N82E16822102002">das</a> on NewEgg). I was initially worried because the 7200.9 drives are not listed on the <a href="http://www.thecus.com/download/other/N2050HDDCompatibilityReport-951003.pdf">Thecus approved drive list</a>, despite the 7200.7 and 7200.10 devices being there. They work fine. The 7200.9 250GB drives were $10 cheaper than the 7200.10 drives. Since I don't really need performance but do want space, the older model was fine for me.

I'm fairly impressed with the setup. The box was easy to install: I slapped the two drives in, installed the included SATA card (Silicon Image, Inc. SiI 3512) in my Linux box, compiled modules for <code>siimage</code> and <code>sata_sil</code> (I'm not yet sure if I need the former; there is no documentation for installation on Linux).

My future plans are to purchase another N2050 in which I'll place the two 250G 7200.9 drives I already have in my Linux box (currently doing LVM to extend the partition space to 500GB, but no RAID). The only problem is that I've run out of PCI slots in my server, so I'd either have to use another server or buy a SATA PCI card with two eSATA connectors. Oh the problems of a geek.

I should also buy another 7200.9 250GB for the event that a drive fails. I'm not sure whether to buy this extra drive now as the 7200.10s are hitting the market or to wait a while. Or, I may not even do this and just rely on Seagate's ridiculous 5 year warranty. On second thought, I'll probably just do that.

<span><strong>Update 12/26/06</strong>: I've been using the <code>sata_sil</code> driver without any problems. Below are my (fairly decent) stats for the drive:</span>

<code>
/dev/sdc:
 Timing cached reads:   1728 MB in  2.00 seconds = 864.08 MB/sec
 Timing buffered disk reads:  156 MB in  3.02 seconds =  51.58 MB/sec
</code>
