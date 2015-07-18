---
wordpress_id: "234"
wordpress_url: "http://www.ericgar.com/?p=234"
title: "ICANN's decision to expand DNS"
date: "2008-06-26"
---
The world has grown up with the internet having DNS addresses (the first part of URLs, web addresess, for the lay person) with a top level domain (TLD) of things like .com, .org, .net and so on. ICANN, the non-governmental body that governs such addresses, <a href="http://www.icann.org/en/announcements/announcement-4-26jun08-en.htm">decided today that it would open up these suffixes</a> to an "application period," whereby anyone can apply for any possible TLD. This means http://icanhas.cheezburger is a real possibility.

Computer scientists and system administrators like the current system, without these extraneous, crazy TLDs, because DNS forms a very nice tree which is easy to manage. (There are specifics here that I should gloss over with the desire to not repeat how DNS actually works). Trees are a very basic computer science data structure, one whose properties are extremely well known, in conjunction with cooperating algorithms. Recursive DNS doesn't work for obvious reasons that iterative DNS does, for example. (Yet, stragely, everyone teaches that DNS works recursively. That annoyance is for another time).

This tree structure gets completely destroyed with the addition of arbitrary TLDs. This is because ICANN is no longer valuing a TLD more than a (second-level) domain name (or the part before the .com). We will have a flat structure for DNS, exactly what it was created to prevent.

My thoughts on this are three-fold:

1. The tree structure of DNS is rather nice, but arbitrary. The tree will now become a forest, in Microsoft Active Directory-speak.

2. That the DNS root servers will have to cope with more queries is a non-issue. Root DNS administrators will step up and get the job done. They want a functioning internet before they want a tree-like DNS structure.

3. The internet doesn't care. Having a naming scheme where there are far more than three letters after the last dot means nothing. DNS is for human consumption and should not purport to be any other type of addressing system. DNS is no longer a very structured system of naming a particular computer on the network; it is now tag-based. That's all.

The last point needs to be expanded. The very notion of having people remember pointers to particular machines on the internet is silly. This is effectively what you do every time you type in a URL. You first identify the protocol your computer should talk to the other computer (typically http://), then you specify a human-readable form of an identifier for a machine (ericgar.com uniquely identifies a particular machine that holds the files that make up this website). Any additional bits of the URL are to specify to the machine what you actually want to see.

Take a step back. Why should a human care about contacting a particular machine on the internet? He or she should not. They want the content. The content is what matters. The service is what matters.

I predict we will see an internet where it isn't unreasonable to search for whatever you need as a first point of contact. Want to get to slashdot? Search for news for nerds. Want to read the news from the US? Type "journalism north america" and your choices will be presented. This will take the form of a combination of Wikipedia, where there is peer-submitted, peer-reviewed links and a search engine, where the information gathering is algorithmic. 

I'd imagine that this seems foreign to most people. We're so used to specifying a URL that we take it for granted.

I think the traditional notion of DNS is going away and faster than I can imagine. I think something like google, today a celebrity, will become a very necessary tool for navigating the internet. In the future, we will consider this kind of search infrastructural, just as we do DNS.

