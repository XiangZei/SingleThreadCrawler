package engine

import (
	"SingleThreadCrawler/单任务爬虫/fetcher"
	"log"
)

func Run(seeds ...Request) {
	//Request 任务队列
	var requests []Request
	//将 seeds Request 放入 []requests ,即初始化[]requests
	for _, r := range seeds {
		requests = append(requests, r)
	}
	//执行任务
	for len(requests) > 0 {
		//1.获取第一个request 并从[]requests 移除，实现一个队列的功能
		r := requests[0]
		requests = requests[1:]
		//2.使用 爬取器 进行对Request.Url 进行爬取
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetch error ,url :%s,  err: %v", r.Url, err)
			continue
		}
		//3.使用 Request的解析函数对怕渠道的内容进行解析
		parseResult := r.ParseFunc(body)
		//4.将解析出来的[]Request 添加到请求任务队列requests的尾部
		requests = append(requests, parseResult.Requests...)

		//5.遍历解析出来的实体，直接打印
		for _, item := range parseResult.Items {
			log.Printf("getItems ,url: %s ,item: %v", r.Url, item)
		}
	}

	//

}
