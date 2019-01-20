package main

import (
	"os/exec"
	"fmt"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
	)

	// Command方法，方法传入的一个参数是要调用的程序是什么
	// 在windows系统下，没有/bin/bash命令，就需要安装cygwin，它的用途就是在windows系统上虚拟一套linux命令
	cmd = exec.Command("/bin/bash", "-c", "echo 1;echo 2;echo 3")
	err = cmd.Run() // 运行命令
	fmt.Println(err)
}
