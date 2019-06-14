package scheduler

import "SingleThreadCrawler/multipleThreadCrawler/model"

//调度器接口
type Scheduler interface {
	ReadyNotifier
	//提交 request 到调度器的request任务通道中
	Submit(request model.Request)
	//初始化当前调度器实例的request任务通道
	WorkerChann() chan model.Request
	Run()
}
