package channels

import "fmt"

func Test_no_block(){
	message :=make(chan string,1)
	signals :=make(chan bool)

	select{
	case msg:=<-message:
		fmt.Println("received message",msg)
	default:
		fmt.Println("no message received")
	}

	msg :="hi"
	select{
	case message <-msg:
		fmt.Println("sent message",msg)
	default:
		fmt.Println("no message sent")
	}
	select{
	case msg:=<-message:
		fmt.Println("received message 3",msg)
	case sig:=<-signals:
		fmt.Println("received signal",sig)
	default:
		fmt.Println("no activity")
	}
}
