//打印重复的行和计数器，从文件或标准输入中读取
package chapter1

import (
	"bufio"
	"fmt"
	"os"
)

func Dup2(){
	counts :=make(map[string]int)
	files :=os.Args[1:]
	if len(files)==0{
		countLines(os.Stdin,counts)
	}else{
		for _,arg:=range files{
			f,err :=os.Open(arg)
			defer f.Close()
			if err !=nil{
				fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
				continue
			}
			countLines(f,counts)
		}
	}
	for line,n:=range counts{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}

func countLines(f *os.File,counts map[string]int){
	input :=bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
	}
}