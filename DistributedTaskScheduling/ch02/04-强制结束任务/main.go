package main

import (
	"os/exec"
	"context"
	"time"
	"fmt"
)

type result struct {
	output []byte
	err    error
}

func main() {
	// 需求：在协程里执行1个cmd，执行两秒钟，sleep2;echo hello;，当执行到1秒钟的时候，把cmd杀死

	var (
		ctx context.Context // 接口，有很多方法的实现
		cancelFunc context.CancelFunc // 是一个取消函数
		cmd *exec.Cmd
		resultChan chan *result
		res *result
	)

	resultChan = make(chan *result, 1000)
	// ctx对象里面，有一个chan byte
	// cancelFunc的作用就是close(chan byte)，就是关闭ctx的chan
	// 其实这是一对，cancelFunc用于关闭chan，ctx用于感知chan被关闭
	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func () {
		var (
			output []byte
			err    error
		)
		// 第一个参数是一个上下文
		// 通过CommandContext构造的cmd对象，内部会做一个事情
		// select {case <- ctx.Done(): }，也就是监听chan是否被关闭
		// 一旦调用cancelFunc，chan就会被关闭，cmd就会监听到chan被关闭了
		// 就会通过kill操作系统调用，把bash程序杀死，杀死子进程
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;echo hello;")
		// 执行任务，捕获输出
		output, err = cmd.CombinedOutput()
		// 把任务输出结果，传给main协程
		resultChan <- &result{output: output, err: err}
	}()

	// 等待1秒钟
	time.Sleep(time.Second)

	// 取消上下文
	cancelFunc()

	// 等待子协程的退出，并打印任务执行结果
	res = <- resultChan
	fmt.Println(res.err, string(res.output)) // signal: killed 
}
