package main

import (
	"github.com/gorhill/cronexpr"
	"fmt"
	"time"
)

func main() {
	var (
		expr *cronexpr.Expression
		err  error
		now time.Time
		nextTime time.Time
	)

	// Parse和MustParse的区别就是前者会进行解析判断，返回错误对象，后者默认会认为表达式没有问题，如果有问题，直接panic
	// cronexpr这个库实现的要比linux复杂，它支持到了秒级和年份级别, 下面表达式表示每五秒钟执行一次
	// 年份的话是从1970年到2099年
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	// 当前时间
	now = time.Now()
	// 下次调度时间
	nextTime = expr.Next(now)

	// 类似于setTimeout, 多久之后调用函数
	time.AfterFunc(nextTime.Sub(now), func () {
		fmt.Println("执行了一次")
	})

	time.Sleep(5 * time.Second)
}