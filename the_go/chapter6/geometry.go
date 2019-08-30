package chapter6

import (
	"fmt"
	"math"
)

type Point struct {
	X,Y float64
}

//普通函数
func Distance(p,q Point) float64{
	return math.Hypot(q.X-p.X,q.Y-p.Y)
}

//方法
func (p Point)Distance(q Point)float64{
	return math.Hypot(q.X-p.X,q.Y-p.Y)
}

func Test_Distance(){
	p:=Point{1,2}
	q:=Point{4,6}
	fmt.Println(Distance(p,q))
	fmt.Println(p.Distance(q))
}


