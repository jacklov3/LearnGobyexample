package worker_pools

import (
	"fmt"
	"time"
)

func worker(id int,jobs <-chan int,result chan<- int){
	for j :=range jobs{
		fmt.Println("worker",id,"started job",j)
		time.Sleep(time.Second)
		fmt.Println("worker",id,"finished job",j)
		result <-j*2
	}
}

func Test_worker(){
	jobs := make(chan int,100)
	results := make(chan int,100)

	for w:=1;w<=3;w++{
		go worker(w,jobs,results)
	}
	for j:=1;j<=5;j++{
		jobs <-j
	}
	close(jobs)
	for a:=1;a<=5;a++{
		<-results
	}
}