---
wordpress_id: "34"
wordpress_url: "http://www.ericgar.com/2006/10/17/eclipse-jbossas-ejb-30-setup-instructions/"
title: "Eclipse + JBossAS + EJB 3.0 setup instructions"
date: "2006-10-17"
---

Today I was tasked with getting my software engineering group's
development environment set up. The requirement for our project
is to build a three-tier application using a major component model
framework. We chose to use:

* **Java/EJB** because it is well documented and in a languagewe all
know (and since we don't want to use COM+ or .Net as we don't have
Windows boxes, the only language we could use)

* **Eclipse** because it is the best Java IDE and is cross platform

* **JBoss AS** because it is a well supported J2EE application
server and seemed to have good Eclipse support. Additionally, it is
one of the few ASes to have J2EE 3.0 support (sort of, see below).

The combination of these tools is supposed to be easy, but I
found it fairly tricky because of lack of documentation. I'm
sure it could be worse had the guys over at JBoss not worked as
hard as they have to bring a J2EE platform to the open-source
world. The rest of this post include the steps that we took to
install these tools and get them running. Hopefully they'll be
of help for those of you who got here through Google.

### Installing Eclipse, JBossAS with EJB 3.0

* **Prerequisites**: A Java 5 JRE must be installed to be able to
use JBoss's EJB 3.0.

* **Download and install JBossAS**: We are using the latest
production release of JBoss AS, version 4.0.4, which can be
[downloaded from Sourceforge][jboss]. To use the bundled EJB 3.0,
you must not download the zip file. Additionally, you must use
the graphical installer to select the EJB 3.0 packages.

   Once downloaded, execute `java -jar
jboss-4.0.4.GA-Patch1-installer.jar` to launch the installer. Select
your preferred language, click "OK," and click "Next." The installer
may freeze for a few moments; this is normal. Continue forward. Choose
a sane install directory: because I am administering my own machine,
I've installed JBossAS into "/usr/local/jboss". Choose something
that makes sense for you, but be aware I reference my own install
directory later and appropriate substitutions should be made.

   On the next panel, be sure to select "ejb3" or else you won't have
EJB 3.0 support. The default values for the remaining options were
sufficient for our purposes.

   When this completes, you now have JBossAS installed. Try to run
it by executing "bin/run.sh" in the JBossAS directory.

