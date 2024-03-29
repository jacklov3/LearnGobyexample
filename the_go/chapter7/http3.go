package chapter7

import (
	"fmt"
	"net/http"
)

func (db database)list(w http.ResponseWriter,req *http.Request){
	for item,price :=range db{
		fmt.Fprintf(w,"%s: %s\n",item,price)
	}
}

func (db database)price(w http.ResponseWriter,req *http.Request)  {
	item :=req.URL.Query().Get("item")
	price,ok :=db[item]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"no such item:%q\n",item)
		return
	}
	fmt.Fprintf(w,"%s\n",price)

}