package scheduler

import "SingleThreadCrawler/单任务爬虫1/model"

type ReadyNotifier interface {
	WorkerReady(chan model.Request)
}
