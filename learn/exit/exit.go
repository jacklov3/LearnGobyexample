package exit

import (
	"fmt"
	"os"
)

func Test_exit(){
	defer fmt.Println("!")
	os.Exit(3)
}