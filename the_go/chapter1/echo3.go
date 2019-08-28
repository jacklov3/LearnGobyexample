package chapter1

import (
	"fmt"
	"strings"
	"os"
)

func Echo3()  {
	//使用strings包的join函数来高效处理多数据量
	fmt.Println(strings.Join(os.Args[1:]," "))
	//fmt.Println(os.Args[1:])

}
