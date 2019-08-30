package chapter8

import (
	"log"
	"net"
)

func Test_clock(){
	listener,err :=net.Listen("tcp","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	//无限循环
	for{
		conn,err :=listener.Accept()
		if err !=nil{
			log.Print(err)
			continue
		}
		go handleConn(conn)//一次只处理一次请求
	}
}

//func handleConn(c net.Conn){
//	defer c.Close()
//	for{
//		_,err:=io.WriteString(c,time.Now().Format("15:04:05\n"))
//		if err!=nil{
//			return
//		}
//		time.Sleep(1*time.Second)
//	}
//
//}
