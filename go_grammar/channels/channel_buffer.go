package channels

import "fmt"

func Test_chan_buff(){
	messages :=make(chan string,2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}