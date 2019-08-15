package sorting

import (
	"sort"
	"fmt"
)
type byLength []string

func (s byLength) Len() int{
	return len(s)
}
func (s byLength)Swap(i,j int){
	s[i],s[j]=s[j],s[i]
}

func (s byLength) Less(i,j int)bool{
	return len(s[i])<len(s[j])
}

func Test_sort_by_func(){
	fruits :=[]string{"peach","banana","kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

