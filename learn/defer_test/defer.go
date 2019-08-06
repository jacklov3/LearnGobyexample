package defer_test

import (
	"fmt"
	"os"
)

func Test_defer(){
	f :=createFile("/tmp/defer.txt")
	defer closFile(f)
	writeFile(f)

}

func createFile(p string) *os.File{
	fmt.Println("creating")
	f,err :=os.Create(p)
	if err !=nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("writring")
	fmt.Fprintln(f,"data")
}

func closFile(f *os.File){
	fmt.Println("closing")
	err :=f.Close()

	if err!=nil{
		fmt.Fprintf(os.Stderr,"error: %v\n",err)
		os.Exit(1)
	}
}