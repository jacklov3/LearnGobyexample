package chapter7

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period",1*time.Second,"sleep period")

func Test_sleep(){
	flag.Parse()
	fmt.Printf("Sleeping for %v...",*period)
	time.Sleep(*period)
	fmt.Println()
}
