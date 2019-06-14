package engine

import (
	"SingleThreadCrawler/单任务爬虫/fetcher"
	"SingleThreadCrawler/单任务爬虫/model"
	"SingleThreadCrawler/单任务爬虫/scheduler"
	"log"
)

type ConcurrentEngine struct {
	//调度器
	Scheduler scheduler.Scheduler
	//开启worker数量
	WorkerCount int
}

func (e ConcurrentEngine) Run(seeds ...model.Request) {
	in := make(chan model.Request)
	out := make(chan model.ParseResult)
	//初始化调度器
	e.Scheduler.ConfigureMasterWorkerChan(in)
	//创建workerCount 个worker
	for i := 0; i < e.WorkerCount; i++ {
		go func(in chan model.Request, out chan model.ParseResult) {
			for {
				r := <-in
				body, err := fetcher.Fetch(r.Url)
				var result model.ParseResult
				if err != nil {
					log.Printf("fetch error, url: %s, err: %v", r.Url, err)
					result = model.ParseResult{}
				}
				result = r.ParseFunc(body)
				out <- result
			}
		}(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out //阻塞获取
		for _, item := range result.Items {
			log.Println("getItems ,items :%v", item)
		}
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}

}
