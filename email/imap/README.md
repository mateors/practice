# IMAP - Internet message Access Protocol
RFC-3501, Old 802

## IMAP
* Connect to an email account
* List all of the available folders in the mailbox
* Download a message


http://bahut.alma.ch/2015/11/roundcube-webmail-with-sqlite-on-debian.html

## How to read email using openssl command

> openssl s_client -crlf -connect mostain.net:993

> tag login postmaster@mostain.net mateors321

> tag LIST "" "*"

> tag SELECT INBOX

> tag STATUS INBOX (MESSAGES)

> tag FETCH 6378:6388 (BODY[HEADER])

> tag FETCH 1:1 (BODY[HEADER])

> tag FETCH 1 (BODY)

> tag FETCH 1 (BODY[1])

> tag LOGOUT

> openssl s_client -connect imap.gmail.com:993 -crlf 


## IMAP commands generally look like this:
> <tag> <command> [<arg1><arg2>…]

## tag1 CAPABILITY
This tagging ability means that it’s possible for a server to handle more than one request at the same time from a client on the same connection, and to indicate their completion by sending back the appropriate tag. In practice, many clients don’t support this ability to send concurrent requests, and simply block waiting for data to arrive on the open socket after a command is sent.

## Resource
* https://www.nylas.com/blog/nylas-imap-therefore-i-am/
* https://tewarid.github.io/2011/05/10/access-imap-server-from-the-command-line-using-openssl.html
