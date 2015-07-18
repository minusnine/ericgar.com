---
wordpress_id: "328"
wordpress_url: "http://www.ericgar.com/?p=328"
title: "Samsung 3253H LCD television recalibration"
date: "2008-11-27"
---

I'm among the huge (and growing) number of people who [really like
to take things apart][make] or to further investigate hidden options
just to know how things work. (One might call this characteristic
"intellectual curiosity" aka *the "break shit open and put it back
together again" syndrome*).

[make]: http://blog.makezine.com"

The other day, I discovered the service menu on my one year old
[Samsung LCD television][tv]. It was glorious: gone are the
terms "color", "brightness", and "tint" for controlling how the
picture looks. In are the terms "LVDS_TX_Bit", "Nor_Roffset", and
"SVP-PX", which clearly are more powerful simply because they are
unintelligible. **"I will attract sexy women with a better picture
using advanced settings!"**, I exclaimed! [^1]

[^1]: Well, I didn't actually say or think that. In reality, I
knew that watching old episodes of Battlestar Galactica alone might
look better.

[tv]: http://reviews.cnet.com/flat-panel-tvs/samsung-ln-t3253h/4869-6482_7-32306314.html?messageSiteID=7&amp;messageID=2408576&amp;cval=2408576;2425005&amp;ctype=msgid;cmsgid&amp;commentMessageID=2425005

Among the new menu items is "Calibration". I figured,
"Sure! Calibration is a good thing! I want my TV to be calibrated!" and
clicked it. Over the next few minutes, my television scanned various
hues in an attempt to find the right color palate for my viewing
pleasure.

But then *everything* was tinted way red. I mean, even *grays* were
showing up reddish. It was bad. So bad that using the normal "tint"
and "color" parameters couldn't fix it. The default factory reset
yielded no ground to my new tint overlords. The service menu factory
reset did not do any justice either. It was bad.

Then I read on the [Samsung TV FAQ on the Service Menu][menu] the following:

[menu]: http://samsungplasmatvfaq.com/index.php/Service_Menu

> Calibration: STAY OUT. DO NOT TOUCH. You have been warned, this is
> designed to be run in the factory or onsite with specialized hardware
> outputting specific images. You can ruin your set if you use this.

**Damn**. Truer words may never have been written.

[Various] [people] [around] [the internet] say that getting a
professional calibrator is the only real way out of that, which
can cost around $300, about half or more of the value of the
equipment. Then I read that all Samsung does is feed a pattern from
a [signal generator] \(which both makes sense and can't actually be
accurate since we're talking about a digital HDMI input\), which means
that a layperson should be able reproduce the calibration image.

[various]: http://www.avsforum.com/avs-vb/showthread.php?t=890457
[people]: http://www.avsforum.com/avs-vb/archive/index.php/t-777497.html
[around]: http://answers.yahoo.com/question/index?qid=20071227072724AALKt6h
[the internet]: http://en.wikipedia.org/wiki/Series_of_tubes
[signal generator]: http://en.wikipedia.org/wiki/Waveform_generator

But what is the image? [Color
tones][tones]? [Grayscale][gray]? [Boxes][boxes]? [A random internet
meme][meme]?

[tones]: http://www.asapproductions.com/images/ColorBars2.gif
[gray]: http://www.hometheaterhifi.com/images/stories/march-2008/panasonic-ag-hvx200-camera-gray-scale-test-for-display.jpg
[boxes]: http://sethresnick.com/photographersonly/srmonitor4.jpg
[meme]: http://www.badgerbadgerbadger.com/

I tried color tones first for no reason other than they seem common enough for television calibration. Colors returned closer to normal, but still wrong.

Some other guy on the internet[^int] says that a Samsung support
guy told him that they use a checkerboard pattern. So, I used [a
checkerboard pattern][pattern] and restarted calibration: colors are
(more or less) back to normal. There is some ambiguity because I
can't really compare the before and after, and I feel like blacks
were blacker before.

[^int]: who should obviously be trusted
[pattern]: /uploads/2008/11/checker.jpg 

But, *at least Jessica Alba looks human-like now.*

My advice? **Don't mess with the service menu unless you actually know
what you're doing**, regardless of level of intellectual curiosity.[^that]

[^that]: That said, you can get to it on Samsung TVs by pressing
"Mute", 1, 8, 2, "Power", one after the other. Caveat emptor.
