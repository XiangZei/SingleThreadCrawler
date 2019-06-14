package scheduler

import "SingleThreadCrawler/单任务爬虫/model"

//并发调度器

type SimpleSchduler struct {
	workChan chan model.Request
}

func (s *SimpleSchduler) ConfigureMasterWorkerChan(in chan model.Request) {
	s.workChan = in
}

func (s *SimpleSchduler) Submit(request model.Request) {
	go func() {
		s.workChan <- request
	}()
}
