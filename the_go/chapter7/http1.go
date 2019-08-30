package chapter7

import (
	"fmt"
	"log"
	"net/http"
)

func Test_http1(){
	db :=database{"shoes":50,"socks":5}
	//mux:=http.NewServeMux()
	//mux.Handle("/list",http.HandlerFunc(db.list))
	//mux.Handle("/price",http.HandlerFunc(db.price))
	//log.Fatal(http.ListenAndServe("localhost:8000",mux))
	http.HandleFunc("/list",db.list)
	http.HandleFunc("/price",db.price)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

type dollars float32
func (d dollars)String() string{return fmt.Sprintf("$%.2f",d)}
type database map[string]dollars

//func (db database) ServeHTTP(w http.ResponseWriter,req *http.Request){
//	for item,price :=range db{
//		fmt.Fprintf(w,"%s:%s\n",item,price)
//	}
//}