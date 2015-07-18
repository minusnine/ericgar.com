---
wordpress_id: "92"
wordpress_url: "http://www.ericgar.com/2007/02/09/mythtv-problem/"
title: "MythTV problem"
date: "2007-02-09"
---
I just fixed a random problem with our MythTV box. The backend crashed for some reason, so we had to power cycle the computer. Upon starting mythbackend up, there was an error message about a failed schema update.  

The fix was to manually decrement the DBSchemaVer setting:

<pre>
mysql> update settings set data=data-1 where value='DBSchemaVer';
</pre>

And reload the backend.
