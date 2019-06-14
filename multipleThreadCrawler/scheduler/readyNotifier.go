package scheduler

import "SingleThreadCrawler/multipleThreadCrawler/model"

type ReadyNotifier interface {
	WorkerReady(chan model.Request)
}
