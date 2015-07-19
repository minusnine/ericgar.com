---
title: About
---

## about me

{{< figure src="/about/IMG_2194_lo.JPG" >}}

### contact

* **Email**: <a href="mailto:eric@ericgar.com">eric@ericgar.com</a>
* **Phone**: <a href="tel:+1-601-884-0354">601-884-0354</a>
* **Web**: <a href="http://ericgar.com">http://ericgar.com</a>
* **Google+**: <a href="https://plus.google.com/112531583832022190935">Eric Garrido</a>
* **XMPP**: <a href="xmpp:eric@ericgar.com">eric@ericgar.com</a>
* **Twitter**: <a href="http://www.twitter.com/minusnine">minusnine</a>
* **GitHub**: <a href="http://www.github.com/minusnine">minusnine</a>
* **Instagram**: <a href="https://instagram.com/minusn1ne/">minusnine</a>


#### gpg public key fingerprint

<br />

<pre>
pub   4096R/EDF93F67 2010-07-01
      Key fingerprint = <a href="/about/gpg.asc">7964 43A2 9880 1FE2 A008  824D DA84 8148 EDF9 3F67</a>
uid                  Enrique Kresla Garrido (Eric) &lt;eric@ericgar.com&gt;
sub   4096R/01462932 2010-07-01
</pre>


### ethos

Most of the problems we have are interesting in that they are questions
looking for answers in a system of composed rules. Solving the problem
is just a matter of solving at least a subset of the system or adding
additional constraints to it.

For me, computers and computation represent an interesting microcosm
of the available problem spaces in the universe. This is a system
that, to anyone but the materials scientist or electrical engineer,
is entirely constructed by humans. We decided to have a given number
of registers in the CPU. We decided to limit ourselves to a particular
transmission limit on the local network link. We meticulously designed
the languages we use to express imperatives to the machines we want
to control. 

And yet despite these limits, computers *seem* limitless in the value
we derive from them.

The part that interests me about all of this is that the blueprints
to the entire computer system are available to solve the problem
at hand. If knowing how the memory caches are interconnected helps
optimized accesses, that information is available. If we need to know
how `gcc` arranges the stack after optimization, that information is
just a `git clone` away.

This is why I'm such an enthusiast for and advocate of open source
software and open standards. The knowledge I've amassed over the years
is *only* possible because the source code or protocol specifications
for most of the things I've worked on is freely available and easily
accessible. For those systems that are not open, the systems on which
they sit on _are_, allowing for a black-box inspection.

The distinguishing skill of a systems programmer and administrator
is not _only_ being able to know where to obtain the copious amount
of information available, is not _only_ being able to quickly learn
and understand it, but is *not* having the fear to delve deep into it
in order to extract facts needed for solving the system. This skill
is often far more important than the specific knowledge base already
accumulated, as the systems programmer and administrator will be able
to learn whatever tool necessary to complete the job.

That's what I do and enjoy: taking apart systems and putting them back
together[^legos] (virtually or otherwise) in order to understand as
much as I need to solve the problem I'm given.[^takeapart]

The one downside is that this requires somewhat of a jack-of-all-trades
approach to computer systems. I think one day I would like to
specialize into being a subject matter expert in a particular subsystem
(say, one of the Linux kernel subsystems, high-speed interconnects,
or something of that nature). For now, I'm having fun investigating
whatever I can and enjoying the fact that I have the liberty to spend
time doing it.

[^legos]: Maybe this results from being a huge user of legos as a kid...

[^takeapart]: One recent example: a colleague and I were working
    on Oracle security integration into an existing, rich Kerberos infrastructure
    when we found an issue where the (non-`mitkrb5`) library Oracle uses for
    Kerberos operations rejected the valid service ticket presented by the
    client. The issue was that the ticket did not include one of the three
    available IP addresses available on the client; this being a virtual address
    that accepts traffic on either physical interface. The question then became:
    given a valid address-based Kerberos TGT, can one obtain address-*less*
    service tickets? We observed that forwarding a address-based TGT through SSH
    produced an address-less TGT on the target host. We subsequently navigated
    both the `openssh` and `mitkrb5` code base to investigate the issue. This was
    fun.

### biographical

I grew up in a sleepy town in Upstate New York called Rhinebeck,
attending 13 years of public school, not getting into trouble, and
always doing my homework (and then some). I spent a solid year of my
childhood coming back from school and watching two out of three *Star
Wars* movies before starting my homework.[^homework]

[^homework]: This is an example of why the US is falling behind in
    the global education race. But, the benefit of this particular example
    is that I probably can recite each script in its entirety.

My father worked at IBM for over 30 years and consequently had copious
amounts of computer hardware lying around the house. It was only a
matter of time before I started hacking on a bunch of Model Ms.

I had the benefit of spending most of my senior year of high school
at Bard College, taking science and math classes at a hugely reduced
tuition rate. Being exposed to advanced ideas and creative people was
a very formative experience for my passion for computer science and
[overall] [geekiness].

[overall]: /2010/05/17/just-did-this/
[geekiness]: /2009/11/04/a-real-but-unfinished-journal-entry

I attended the engineering undergraduate school at Columbia University
where I studied computer science. I was planning on studying applied
physics, until I had a much needed existential crisis and decided to
follow my intuition and engorged myself in computer science. I
under-appreciated having so much time to hack on and investigate new
systems; my perspective was at the 1.5-meter level when it should have
been at 10,000 meters. Still, it was a [super valuable experience][xp].

[xp]: /2009/10/29/columbia-advice

I attribute my work ethic to having a short career in technical
theater, starting in high school and ending in college. I was fortunate
to gain a "do everything necessary" attitude to getting the task
at hand completed. I transitioned to work for the university's IT
group, working on web applications ranging from employee scheduling
to webmail once I knew that computing would be my future.

I currently work as a Site Reliabilty Engineer at Google in New York
on our offline storage system. I love my job. I get to work with
really smart people on improving the reliability and operation of a
massive system that includes robots.
