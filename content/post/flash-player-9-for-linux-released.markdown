---
wordpress_id: "36"
wordpress_url: "http://www.ericgar.com/2006/10/18/flash-player-9-for-linux-released/"
title: "Flash Player 9 for Linux Released"
date: "2006-10-18"
---
(I'd like to do my part for contributing to a PageRank:) I'm very pleased to see that the Flash team over at Adobe has released a very working <a href="http://labs.adobe.com/technologies/flashplayer9/">Flash Player 9 Beta for Linux</a>.

I've tested it on a few videos in YouTube and the results are obvious: Sound and video are now synchronized well. I may very well be using my Linux desktop a whole lot more now.

I give a big thanks to the Flash team, especially <a href="http://weblogs.macromedia.com/emmy/archives/2006/10/beta_refresh_on.cfm">Emmy Huang</a>, the team leader.

For reference, I had to place the <code>libflashplayer.so</code> file into <code>/usr/lib/nsbrowser/plugins</code> and delete <code>/usr/lib/nsbrowser/plugins/flashplayer.xpt</code>. It kind of seems odd to me that my installation of Firefox is using the <code>nsbrowser</code> directory rather than the <code>/usr/lib/mozilla-firefox/plugins</code> directory. Whatever floats its boat, I guess.

<span><strong>Update</strong>: I unmerged ebuild <code>netscape-flash</code> from Portage to get rid of <code>/usr/lib/nsbrowser</code>. Firefox still believed I had Flash Player 7 installed. Then, I enabled the configuration option <code>plugin.expose_full_path</code> in <code>about:config</code> and figured out that the plugin it found was in <code>/usr/lib/mozilla-firefox/plugins</code> and despite the fact that that file has a date of October 19, it wasn't Flash Player 9. Deleted the libflashplayer.so there, restarted, and works well.</span>
