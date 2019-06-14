package main

import (
	"SingleThreadCrawler/单任务爬虫/engine"
	"SingleThreadCrawler/单任务爬虫/model"
	"SingleThreadCrawler/单任务爬虫/scheduler"
	"SingleThreadCrawler/单任务爬虫/zhenai/parser"
)

type person struct {
	name string
	age  int
}

func main() {
	//并发版的多线程
	engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleSchduler{},
		WorkerCount: 1000,
	}.Run(model.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
	//队列调度器

}
