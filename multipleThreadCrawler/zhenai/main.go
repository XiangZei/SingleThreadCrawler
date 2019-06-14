package main

import (
	"SingleThreadCrawler/multipleThreadCrawler/engine"
	"SingleThreadCrawler/multipleThreadCrawler/model"
	"SingleThreadCrawler/multipleThreadCrawler/scheduler"

	"SingleThreadCrawler/multipleThreadCrawler/zhenai/parser"
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
