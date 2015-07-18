---
wordpress_id: "208"
wordpress_url: "http://www.ericgar.com/2008/02/04/some-feedback-from-my-screen-patch/"
title: "Some feedback from my screen patch"
date: "2008-02-04"
---
A person on <a href="https://savannah.gnu.org/projects/screen/">screen-devel</a>, who shall respectfully remain nameless since he didn't post to the list proper, sent me this in response to <a href="http://www.ericgar.com/2008/02/03/screen-renumbering-windows-to-fill-gaps/">my proposed patch</a>:

> I would never use this feature because I would rather that window #n
> always remain window #n, but I can see the usefulness of the feature
> if you used to have more than 10 windows and now have fewer than 10
> but some windows still bear numbers greater than 9, so you can go
> back to using Ctrl-A &amp;lt;n&amp;gt; to switch to them quickly.
> 
> My recommendation is that you call it compacting, not renumbering.
> "renumber" doesn't make it clear enough HOW they get new numbers.


To which I replied:

> Thanks for the mail. Your comments are fair enough and definitely
> anticipated. I often can peak way above 10 windows, begrudgingly,
> and often want to migrate down to as few as possible so that the
> next window allocated is the highest available number. And, I'm an
> active user of the hardstatus line, including labeling windows. I
> implemented this because I know several people who use screen like
> I do, though knowing many people do not.
>
> I do like your suggestion that the patch feature be called
> "compacting." I was in fact struggling with how do best describe the
> action, and that does, in a single word. I will create a new patch
> sometime in the next week reflecting this, even just for my purposes.
>
> I'm curious, though, for my own usage: how do you remember which
> window is which? What is your typical screen workflow?
>
> Thanks,
> Eric
