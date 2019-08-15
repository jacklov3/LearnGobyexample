package interfaces

import (
	"fmt"
	"reflect"
)

func Test_reflect(){
	var x float64=3.4
	t := reflect.TypeOf(x)
	fmt.Println(t)
	v :=reflect.ValueOf(x)
	fmt.Println(v.Type())
	fmt.Println(v.Float())
	fmt.Println(v.Kind())

}