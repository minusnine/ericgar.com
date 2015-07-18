---
wordpress_id: "180"
wordpress_url: "http://www.ericgar.com/2008/01/19/psa-https-and-gmail/"
title: "PSA: https and gmail"
date: "2008-01-19"
---
Those of you who know me know that I don't use Gmail for a variety
of reasons. But, I know you do. Here's looking at you, kid.

When you go to [http://gmail.com][gmail] to login, your browser greets
you with a happy "this webpage is secure" notification.

![baby2](http://www.ericgar.com/uploads/2008/01/happy_baby.jpg)

And you sign in. Your username and password is sent using an encryption
technology called SSL/TLS so that people who see your information go
by can't actually read it.

![alice](http://www.ericgar.com/uploads/2008/01/alice.jpg)

Google then sends you, over the same encrypted connection, a
**delicious cookie** to identify you so that you don't have to sign
in every time you request something from them.

![baby](http://www.ericgar.com/uploads/2008/01/baby_cookie.jpg)

This is all standard practice. But then Google does something
sneaky. It redirects you to the **non-encrypted** version of Gmail.

![unecr](http://www.ericgar.com/uploads/2008/01/unencrgoogle.png)

All subsequent information you retrieve is sent over the internet
unencrypted, available for any **eavesdropper**[^1]
to see.

![eva longoria](http://www.ericgar.com/uploads/2008/01/evalongoria.jpg)

This is particularly important when you're browsing over an untrusted
network, like the wireless network at Starbucks, the connection you
happen to use on a park bench, or even **my wireless network** when
you come to my apartment (where I may or may not log packets).

![wifi](http://www.ericgar.com/uploads/2007/06/leo.png)[^2]

Now, we all know that you don't want your correspondence with the
new **half-orc you met at the Friday Dungeons and Dragons session**
to be known to the world.

Worse than anyone being able to see everything you send back and forth
to Google, the eavesdropper could intercept the delicious cookie,
install it in their browser, and impersonate you. They would have
complete access to all of your information at Google.

There is a simple fix to avoid this potential embarrassment,
however cute the half-orc may in fact be. **Instead of going to
[http://gmail.com][gmail], use [https://gmail.google.com][gmails]**
which will encrypt everything you send and receive to and from Google.

[gmail]: http://gmail.com
[gmails]: https://gmail.google.com

Remember, your love life is counting on it.

![orc](http://www.ericgar.com/uploads/2008/01/orc.jpg)

[^1]: "Alice" is the name used for the [unassuming victim of computer security][wiki]. "Eve" is the typical name for the "eavesdropper."
[wiki]: http://en.wikipedia.org/wiki/Alice_and_Bob

Picture of <a href="http://flickr.com/photos/kcorbyndyc/36666332/">happy baby</a> by <a href="http://flickr.com/photos/kcorbyndyc/">cnbyates</a>.

Picture of <a href="http://flickr.com/photos/trommetter/697551/">cookie baby</a> by <a href="http://flickr.com/photos/trommetter/">Jason Trom</a>. 

Picture of <a href="http://flickr.com/photos/steature/38853035/">Eva Longoria</a> by <a href="http://flickr.com/photos/steature/">steature</a>.

Picture of <a href="http://flickr.com/photos/christajoy42/290276586/">Orc Donny</a> by <a href="http://flickr.com/photos/christajoy42">cristajoy42</a>

All are licensed under <a href="http://creativecommons.org/licenses/by-nc/2.0/deed.en">CC Attribution-Noncommercial 2.0 Generic</a>.
