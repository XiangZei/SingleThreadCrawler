package main

import (
	"SingleThreadCrawler/单任务爬虫/engine"
	"SingleThreadCrawler/单任务爬虫/zhenai/parser"
)

type person struct {
	name string
	age  int
}

func main() {

	engine.Run(engine.Request{
		// 种子 Url
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
