package functions

import "fmt"

func intSeq() func() int{
	i :=0
	return func() int {
		i++
		return i
	}
}

func Test_closures(){
	nextInt :=intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts :=intSeq()
	fmt.Println(newInts())
}