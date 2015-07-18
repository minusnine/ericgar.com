---
wordpress_id: "401"
wordpress_url: "http://ericgar.com/?p=401"
title: "ratpower.sh: a power menu for ratpoison"
date: "2009-04-19"
---
I've started to use ratpoison on my netbook. I've made a script that displays a menu of power-related actions and the current state of the battery. Here is a screenshot:

<blockquote>
[discharging 58%]
SLEEP
HIBERNATE
REBOOT
SHUTDOWN
LOCK
</blockquote>

The script uses <a href="http://packages.ubuntu.com/dapper/ratmenu">ratmenu</a> to actually display the menu and uses dbus to send signals.

<a href="/uploads/ratpower.sh">download ratpower.sh</a>
