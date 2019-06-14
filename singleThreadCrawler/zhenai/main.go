package main

import (
	"SingleThreadCrawler/singleThreadCrawler/engine"
	"SingleThreadCrawler/singleThreadCrawler/model"
	"SingleThreadCrawler/singleThreadCrawler/scheduler"
	"SingleThreadCrawler/singleThreadCrawler/zhenai/parser"
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
