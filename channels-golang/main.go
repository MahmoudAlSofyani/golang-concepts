package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
		"https://mahmoudalsofyani.dev",
	}

	// declare a channel of type string
	c := make(chan string)

	for _, l := range links {
		go isWebsiteLive(l, c)
	}

	// after every 5 seconds, loop the links from the channel
	// and make the http call while updating link with l
	// in the function literal (anonymous function)
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			isWebsiteLive(link, c)
		}(l)

	}
}

func isWebsiteLive(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// send the link to the channel
		c <- link
		return
	}

	// send the link to the channel
	fmt.Println(link, "is live!")
	c <- link

}
