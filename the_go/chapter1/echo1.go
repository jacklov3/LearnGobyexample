//Echo1 打印命令行参数
package chapter1

import (
	"fmt"
	"os"
)

func Echo1(){
	var s,sep string
	for i:=1;i<len(os.Args);i++{
		s+=sep+os.Args[i]
		sep=" "
	}
	fmt.Println(s)
}