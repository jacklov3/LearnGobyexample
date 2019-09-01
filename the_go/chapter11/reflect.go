package chapter11

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
}

func Reflect(){
	u:=User{"张三",20}
	t:=reflect.TypeOf(u)
	v:=reflect.ValueOf(u)
	fmt.Println(t)
	fmt.Println(v)
	fmt.Printf("%T\n",u)
	fmt.Printf("%v\n",u)

	//从reflect value还原到type
	u1:=v.Interface().(User)
	fmt.Println(u1)

	//方法二
	t1:=v.Type()
	fmt.Println(t1)
	fmt.Println(t.Kind())
}