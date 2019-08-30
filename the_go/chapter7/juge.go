package chapter7

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Test_juge(){
	var w io.Writer
	w=os.Stdout
	f:=w.(*os.File)
	c,_:=w.(*bytes.Buffer)
	fmt.Printf("%T %T",f,c)
	rw:=w.(io.ReadWriter)
	fmt.Printf("%T",rw)
}