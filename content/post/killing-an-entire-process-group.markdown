---
wordpress_id: "777"
wordpress_url: "http://ericgar.com/?p=777"
title: "Killing an entire process group"
date: "2010-08-19"
---
Today, one of my colleagues inadvertently set off a linear fork bomb of ksh processes on a critical infrastructure machine. Linux held up well to this, since basically these processes were just entries in the scheduling table and very little else. The system was pretty responsive, albeit with a load average hovering between 800 and 3000, and we ran out of process identifiers fairly quickly.

Clearly, all of these processes would have the same process group, so we just needed to kill the process group as a whole. My other colleague and I came up with a good, stupid simple solution that neither of us had immediately thought of. As it turns out, Perl's kill() system call interface has an interesting caveat that allows you to kill processes in exactly this way: when the signal is specified as a negative number, the function treats the other parameters as process groups to kill.

Actually, it turns out a lot of systems and shells allow specifying the process group ID as negative numbers to kill command/built-in where the pids would ordinarily go to specify the same action. It's too bad my colleague was using an old version of ksh.
