package scheduler

import "SingleThreadCrawler/单任务爬虫1/model"

type QueueScheduler struct {
	requestChann chan model.Request

	workerChan chan chan model.Request
}

func (s *QueueScheduler) WorkerChann() chan model.Request {
	return make(chan model.Request)
}

func (s *QueueScheduler) Submit(request model.Request) {
	s.requestChann <- request
}
func (s *QueueScheduler) WorkerReady(w chan model.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) Run() {
	s.requestChann = make(chan model.Request)
	s.workerChan = make(chan chan model.Request)

	go func() {
		var requestQ []model.Request
		var workerQ []chan model.Request
		for {
			var activeRequest model.Request
			var activeWorker chan model.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChann:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
