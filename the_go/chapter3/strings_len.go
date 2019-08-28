package chapter3

import (
	"fmt"
	"unicode/utf8"
)

func Lenofstring(){
	s:="Hello,世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i,r:=range s{
		fmt.Printf("%d\t%q\t",i,r)
	}
	n :=0
	for range s{
		n++
	}
	fmt.Println(n)
}

