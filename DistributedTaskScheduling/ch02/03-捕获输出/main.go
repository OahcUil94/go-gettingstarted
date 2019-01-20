package main

import (
	"os/exec"
	"fmt"
)

func main() {
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	// 创建cmd
	cmd = exec.Command("/bin/bash", "-c", "sleep 5;ls -l")
	// 这个方法其实是执行子进程并且捕获了子进程的输出（通过pipe）
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}
