package web_basic

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)
func sayhelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v :=range r.Form{
		fmt.Println("key :",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello astaxie!")
}
func login(w http.ResponseWriter,r *http.Request){
	fmt.Println("method:",r.Method)//获取请求的方法
	if r.Method == "GET"{
		t,_:=template.ParseFiles("./web_basic/login.gtpl")
		t.Execute(w,nil)
	}else{
		r.ParseForm()
		fmt.Println("password:",r.Form["password"])
		fmt.Println(r.Form["username"][0])
		fmt.Fprintf(w,r.Form["username"][0])
	}
}

func upload(w http.ResponseWriter,r *http.Request){
	fmt.Println("method:",r.Method)
	if r.Method == "GET"{
		crutime :=time.Now().Unix()
		h :=md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token :=fmt.Sprintf("%x",h.Sum(nil))
		t,_:=template.ParseFiles("./form/upload.gtpl")
		t.Execute(w,token)
	}else {
		r.ParseMultipartForm(32<<20)
		file,handler,err :=r.FormFile("uploadfile")
		if err!=nil{
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",handler.Header)
		f,err :=os.OpenFile("./test/"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err !=nil{
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f,file)
	}
}
func MyHttp(){
	http.HandleFunc("/",sayhelloName)
	http.HandleFunc("/login",login)
	http.HandleFunc("/upload",upload)
	err :=http.ListenAndServe(":9090",nil)
	if err !=nil{
		log.Fatal("ListenAndServe: ",err)
	}
}
