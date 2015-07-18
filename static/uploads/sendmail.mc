# The Following removes any comments placed between the two diverts. (meaning: 
# this is useless, but common practice.)
divert(-1)
divert(0)

# Includes the default sendmail configuration file
include(`/usr/share/sendmail-cf/m4/cf.m4')

# The following three Lines are required for all sendmail configs
VERSIONID(`$Id: sendmail-procmail.mc,v 1.2 2004/12/07 01:59:31 g2boojum Exp $')
OSTYPE(linux)
DOMAIN(generic)

# Provide a separate map for client side authentication information.         
# This is used for SMTP authentication. My authinfo file contains one line:
# 	AuthInfo:<smtp hostname> "I:<username>" "P:<password>"
# See: http://www.sendmail.org/m4/smtp_auth.html
FEATURE(`authinfo',`hash /etc/mail/authinfo.db')

# Prevents the creation an MSA with the default options
# See: http://www.sendmail.org/m4/features.html
FEATURE(`no_default_msa')

# Creates an MSA with new options
# See: http://www.sendmail.org/m4/tweaking_config.html
DAEMON_OPTIONS(`Name=MSA, Port=587, Addr=127.0.0.1, M=E')

# This tells sendmail which hosts to accept mail from. I only want
# localhost to be able to connect to my smtp server, so that I don't
# have to deal with being an open relay that will only send mail through
# my secure account. The file indicated has the following:
# localhost.localdomain           RELAY
# localhost                       RELAY
# 127.0.0.1                       RELAY
FEATURE(`access_db',`hash -T<TMPF> -o /etc/mail/access.db')dnl

# Means that you can send mail to local users.
# See: http://www.sendmail.org/m4/mailers.html
MAILER(local)

# Means you can deliver mail to other servers
# through SMTP traffic
MAILER(smtp)

# Makes processing through procmail possible
MAILER(procmail)

# If recipient addresses are fully qualified to another
# host, send mail through this server.
#See: http://www.sendmail.org/m4/masquerading_relaying.html
define(`SMART_HOST',`send-vif.cc.columbia.edu')

# Creates an MTA with non-default options (listen on localhost only)
# See: http://www.sendmail.org/m4/tweaking_config.html
DAEMON_OPTIONS(`Port=25,Addr=127.0.0.1, Name=MTA')

# The following four lines allow sendmail to negotiate SMTP AUTH with 
# the external mail servers. In my case, this is only to my ISP's SMTP
# AUTH server.
# See: http://www.sendmail.org/m4/tweaking_config.html
define(`confCACERT_PATH',`/etc/ssl/certs')
define(`confCACERT',`/etc/ssl/certs/ca-certificates.crt')

# You must manually generate these files so that your server has
# a self-signed certificate to present to the external SMTP AUTH server.
# See: http://sial.org/howto/openssl/self-signed/
# See: http://www.sendmail.org/m4/tweaking_config.html
# See: http://www.sendmail.org/~ca/email/starttls.html
define(`confSERVER_CERT',`/usr/share/ca-certificates/server.pem')
define(`confSERVER_KEY',`/usr/share/ca-certificates/server.pem')

# List the available authentication mechanisms presented to clients.
# This is by default "GSSAPI KERBEROS_V4 DIGEST-MD5 CRAM-MD5"
define(`confAUTH_MECHANISMS', `LOGIN PLAIN')

# This is strange. The 'A' means "Use the AUTH= parameter for the MAIL FROM
# command only when authentication succeeded. This can be used as a workaround 
# for broken MTAs that do not implement RFC 2554 correctly.".
# Then, the 'p' means "don't permit mechanisms susceptible to passive
# dictionary attack" and 'y' means "don't permit mechanisms that allow #
# anonymous login".
# According to the third URL below, "Next line stops sendmail from allowing 
# auth without encryption"
# See: http://www.sendmail.org/m4/tweaking_config.html
# See: http://lists.freebsd.org/pipermail/freebsd-questions/2003-June/008118.html
# See: http://news.umailcampaign.com/message/106448.aspx
define(`confAUTH_OPTIONS', `A p y')
TRUST_AUTH_MECH(`LOGIN PLAIN')

# Tells sendmail how verbose it should be for logging. For a long time, I had
# mine set at 40 to get tons of output and diagnose issues. The default is
# 9.
define(`confLOG_LEVEL', `9')

