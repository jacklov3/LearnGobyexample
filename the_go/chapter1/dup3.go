//dup3一次性读取问题的全部内容到内存再处理
package chapter1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
func Dup3(){
	counts :=make(map[string]int)
	for _,filename :=range os.Args[1:]{
		data,err :=ioutil.ReadFile(filename)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
			continue
		}
		for _,line:=range strings.Split(string(data),"\n"){
			counts[line]++
		}
	}
	for line,n:=range counts{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}
