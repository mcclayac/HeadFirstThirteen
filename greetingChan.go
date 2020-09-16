package main

import "fmt"

func gretting(myChannel chan string) {
	myChannel <- "Theze Nuts"
}

func main() {
	myChannel := make(chan string)
	go gretting(myChannel)
	//fmt.Println(<-myChannel)
	recievedValue := <-myChannel
	fmt.Println(recievedValue)
}
