package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type page struct {
	url  string
	size int
}

func main() {

	// make channel
	pages := make(chan page)
	urls := []string{
		"https://www.msnbc.com/",
		"https://golang.org/",
		"https://golang.org/doc",
		"https://git-scm.com/docs/gitignore",
		"https://github.com/mcclayac/HeadFirstThirteen",
		"https://learning.oreilly.com/library/view/learning-go/9781492077206/",
		"http://headfirstgo.com/",
		"https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/",
		"https://twitter.com/home",
		"https://totale.rosettastone.com/sign_in",
	}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.url, page.size)
	}
}

func responseSize(url string, channel chan page) {

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
	//lenStr := fmt.Sprintf("%d", len(body))
	// channel <- len(body)
	channel <- page{url: url, size: len(body)}
	//fmt.Println(lenStr + "  :" + url)
}
