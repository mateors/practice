# practice
Golang practice

## Move all files except one
> cut evrything from current directory and paste into stdinout/ directory \
> mv !(stdinout) stdinout/

### another way
> mv SOURCE_DIRECTORY/^UNWANTED_FILENAME TARGET_DIRECTORY

## What is the use of carriage return?
A carriage return, sometimes known as a cartridge return and often shortened to CR, <CR> or return, is a control character or mechanism used to reset a device's position to the beginning of a line of text.


## SPF TXT record for other domain
> "v=spf1 +a +mx +a:host.mateors.com -all" \
> "v=spf1 include:_spf.google.com ~all"

"v=spf1 +a +mx +a:host.mateors.com -all"
"v=spf1 +a +mx +a:host.mateors.com -all"
"v=spf1 +a +mx +a:host.mateors.com -all"

_spf.google.com

"v=spf1 include:_spf.google.com ~all"

"v=spf1 include:_spf_ipv4.netflix.com include:_spf.google.com include:amazonses.com include:servers.mcsv.net -all"

"v=spf1 include:spf1.amazon.com include:spf2.amazon.com include:amazonses.com -all"

v=spf1 a mx ip4:xxx.xxx.xxx.xxx -all

v=spf1 a MX include:spf.sampledomain.com ~all

_spf.google.com.	95	IN	TXT	
"v=spf1 include:_netblocks.google.com include:_netblocks2.google.com include:_netblocks3.google.com ~all"

"v=spf1 ip4:172.217.0.0/19 ip4:172.217.32.0/20 ip4:172.217.128.0/19 ip4:172.217.160.0/20 ip4:172.217.192.0/19 ip4:172.253.56.0/21 ip4:172.253.112.0/20 ip4:108.177.96.0/19 ip4:35.191.0.0/16 ip4:130.211.0.0/22 ~all"

Ideally your third parties will already have a generic SPF record and you can just add the include:spf.thirdparty.dom element to your record. If they don't you might well want to create your own record for them and chain it youself anyway, so that it's easy for you manage administratively.

## Reference
* https://www.baeldung.com/linux/move-all-files-except-one
* [SPF Record tutorial](https://serverfault.com/questions/734297/adding-an-spf-record-for-a-3rd-party-but-dont-have-one-for-my-own-domain/734308)