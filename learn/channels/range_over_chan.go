package channels

import "fmt"

func Test_range_over_chan(){
	queue :=make(chan string,2)
	queue <- "one"
	queue <-"two"
	close(queue)

	for elem:=range queue{
		fmt.Println(elem)
	}
}