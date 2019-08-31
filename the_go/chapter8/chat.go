package chapter8

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Test_chat(){
	listener,err :=net.Listen("tcp","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	go broadcaster()
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			log.Print(err)
			continue
		}
		go handleConn2(conn)
	}
}

type client chan<- string//类型别名定义
var (
	entering = make(chan client)//连接的客户端消息
	leaving = make(chan client)//离开的客户端
	messages = make(chan string)//消息
)

func broadcaster(){
	clients :=make(map[client]bool)//客户端字典集合
	for{
		select {
		case msg:=<-messages://收到消息
			for cli:=range clients{
				cli <-msg
			}

			case cli:=<-entering:
				clients[cli]=true
			case cli :=<-leaving:
				delete(clients,cli)
				close(cli)
		}
	}
}

func handleConn2(conn net.Conn){
	ch :=make(chan string)
	go clientWrite(conn,ch)
	who :=conn.RemoteAddr().String()
	ch <- "You are"+who
	messages <- who+"has arrived"
	entering<-ch
	input :=bufio.NewScanner(conn)
	for input.Scan(){
		messages<-who+":"+input.Text()
	}

	leaving<-ch
	messages<-who +"has left"
	conn.Close()
}

func clientWrite(conn net.Conn,ch <-chan string){
	for msg := range ch{
		fmt.Fprintf(conn,msg)
	}
}
