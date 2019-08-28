package main

import (
	//"learn_go/the_go/chapter1"
	//"learn_go/the_go/chapter2"
	"learn_go/the_go/chapter3"
)

func main(){
	s:=[]string{"a/b/c.go","c.d.go","abc"}
	for _,str:=range s{
		chapter3.GetBasename2(str)
	}
}

