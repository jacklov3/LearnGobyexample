package chapter4

import (
	"fmt"
)
//添加元素到切片中
func appendtoslice(x[]int,y int)[]int{
	var z[]int
	zlen :=len(x)+1
	if zlen<=cap(x){
		z=x[:zlen]
	}else{
		zcap:=zlen
		if zcap<2*len(x){
			zcap=2*len(x)
		}
		z=make([]int,zlen,zcap)
		copy(z,x)
	}
	z[len(x)]=y
	return z
}

//添加多元素到切片中
func appendmultitoslice(x []int,y ...int)[]int{
	var z[]int
	zlen:=len(x)+len(y)
	if zlen<=cap(x){
		copy(x[len(x):],y)
		z=x
	}else {
		zcap:=zlen
		if zcap<2*len(x){
			zcap=2*len(x)
		}
		z=make([]int,zlen,zcap)
		copy(z,x)
	}
	copy(z[len(x):],y)
	return z

}


//测试函数
func Test_appendtoslice(){
	//var x,y[]int
	//for i:=0;i<10;i++{
	//	y=appendtoslice(x,i)
	//	fmt.Printf("%d cap=%d\t%v\n",i,cap(y),y)
	//	x=y
	//}
	var x[]int
	x=appendmultitoslice(x,1)
	x=appendmultitoslice(x,2,3)
	x=appendmultitoslice(x,4,5,6)
	x=appendmultitoslice(x,x...)
	fmt.Println(x)
}


