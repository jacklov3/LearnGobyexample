package channels

import "fmt"

func Test_channels(){
	messages := make(chan string)
	go func(){messages <- "ping"}()
	msg := <-messages
	fmt.Println(msg)
}