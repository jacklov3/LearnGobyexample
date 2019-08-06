package panic

import "os"

func Test_panic(){
	panic("a problem")

	_,err :=os.Create("/tmp/file")
	if err!=nil{
		panic(err)
	}
}