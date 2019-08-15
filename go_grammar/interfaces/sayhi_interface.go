package interfaces

import (
	"fmt"
)

type Humen struct{
	name string
	age int
	phone string
}

type Student struct{
	Humen //匿名字段
	school string
	locan float32
}

type Employee struct{
	Humen //匿名字段
	company string
	money float32
}
//Humen 对象实现了Sayhello方法
func (h *Humen) Sayhello(){
	fmt.Printf("你好，我是%s,今年%d岁，我的电话号码是%s",h.name,h.age,h.phone)
}

//Humen 对象实现了Sing方法
func (h *Humen)Sing(geci string){
	fmt.Printf("不如唱歌，我正在唱: %s\n",geci)
}

//Humen 对象实现了Dump方法
func (h *Humen)Dump(length int){
	fmt.Println("我可以跳%d米\n",length)
}

//Employee 重载Humen的Sayhello方法

func (e *Employee) Sayhello(){
	fmt.Println("我叫%s，我今年%d岁了，我的电话号码是%s，我目前工作在%s公司，月薪%f\n",e.name,e.age,e.phone,e.company,e.money)
}

//Student 实现了BorrowMoney方法
func (s *Student)BorrowMoney(amount float32){
	s.locan +=amount
}
//Employee实现SpendSalary方法
func(e *Employee) SpendSalary(amount float32){
	e.money -=amount
}

//此接口被Humen，Student和Employee实现
type Men interface{
	Sayhello()
	Sing(geci string)
	Dump(length int)
}

//这个接口被Student实现
type YoungChap interface{
	Sayhello()
	Sing(geci string)
	BorrowMoney(amount float32)
}

//这个接口被Employee实现
type ElderlyGent interface{
	Sayhello()
	Sing(geci string)
	SpendSalary(amount float32)
}

func Test_multi_interface(){
	humen1 :=Humen{"张三",25,"110"}
	student1 :=Student{Humen{"李四",18,"119"},"清华大学",100}
	employee1 :=Employee{Humen{"王五",23,"911"},"腾讯",20000}
	//定义Men接口的变量
	var i Men
	i = &humen1
	i.Sayhello()
	i.Sing("找个天使替我去爱你")
	i.Dump(100)

	i = &student1
	i.Sayhello()
	i.Sing("唱不完一首歌.....")
	i.Dump(200)

	i = &employee1
	i.Sayhello()
	i.Sing("冬天的秘密...")
	i.Dump(300)

	
}