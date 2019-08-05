package variables

import "fmt"

func PrintVariables() {
	var a = "initial"
	fmt.Println(a)
	var b, c = 1, 2
	fmt.Println(b, c)
	var e int
	fmt.Println(e)
	f := "apple"
	fmt.Println(f)
}

