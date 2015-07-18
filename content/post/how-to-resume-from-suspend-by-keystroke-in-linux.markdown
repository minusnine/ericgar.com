---
wordpress_id: "449"
wordpress_url: "http://ericgar.com/?p=449"
title: "how to: resume from suspend by keystroke in Linux"
date: "2009-07-08"
---
I've always wondered why Macs could wakeup from sleep by a mere keystroke, but my Linux boxes required me to press the power button. It turns out you can enable wake from suspend/hibernate in Linux by adding the following to /etc/rc.local, which is run at startup:

{{< highlight bash >}}
for i in `/bin/grep USB /proc/acpi/wakeup | /usr/bin/awk '{print $1}'`; 
do 
    echo $i > /proc/acpi/wakeup; 
done
{{< /highlight >}}

/proc/acpi/wakeup will then look something like:

<pre>
$ cat /proc/acpi/wakeup
Device  S-state   Status   Sysfs node
PCI0      S5     disabled  no-bus:pci0000:00
PEX0      S5     disabled  pci:0000:00:1c.0
PEX1      S5     disabled  pci:0000:00:1c.1
PEX2      S5     disabled  
PEX3      S5     disabled  
PEX4      S5     disabled  
PEX5      S5     disabled  
HUB0      S5     disabled  pci:0000:00:1e.0
IGBE      S5     disabled  
USB0      S3     enabled   pci:0000:00:1d.0
USB1      S3     enabled   pci:0000:00:1d.1
USB2      S3     enabled   pci:0000:00:1d.2
USB3      S3     enabled   pci:0000:00:1a.0
USB4      S3     enabled   pci:0000:00:1a.1
USB5      S3     enabled   pci:0000:00:1a.2
EHC1      S3     disabled  pci:0000:00:1d.7
EHC2      S3     disabled  pci:0000:00:1a.7
AZAL      S5     disabled  pci:0000:00:1b.0
</pre>

and voila: when your Linux box suspends, you can wake it up by pressing any key on your USB keyboard.
