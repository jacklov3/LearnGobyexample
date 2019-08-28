package chapter3

import (
	"fmt"
	"strings"
)

func basename2(s string)string  {
	slash:=strings.LastIndex(s,"/")
	s=s[slash+1:]
	if dot:=strings.LastIndex(s,".");dot>=0{
		s=s[:dot]
	}
	return s
}

func GetBasename2(s string){
	fmt.Println(basename2(s))
}
