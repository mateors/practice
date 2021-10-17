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
> "v=spf1 +a +mx +a:host.mateors.com -all"

## Reference
* https://www.baeldung.com/linux/move-all-files-except-one