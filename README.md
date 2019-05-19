# go-jamcracker
a simple [animal jam](https://animaljam.com) brute force password cracker using concurrency written in [golang](https://golang.org).

## disclaimer
**DO NOT ACTUALLY USE THIS UTILITY TO CRACK ACCOUNTS** - you will most likely get banished **permanently** from jamaa and have all your rare long spikes, headdresses, and beta tails stripped away.  i originally wrote this application with my daughter [@mandarinp](https://github.com/mandarinp) to teach her some programming + security basics, and to demonstrate how easy it is to bootstrap useful applications in go.  this was all originally done in our private repo, but i have decided to make the utility public and comment it accordingly for aspiring young coders like my daughter to follow along and hopefully travel down the path of True Ultimate Power.

## installation
this application makes use of the golang standard libraries.
set up your golang environment and run `go get github.com/realytcracker/go-jamcracker`.
the binary will be located in `~/go/bin`.

## known issues
windows builds may fire back an `unexpected EOF` error during the `POST` operation when submitting a potential username and password combination.  this error has been discussed in other contexts on the Holy StackOverflow, and according to the Divine Verses Within, i suspect it is triggered because floating the `Connection: close` header is not sufficient to ensure the stream is terminated; `req` itself must contain `Close = true` so the connection does not get mistakenly reused.

again, this is purely conjecture, and i could totally be wrong.

workaround: run the application on linux or osx, both of which i have tested myself with success.

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
good wordlists can be found in [@danielmiessler](https://github.com/danielmiessler)'s [SecLists](https://github.com/danielmiessler/SecLists) repository.

## license
MIT
