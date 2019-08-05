package slices

import "fmt"

func Test_slice(){
	s :=make([]string,3)
	fmt.Println("emp:",s)
	fmt.Println(&s[0])

	s[0]="a"
	s[1]="b"
	s[2]="c"
	fmt.Println("set:",s)
	fmt.Println("get:",s[2])
	fmt.Println("len:",len(s))
	fmt.Println(&s[0])
	s=append(s,"d")
	fmt.Println(&s[0])
	s=append(s,"e","f")
	fmt.Println("apd:",s)
	fmt.Println(&s[0])

	c:=make([]string,len(s))

	copy(c,s)
	fmt.Println("copy:",c)
	l:=s[2:5]
	fmt.Println("sl1:",l)
	fmt.Println(&s[2],&l[0])
	l=s[2:]
	fmt.Println("sl3:",l)

	l=s[:5]
	fmt.Println("sl3:",l)

	//声明并初始化一个切片
	t:=[]string{"g","h","i"}
	fmt.Println("dcl:",t)

	twoD :=make([][]int,3)
	for i:=0;i<3;i++{
		innerlen :=i+1
		twoD[i]=make([]int,innerlen)
		for j:=0;j<innerlen;j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d:",twoD)

	x :=[5]int{1,2,3,4,5}
	sb :=x[:]
	fmt.Println(x)
	fmt.Println(sb)
	ne := make([]int,len(sb))
	copy(ne,sb)
	ne[0]=100
	fmt.Println(ne)
	fmt.Println(sb)
	fmt.Println(x)
}