package engine

import (
	"SingleThreadCrawler/multipleThreadCrawler/fetcher"
	"SingleThreadCrawler/multipleThreadCrawler/model"
	"SingleThreadCrawler/multipleThreadCrawler/scheduler"
	"log"
)

type ConcurrentEngine struct {
	//调度器
	Scheduler scheduler.Scheduler
	//开启worker数量
	WorkerCount int
}

func (e ConcurrentEngine) Run(seeds ...model.Request) {
	e.Scheduler.Run()
	out := make(chan model.ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		go func(in chan model.Request, out chan model.ParseResult, notifier scheduler.ReadyNotifier) {
			for {
				notifier.WorkerReady(in)
				r := <-in
				log.Printf("fetching url:%s", r.Url)
				var result model.ParseResult
				body, err := fetcher.Fetch(r.Url)
				if err != nil {
					log.Printf("fetch error ,url:%s ,err:%v", r.Url, err)
					result = model.ParseResult{}
				}
				result = r.ParseFunc(body)
				out <- result
			}
		}(e.Scheduler.WorkerChann(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out //阻塞获取
		for _, item := range result.Items {
			log.Printf("getItems ,items: %v", item)

		}
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}

}
