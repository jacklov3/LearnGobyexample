package goroutines

import "fmt"

func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
	}
}

func Test_goroutines(){
	f("direct")
	go f("goroutine")

	go func(msg string){
		fmt.Println(msg)
	}("going")
	fmt.Scanln()
	fmt.Println("done")
}