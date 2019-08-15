package functions

import "fmt"

//多返回值
func vals()(int,int){
	return 3,7
}

func Test_mul_return(){
	a,b:=vals()
	fmt.Println(a)
	fmt.Println(b)

	_,c := vals()
	fmt.Println(c)
}