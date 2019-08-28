//获取URL
package chapter1

import (
	"fmt"
	"io"
	//"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Fetch(){
	for _,url :=range os.Args[1:]{
		resp,err :=http.Get(url)
		if err !=nil{
			fmt.Fprintf(os.Stderr,"fetch:%v\n",err)
			os.Exit(1)
		}
		//b,err :=ioutil.ReadAll(resp.Body)
		//resp.Body.Close()
		//if err!=nil{
		//	fmt.Fprintf(os.Stderr,"fetch:reading %s: %v\n",url,err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s",b)
		fmt.Println(resp.Status)
		//使用io.Copy直接将内容拷贝到标准输出，而不需要一个b缓存
		if _,err :=io.Copy(os.Stdout,resp.Body);err!=nil{
			log.Fatal(err)
		}
	}
}