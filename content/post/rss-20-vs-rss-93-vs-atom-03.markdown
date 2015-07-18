---
wordpress_id: "202"
wordpress_url: "http://www.ericgar.com/?p=202"
title: "RSS 2.0 vs. RSS .93 vs. Atom 0.3 ..."
date: "2008-04-05"
---
The other day, I was visiting a weblog that I wanted to include in my RSS aggregator. I clicked on the icon my web browser that indicates that the site provides such a feed and was presented with this[^*]:

<a href='http://www.ericgar.com/uploads/2008/01/chooserss.png' title=''><img class="bo" src='http://www.ericgar.com/uploads/2008/01/chooserss.png' alt='' /></a>

Great. Which one do I choose?  I guess it's clear: 2.0 is an order of magnitude better than .93, which it self must be three times better than 0.3. Right? 

Uhh...No.

Okay. I've been a developer for a while and I've even developed RSS-related stuff. If <em>I</em> don't know what the real differences are and how it affects my choice and subsequent enjoyment of the content, then I feel like most people wouldn't either.

As <a href="http://en.wikipedia.org/wiki/Syndication_format_family_tree">syndication format family tree</a> clearly shows, RSS .93 is the wicked step-child of earlier RSS 0.9x versions and the extinct scriptingNews formats. Basically, Dan Libby at Netscape borrowed (ahem.. <em>stole</em>, really, but for the better) in an effort toward standards. Then, RSS 2.0 is the inbred child of all RSS 0.9x versions, and, strangely, RSS 1.0. Then, the <a href="http://en.wikipedia.org/wiki/Atom_(standard)">Atom</a> format was created to make a fresh technology and leave all of the accumulated crud that an old protocol takes with it. 

What does that all mean? <em>Nothing</em>. Not when the end user doesn't care, just randomly picks one from the list, and hopes his or her client works well with it. Even after reading <a href="http://en.wikipedia.org/wiki/RSS">a more detailed account of RSS lineage</a>, do you care which version of RSS you use? [^**]

Any software developer will tell you that they've had the urge to throw out a piece of legacy code and start all over from scratch, applying best practices and lessons learned. That's what Atom is supposed to be. It's raison d'etre is to be the child who observes what his parents don't like about themselves and improves upon those aspects. 

On the Atom wikipedia page, these two points are listed among others under "Barriers to adoption":

<blockquote>Many sites choose to publish their feeds in only a single format. For example CNN, the New York Times, and the BBC offer their web feeds only in RSS 2.0 format.</blockquote>

...which is actually doing a service to the user. This shouldn't be criticized as a "barrier to adoption", but a embrace of usability.

<blockquote>Sites that publish Atom will often publish RSS as well.</blockquote>

But why? If backward compatibility is a concern, then continue <em>publishing</em> content in all formats that you've given to users in the past, but <em>advertise</em> only the current best format. 
I understand that backward support is good, so that people who subscribed to the RSS 0.93 feed don't get burned when support for RSS 2.0 comes along. I also understand that too much meaningless choice for an unknowing consumer is just that: meaningless. And, if we're supposed to be using standard technology, why are there three competing standards with no winner in sight? [^***]

Incompatibilities may exist with software being able to read the formats. [^****] Here is an informal survey of some popular feed readers on the formats discussed here:

* Google Reader: <em>RSS 0.92, RSS 1.0, RSS 2.0, Atom 0.3, Atom 1.0</em>
* NewsFire: <em>RSS 0.92, RSS 1.0, RSS 2.0, Atom 0.3, Atom 1.0</em>
* RSSOwl: <em>RSS 0.92,  RSS 1.0, RSS 2.0, Atom 0.3, Atom 1.0</em>
* Bloglines: <em>RSS 0.92,  RSS 1.0, RSS 2.0, Atom 0.3, Atom 1.0</em>
* NetNewsWire: <em>RSS 0.92,  RSS 1.0, RSS 2.0, Atom 0.3, Atom 1.0</em>
* FeedDaemon: <em>RSS 0.92,  RSS 1.0, RSS 2.0, Atom 0.3, Atom 1.0</em>

Do you get the point?

<span><a href="http://www.feedburner.com/">Feedburner</a>, a recent acquisition by Google, at least is <a href="http://feeds.feedburner.com/BurnThisRSS2">tending toward the user</a>:</span>

<a href='/uploads/2008/04/feedburner.png'><img src="/uploads/2008/04/feedburner-300x257.png" alt="" title="feedburner" width="300" height="257" class="bo alignnone size-medium wp-image-221" /></a>

Here we see a application-centric model of how to advertise syndication formats. Feedburner presents icons denoting popular applications that the user might use. If I'm a user of Pageflakes, I may not know anything about RSD 3.2 vs Atoms 1.4, but I do know that I go to www.pageflakes.com to see this week's Dilbert cartoon on my homepage.

<span><strong>Here's the bottom line</strong>: Stop advertising the older formats. It's fine to continue to serve up the others, just don't actively advertise it. No one cares what formats you advertise, or the format they click on, as long as they get the content they want.</span>

I've chosen Atom. I think it's in the winner in its modularity, feature-set, and future growth. I could go on about why I think it's the right choice for this application, but here's the point: <strong>no one cares</strong>.



[^*]: Admittedly, www.ericgar.com suffered from this affliction, which
are the default options for Wordpress. This has been locally remedied.

[^**]: Ironically, that wikipedia article has an "Incompatibilities"
section, with no "Features" section or similar. What is the (probably
unintended) implication of that?

[^***]: The Blu-ray vs. HD-DVD competition at least was better in this
regard: there was a financial motive that would produce a winner. This
is not so in RSS .93 vs 1.0 vs 2.0 vs Atom .93, Atom 1.0

[^****]: Strangely, this isn't listed under as a "barrier to adoption"
on the Atom wikipedia page. I wonder why?
