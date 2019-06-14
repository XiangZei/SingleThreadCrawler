package main

import (
	"SingleThreadCrawler/单任务爬虫1/engine"
	"SingleThreadCrawler/单任务爬虫1/model"
	"SingleThreadCrawler/单任务爬虫1/scheduler"

	"SingleThreadCrawler/单任务爬虫1/zhenai/parser"
)

type person struct {
	name string
	age  int
}

func main() {

	engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
	}.Run(model.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
