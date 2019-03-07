# go-jamcracker
a simple [animal jam](https://animaljam.com) brute force password cracker using concurrency written in [golang](https://golang.org).

## disclaimer
**DO NOT ACTUALLY USE THIS UTILITY TO CRACK ACCOUNTS** - you will most likely get banished **permanently** from jamaa and have all your rare long spikes, headdresses, and beta tails stripped away.  i originally wrote this application with my daughter @mandarinp to teach her some programming + security basics, and to demonstrate how easy it is to bootstrap useful applications in go.  this was all originally done in our private repo, but i have decided to make the utility public and comment it accordingly for aspiring young coders like my daughter to follow along and hopefully travel down the path of True Ultimate Power.

## installation
this application makes use of the golang standard libraries.
set up your golang environment and run `go install github.com/realytcracker/go-jamcracker`.
the binary will be located in `~/go/bin`.

## usage
```
$ ./go-jamcracker -h
  ,---.       ,--.
 /  O  \      |  | go-jamcracker
|  .-.  |,--. |  | by ytcracker and mandarinp
|  | |  ||  '-'  / for educational purposes only
`--' `--' `-----' 
Usage of ./go-jamcracker:
  -c string
    	path to file for saving successful cracks (default "cracks.txt")
  -l string
    	path to file containing passwords (default "passwords.txt")
  -p string
    	path to file containing HTTP proxies in ip:port format (optional)
  -t uint
    	amount of simultaneous cracking threads (default 10)
  -u string
    	path to file containing usernames (default "usernames.txt")
```
good wordlists can be found in @danielmiessler's [SecLists](https://github.com/danielmiessler/SecLists) repository.

## license
MIT