package file_paths

import(
	"fmt"
	"path/filepath"
	"strings"
)

func Test_file_path(){
	p :=filepath.Join("dir1","dir2","filename")
	fmt.Println("p:",p)

	fmt.Println(filepath.Join("dir1//","filrname"))
	fmt.Println(filepath.Join("dir1/../dir1","filename"))

	fmt.Println("Dir(p):",filepath.Dir(p))
	fmt.Println("Base(p):",filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	filename :="config.json"
	ext :=filepath.Ext(filename)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(filename,ext))
	rel,err :=filepath.Rel("a/b","a/b/t/file")
	if err !=nil{
		panic(err)
	}
	fmt.Println(rel)
	rel,err =filepath.Rel("a/b","a/c/t/file")
	if err!=nil{
		panic(err)
	}
	fmt.Println(rel)
}