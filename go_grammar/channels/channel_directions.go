package channels

import "fmt"

func ping(pings chan<- string,msg string){
	pings <-msg
}

func pong(pings <-chan string,pongs chan<- string){
	msg :=<-pings
	pongs <-msg
}

func Test_chan_direct(){
	pings :=make(chan string,1)
	pongs :=make(chan string,1)
	ping(pings,"passed message")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}