package line_filters

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Test_line_filters(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl :=strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err();err !=nil{
		fmt.Fprintln(os.Stderr,"error:",err)
		os.Exit(1)
	}
}