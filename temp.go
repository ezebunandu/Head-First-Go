package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)
	go responseSize("https://example.com", channel1)
	fmt.Println(<-channel1)
	go responseSize("https://golang.org/", channel2)
	fmt.Println(<-channel2)
	go responseSize("https://golang.org/doc", channel3)
	fmt.Println(<-channel3)
}

func responseSize(url string, channel chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- len(body)
}
