package chapter9


//var balance int
//
//func Deposit(amount int){balance =balance+amount}
//
//func Balance() int {return balance}
//
//
//func Test_bank(){
//	go func() {
//		Deposit(200)
//		fmt.Println("=",Balance())
//	}()
//
//	go Deposit(100)
//}


//不要使用共享变量来通信，而要使用通信来共享变量

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int){deposits <-amount}
func Balance() int {return <-balances}

//监听goroutine，共享变量被限制在单一的groutine里，只能通过channel来共享变量
func teller(){
	var balance int//仅仅限制在该单一groutine里
	for{
		select {
		case amount :=<-deposits:
			balance +=amount
			case balances <-balance:
		}
	}
}

func init(){
	go teller()
}