package chapter9

import (
	"fmt"
	"sync"
)

////保证只能有一个goroutine在同一时刻访问一个共享变量
//var(
//	sema = make(chan struct{},1)//一个二元信号量
//	balance int
//)
//
//func Deposit2(amount int,done chan string){
//	sema <- struct{}{}//获得token
//	balance = balance+amount
//	<-sema //释放token
//	done<-"done"
//}
//
//func Balance2() int{
//	sema <- struct{}{}//获得token
//	b:=balance
//	<-sema//释放token
//	return b
//}


var (
	mu sync.Mutex
	balance int
)

func Deposit2(amount int,done chan string){
	mu.Lock()
	balance = balance+amount
	mu.Unlock()
	done<-"done"
}

func Balance2() int{
	mu.Lock()
	defer mu.Unlock()
	return balance

}
func Test_bank2(){
	done := make(chan string,2)
	go Deposit2(200,done)
	go Deposit2(100,done)
	for x:=range <-done {
		fmt.Println(x)
	}
	fmt.Println(Balance2())
}