package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//* for i := 0; i < len(links); i++ {
	//* 	fmt.Println(<-c)
	//* }

	//* for l := range c {
	//* 	time.Sleep(5 * time.Second) // This will block channel
	//* 	go checkLink(l, c)
	//* }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

	// This extra channel will make Main routine go sleep
	//* fmt.Println(<-c)

}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}
