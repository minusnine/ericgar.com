---
wordpress_id: "512"
wordpress_url: "http://ericgar.com/?p=512"
title: "Auto-Dependency Generation in GNU Make"
date: "2009-09-19"
---
I renovated some makefiles this week and found myself wanting to generate dependencies of a piece of code automatically from 'gcc -M'. The existing Makefile had an explicit "make depends" step that would invoke 'gcc -M' for each of the source files. This is annoying because all generated Makefiles would have to be recreated if one file's dependencies changed.

The <a href="http://www.gnu.org/software/make/manual/make.html">GNU Make Manual</a> has a section on <a href="http://www.gnu.org/software/make/manual/make.html#Automatic-Prerequisites">automatically generating prerequisites</a>, which only gets you part of the way.

I owe a good deal of gratitude to Tom Tromey and Paul D. Smith for creating and writing up <a href="http://make.paulandlesley.org/autodep.html">instructions on how GNU Make can automatically generate dependencies for source files</a>. They discuss even better ways of generating dependency files and build upon the GNU Make Manual method.

