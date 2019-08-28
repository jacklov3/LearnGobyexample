//处理ehco和http响应
package chapter1

import (
	//"fmt"
	//"log"
	"net/http"
)

func Server3()  {
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		lissajous(w)
	})
	http.ListenAndServe("localhost:8080",nil)

}

func handler3(w http.ResponseWriter,r *http.Request){
//	fmt.Fprint(w,"%s %s %s\n",r.Method,r.URL,r.Proto)
//	for k,v :=range r.Header{
//		fmt.Fprintf(w,"Header[%q]=%q\n",k,v)
//	}
//	fmt.Fprintf(w,"Host=%q\n",r.Host)
//	fmt.Fprintf(w,"RemoteAddr=%q\n",r.RemoteAddr)
//	if err:=r.ParseForm();err!=nil{
//		log.Print(err)
//	}
//	for k,v :=range r.Form{
//		fmt.Fprintf(w,"Form[%q]=%q\n",k,v)
//	}
	lissajous(w)
}
