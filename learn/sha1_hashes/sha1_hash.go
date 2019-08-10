package sha1_hashes

import (
	"crypto/sha1"
	"fmt"
)

func Test_sha1_hash(){
	s :="sha1 this string "
	h :=sha1.New()

	//转换成字节序列
	h.Write([]byte(s))
	bs :=h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n",bs)

}