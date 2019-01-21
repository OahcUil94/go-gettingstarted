package main

import (
	"github.com/gorhill/cronexpr"
	"time"
	"fmt"
)

// 代表一个任务
type CronJob struct {
	expr *cronexpr.Expression
	nextTime time.Time
}

func main() {
	// 需要有一个调度协程, 定时检查它所有的cron任务, 谁过期了，就执行谁

	// 1. 定义两个cronjob
	var (
		expr *cronexpr.Expression
		now      time.Time
		cronJob *CronJob
		scheduleTable map[string]*CronJob
	)

	scheduleTable = make(map[string]*CronJob)
	now = time.Now()

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr: expr, nextTime: expr.Next(now),
	}

	// 注册任务到调度表中
	scheduleTable["job1"] = cronJob

	now = time.Now()
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr: expr, nextTime: expr.Next(now),
	}

	scheduleTable["job2"] = cronJob

	// 启动一个调度协程
	go func () {
		var (
			jobName string
			cronJob *CronJob
			now time.Time
		)

		// 定时检查一下任务调度表
		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				// 判断是否过期
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					// 启动一个协程执行任务
					go func (jobName string) {
						fmt.Println("执行任务：", jobName)
					}(jobName)

					// 计算下一次的调度时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "的下次调度时间是：", cronJob.nextTime)
				}
			}

			// 为了避免死循环不停的扫描, 防止耗费CPU, 所以让协程睡眠100毫秒, 让调度协程有一定的间隔, 避免把CPU打满
			// 原理就是Timer的C属性是一个chan，时间到期后，chan写入一个对象
			// 最直接的就是使用time.Sleep(100 * time.Millisecond)
			// Timer实际上是通过一个chan来通知我们到期了，sleep是睡眠
			select {
			case <- time.NewTimer(100 * time.Millisecond).C: // 将在100毫秒可读
			}
		}
	}()

	// 防止主协程退出
	time.Sleep(100 * time.Minute)
}
