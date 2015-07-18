---
wordpress_id: "340"
wordpress_url: "http://ericgar.com/?p=340"
title: "Tired. Exhausted. Smelly"
date: "2008-12-07"
---
Ended my last primary on-call shift of the year just now, and not a minute too soon. I like to say that my job is at the intersection of system administration and adrenaline because for 37 hours a weekend we are constantly solving issues where data loss is not an option and speed is everything.

Anyway, two choice quotes from my weekend, emphasis mine:

From a Sybase ASE error log:
<blockquote>
2008/12/07 02:34:34.40 server  <strong>Error: 9970, Severity: 20, State: 1</strong>
2008/12/07 02:34:34.40 server  DBCC cannot update the finish time in dbcc_operation_log table for this operation(opid = '182') of database 'sybsystemprocs'. <strong>This can be patched by executing sp_dbcc_patch_finishtime</strong>.
</blockquote>

I love software that has been patched to tell you how to actually fix the issue but won't actually fix it for you.

Next up, an email from a first-level system administrator:

<blockquote>
We see that the host is now responding.
But <strong>host is on high load.</strong>
Please check your applications.

piias1332 /ms/user/v/venusrin 63$ rup pisdb38
pisdb38                  up  29 days,  2:42,    <strong>load average: 1.77 1.73  1.62</strong>
</blockquote>

Um. No. Just no.
