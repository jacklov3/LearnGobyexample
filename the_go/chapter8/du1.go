package chapter8

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string,fileSize chan<-int64){
	for _,entry:=range dirents(dir){
		if entry.IsDir(){
			subdir :=filepath.Join(dir,entry.Name())
			walkDir(subdir,fileSize)
		}else{
			fileSize<-entry.Size()
		}
	}
}


func dirents(dir string)[]os.FileInfo{
	entries,err :=ioutil.ReadDir(dir)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"du1:%v\n",err)
		return  nil
	}
	return entries
}

func Test_du(){
	flag.Parse()
	roots :=flag.Args()
	if len(roots)==0{
		roots =[]string{"."}
	}
	fileSizes :=make(chan int64)
	go func() {
		for _,root:=range roots{
			walkDir(root,fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles,nbytes int64
	for size :=range fileSizes{
		nfiles++
		nbytes +=size
	}
	printDiskUsage(nfiles,nbytes)
}

func printDiskUsage(nfiles,nbytes int64){
	fmt.Printf("%d files %.1f GB\n",nfiles,float64(nbytes)/1e9)

}

var verbose = flag.Bool("v",false,"show verbose progress messages")

func Test_du2(){
	flag.Parse()
	roots :=flag.Args()
	if len(roots)==0{
		roots =[]string{"."}
	}
	fileSizes :=make(chan int64)
	go func() {
		for _,root:=range roots{
			walkDir(root,fileSizes)
		}
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose{
		tick = time.Tick(500*time.Millisecond)
	}
	var nfiles,nbytes int64
	loop:
		for{
			select {
			case size,ok :=<-fileSizes:
				if !ok{
					break loop;
				}
				nfiles++
				nbytes +=size
				case <-tick:
					printDiskUsage(nfiles,nbytes)
			}

		}
	printDiskUsage(nfiles,nbytes)
}

func Test_du3(){
	flag.Parse()
	roots :=flag.Args()
	if len(roots)==0{
		roots =[]string{"."}
	}
	fileSizes :=make(chan int64)
	var n sync.WaitGroup

	for _,root :=range roots{
		n.Add(1)
		go walkDir2(root,&n,fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose{
		tick = time.Tick(500*time.Millisecond)
	}
	var nfiles,nbytes int64
loop:
	for{
		select {
		case size,ok :=<-fileSizes:
			if !ok{
				break loop;
			}
			nfiles++
			nbytes +=size
		case <-tick:
			printDiskUsage(nfiles,nbytes)
		}

	}
	printDiskUsage(nfiles,nbytes)

}
func walkDir2(dir string,n *sync.WaitGroup,fileSizes chan<- int64){
	defer n.Done()
	for _,entry :=range dirents2(dir){
		if entry.IsDir(){
			n.Add(1)
			subdir :=filepath.Join(dir,entry.Name())
			go walkDir2(subdir,n,fileSizes)
		}else{
			fileSizes<-entry.Size()
		}
	}
}

//定义计数信号量；来阻止该函数打开太多文件
var sema = make(chan struct{},20)

func dirents2(dir string)[]os.FileInfo{
	sema <- struct{}{}
	defer func() {<-sema}()
	entries,err :=ioutil.ReadDir(dir)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"du1:%v\n",err)
		return  nil
	}
	return entries
}