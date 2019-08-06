package tickers

import (
	"time"
	"fmt"
)

func Test_tickers(){
	ticker :=time.NewTicker(500*time.Millisecond)
	go func(){
		for t:=range ticker.C{
			fmt.Println("Tick at",t)
		}
	}()

	time.Sleep(1600*time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}