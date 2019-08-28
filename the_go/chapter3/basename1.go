package chapter3

import (
	"fmt"
)

func basename(s string)string{
	for i:=len(s)-1;i>=0;i--{
		if s[i]=='/'{
			s=s[i+1:]
			break
		}
	}

	for i:=len(s)-1;i>=0;i--{
		if s[i]=='.'{
			s=s[:i]
			break
		}
	}
	return s
}

func GetBasename(s string){
	fmt.Println(basename(s))
}
