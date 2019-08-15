package main


import (
	"fmt"
)

type Humen struct{
	name string
	age int
	phone string
}

type Student struct{
	Humen 
	shcool string
}

type Employee struct{
	Humen 
	company string
}

func (h *Humen) SayHi(){
	fmt.Printf("Hi,I am %s you can call me on %s\n",h.name,h.phone)
}
func (e *Employee) SayHi(){
	fmt.Printf("Hi, I am %s,I work at %s. Call me on %s\n",e.name,e.company,e.phone)
}

func main(){
	mark := Student{Humen{"Mark",25,"222-222-YYYY"},"MIT"}
	sam :=Employee{Humen{"Sam",45,"111-888-XXXX"},"Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}