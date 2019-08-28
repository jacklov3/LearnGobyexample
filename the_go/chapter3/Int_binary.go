package chapter3

import (
	"fmt"
)

func Binary_int() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n",y)
	fmt.Printf("%08b\n",x&y)
	fmt.Printf("%08b\n",x|y)
	fmt.Printf("%08b\n",x&^y)
}