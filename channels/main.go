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
		"http://go.org",
		"http://amazon.com",
	}

	c := make(chan string)

	// for _, link := range links {
	// 	go checkLink(link, c)
	// }

	// Anonymous Function
	for _, link := range links {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(link)

	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yup its up"
}