* **Install / Upgrade Eclipse**: Eclipse is a powerful Java
IDE that is convenient because it is free and open-source. A
plethora of plugins exist for this development platform, some
of which we'll install later. You can download Eclipse from the
[Eclipse website][eclipse]. (Alternatively, you can install
[JBoss's bundle][bundle] that includes Eclipse and some of the
plugins we'll install later. I had an existing installation of
Eclipse and would prefer to run my existing vanilla version over
the JBoss-branded one. I had originally installed Eclipse from
Gentoo's Portage system. YMMV).

   If you already have Eclipse installed, launch it and upgrade it:
Click on "Help -&amp;gt; Software Updates -&amp;gt; Find and Install" and
select "Search for updates of the currently installed features". Click
appropriate buttons to continue.

    (Caveat: if you've installed Eclipse into a global install
directory like "/usr" or "/opt", be sure to run Eclipse as root or
some user that can write to these directories or more preferably
change the install directories).

[eclipse]: http://www.eclipse.org/downloads/
[bundle]: http://labs.jboss.com/portal/jbosside/download/index.html
[jboss]: http://prdownloads.sourceforge.net/jboss/jboss-4.0.4.GA-Patch1-installer.jar?download

* **Install Eclipse Plugins**: Our team is using a Subversion
repository for source code management; the best Eclipse SVN plugin
is the new but complete Subversive from the Polarion group. This
plugin integrates SVN actions and management into the Eclipse
GUI. We also want to use the JBoss Eclipse plugins to interface our
project to the application server we just installed. This plugin
has a dependency on an upgraded version of the Eclipse EMF tool,
so we need to perform this upgrade.

   In Eclipse, Click on "Help -> Software Updates -> Find and Install"
and select "Search for new features to install". For each of the
following URLs, create a new update site by clicking the "New
Remote Site" button:

   * Subversive `http://www.polarion.org/projects/subversive/download/1.1/update-site/1`
   * Eclipse's EMF `http://www.eclipse.org/emf/updates/`
   * JBoss Eclipse IDE Plugins `http://download.jboss.org/jbosside/updates/development`

   ***Update***: Thanks to commenters below (!!), it should be
   reiterated that ***only the development JBoss plugins will
   work*** at this point in time and are the ones used for the
   remainder of this setup doc.

   Despite it being installed by default with Eclipse 3.2.1, I also
   had to install APT on one of my two installs for some reason. It can be
   found at <code>http://www.eclipse.org/jdt/apt/JdtAptUpdateSite</code> .

   Once complete, be sure that all remote sites are checked. Click
   "Finish" and install all available packages. There may be some
   additional dependencies, be sure to download those first if there
   seems to be an error.

* Create a new Eclipse Java project</strong>: We now must
begin setting up the source code locations and supporting
libraries. Create a new Java project in Eclipse by clicking
"File -> New -> Project...". Select the "EJB 3.0 Project"
item. This item is inspired by the skeleton Eclipse
project that is described and downloaded at the [JBoss
EJB 3.0 Wiki][wiki].

    Give the project a name (mine is "ejb3-project") and click
next. Now, click the "Create a JBoss Server". We must first create
a JBoss configuration In the directory tree, select the JBoss Inc /
JBoss AS 4.0 (NOT JBoss / JBoss v4.0) item. Give a name of (mine
is "default") and specify where you installed JBossAS (mine is
"/usr/local/jboss"). If this is successful, Eclipse will detect
the configuration we created when we installed JBossAS. Click next.

    Create a new JBoss Server. Give a name to your new JBoss server
(I'll call mine "Eric"). Click "Finish". Select the configuration
we just created and click "Finish." If you can't click finish,
Eclipse does not detect EJB3 libraries in the installation of
JBossAS you specified. We now have an EJB project framework. There
are still a few holes we need to patch in this folder tree.

[wiki]: http://wiki.jboss.org/wiki/Wiki.jsp?page=StarterSkeletonEclipseProject"

* **Providing necessary libraries**: We need to add some supplied
EJB library code to the CLASSPATH of our project. This allows our
project to use methods that are provided by JBoss and TestNG.

   Right click on our new project and select "Preferences". Navigate
to the "Java Build Path" pane and click "Add JARs...". Select the
"jboss-ejb3.jar" element contained within the folder structure
"ejb3-project/lib/embeddable-ejb3/". Click "OK". Now, add the
"testng-4.5.2-jdk15.jar" from within the folder structure
"ejb3-project/lib/testng/". Click "OK". Now, click "OK" on the
Properties pane to return to our project.

   The entire project will be refreshed and compiled. But there
are errors. For whatever reason, all \*.java files are placed in
the default package, despite being declared as being within a
package. For each java source file, click on the first error in
the file and then click "Move xxx to package yyy". (Alternatively,
you can do this manually by creating the appropriate package
and moving the files manually). The project should now compile
successfully.

* **Edit build.xml to your environment**: Personally, I didn't have
to edit this file. You may have to depending on your environment,
but I think it gets generated from the information supplied earlier.

* **Build and deploy**: Now, click the external executable button
and select "Build and deploy". If you configured the build.xml
file properly, this will succeed and give some status messages in
the console.

* **Start the server**: Click "Window" -> "Show View" -> "Other",
scroll down to the "Server" folder and select "JBoss Server View".
You should see the server we configured earlier. Right click your
server and click "Start". Hooray your server should start!!

   But, you'll get an error indicating that the
ManagedConnectionFactory is not yet installed. This is because you
must install a JDBC driver for either MySQL or PostreSQL in the lib/
directory of your project.

That's where I end these instructions because that's
where real development starts. I refer you to <a
href="http://trailblazer.demo.jboss.com/EJB3Trail/">the JBoss
TrailBlazer demo</a> to get a better feel for what comes next in
JBoss EJB 3.0 developement and deployment.

