package channels

import (
	"fmt"
	"time"
)

func worker(done chan bool){
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <-true
}

func Test_worker(){
	done :=make(chan bool,1)
	go worker(done)
	<-done
}