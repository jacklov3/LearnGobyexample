package waitgroups

import (
	"fmt"
	"time"
	"sync"
)

func worker(id int,wg *sync.WaitGroup){
	fmt.Println("worker %d starting\n",id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n",id)
	wg.Done()
}

func Test_waitgroups(){
	var wg sync.WaitGroup

	for i :=1;i<=5;i++{
		wg.Add(1)
		go worker(i,&wg)
	}
	wg.Wait()
}