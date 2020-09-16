package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// make channel
	sizes := make(chan int)
	go responseSize("https://www.msnbc.com/", sizes)
	go responseSize("https://golang.org/", sizes)
	go responseSize("https://golang.org/doc", sizes)
	go responseSize("https://git-scm.com/docs/gitignore", sizes)
	go responseSize("https://github.com/mcclayac/HeadFirstThirteen", sizes)
	go responseSize("https://learning.oreilly.com/library/view/learning-go/9781492077206/", sizes)
	go responseSize("http://headfirstgo.com/", sizes)
	go responseSize("https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/", sizes)
	go responseSize("https://twitter.com/home", sizes)
	go responseSize("https://totale.rosettastone.com/sign_in", sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
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
	//lenStr := fmt.Sprintf("%d", len(body))
	channel <- len(body)
	//fmt.Println(lenStr + "  :" + url)
}
