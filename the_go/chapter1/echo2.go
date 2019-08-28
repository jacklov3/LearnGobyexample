//使用range语言解析命令行参数
package chapter1

import (
	"fmt"
	"os"
)

func Echo2()  {
	s,sep :="",""
	for _,arg:=range os.Args[1:]{//忽略索引
		s +=sep+arg
		sep=" "
	}
	fmt.Println(s)
}

