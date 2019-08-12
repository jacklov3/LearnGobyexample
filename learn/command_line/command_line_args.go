package command_line

import (
	"os"
	"fmt"
)

func Test_command_line_args(){
	argsWithProg := os.Args
	argsWithoutProg :=os.Args[1:]
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}