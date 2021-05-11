package main

import (
	//"errors"
	"fmt"
	"net/http"
)

type urlStatus struct {
	url    string
	status bool
}

func main() {
	// A slice of sample websites
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	c := make(chan urlStatus)
	for _, url := range urls {
		go checkUrl(url, c)
	}
	result := make([]urlStatus, len(urls))
	for i, _ := range result {
		result[i] = <-c
		if result[i].status == true {
			fmt.Println(result[i].url, "is up.")
		} else {
			fmt.Println(result[i].url, "is down !!")
		}
	}
}
func checkUrl(url string, c chan urlStatus) {
	fmt.Println("control in go hello")
	_, err := http.Get(url)
	if err != nil {
		c <- urlStatus{url: url, status: false}
	}
	c <- urlStatus{url: url, status: true}
}
