package chapter8

import (
	"fmt"
	"os"
	"time"
)

func Test_countdown1(){
	fmt.Println("Commencing countdown. Press return to abort")

	//打断发射的channel
	abort :=make(chan struct{})
	go func(){
		os.Stdin.Read(make([]byte,1))//读取一个字节
		abort<- struct{}{}
	}()

	tick :=time.Tick(1*time.Second)
	for countdown :=10;countdown>0;countdown--{
		fmt.Println(countdown)
		<-tick
	}
	select {
	case <-time.After(10*time.Second):
		fmt.Println("good")
	case <-abort:
		fmt.Println("launch aborted!")
		return
	}
	fmt.Println("launch...")
}

func Test_countdown2()  {
	abort :=make(chan struct{})
	go func(){
		os.Stdin.Read(make([]byte,1))//读取一个字节
		abort<- struct{}{}
	}()
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	fmt.Println("launch...")


}

func Test_odd(){
	ch :=make(chan int ,1)
	for i:=0;i<10;i++{
		select {
		case x:=<-ch:
			fmt.Println(x)
			case ch<-i:
		}
	}
}
