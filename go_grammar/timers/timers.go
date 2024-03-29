package timers

import (
	"time"
	"fmt"
)

func Test_timers(){
	timer1 :=time.NewTimer(2*time.Second)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 :=time.NewTimer(time.Second)
	go func(){
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 :=timer2.Stop()
	if stop2{
		fmt.Println("Timer 2 stopped")
	}

}