package spawning_processes

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Test_spawning_processes(){
	dateCmd :=exec.Command("ls")
	dateOut,err :=dateCmd.Output()
	if err!=nil{
		panic(err)
	}
	fmt.Println(">date")
	fmt.Println(string(dateOut))

	grepCmd :=exec.Command("grep","hello")
	grepIn,_:=grepCmd.StdinPipe()
	grepOut,_:=grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes,_ :=ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd :=exec.Command("bash","-c","ls -a -l -h")
	lsOut,err :=lsCmd.Output()
	if err !=nil{
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}