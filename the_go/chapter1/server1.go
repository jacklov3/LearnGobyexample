//最小的echo服务器
package chapter1

import (
	"fmt"
	"log"
	"net/http"
)

func Server1(){
	http.HandleFunc("/",handler1)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}
func handler1(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"URL.Path=%q\n",r.URL.Path)
}