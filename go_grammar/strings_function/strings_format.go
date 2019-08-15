package strings_function

import (
	"fmt"
	"os"
)

type point struct{
	x,y int
}

//格式化字符串
func Test_strings_format(){
	
	p := point{1,2}
	//打印结构体
	fmt.Printf("%v\n",p)
	//打印包括字段的结构体
	fmt.Printf("%+v\n",p)
	//源代码表示
	fmt.Printf("%#v\n",p)
	//变量类型
	fmt.Printf("%T\n",p)

	//布尔
	fmt.Printf("%t\n",true)

	//十进制
	fmt.Printf("%d\n",123)
	//二进制
	fmt.Printf("%b\n",14)
	//字符
	fmt.Printf("%c\n",33)
	//十六进制
	fmt.Printf("%x\n",456)

	fmt.Printf("%f\n",52.1)
	fmt.Printf("%e\n",123400000.0)
	fmt.Printf("%E\n",123400000.0)

	fmt.Printf("%s\n","\"string\"")
	//原字符显示
	fmt.Printf("%q\n","\"string\"")

	//字节显示
	fmt.Printf("%x\n","hex this")
	//指针
	fmt.Printf("%p\n",&p)
	//宽度，左对齐，右对齐
	fmt.Printf("|%-6d|%6d|\n",12,345)

	fmt.Printf("|%6.2f|%6.2f|\n",1.2,3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n",1.2,3.45)
	fmt.Printf("|%6s|%6s|\n","hello","world")
	fmt.Printf("|%-6s|%-6s|\n","hello","world")

	//返回一个字符串
	s:=fmt.Sprintf("a %s","string")
	fmt.Println(s)
	//打印到文件流
	fmt.Fprintf(os.Stderr,"an %s\n","error")





}