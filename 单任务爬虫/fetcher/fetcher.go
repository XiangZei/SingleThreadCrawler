package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	//1.爬取URL
	if url[len(url)-1] <= '9' && url[len(url)-1] >= '0' {
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
		req.Header.Add("Referer", url)
		resp, err := client.Do(req)
		if err != nil {

			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Wrong statusCode , %d", resp.StatusCode)
		}
		//读取响应体 并返回
		return ioutil.ReadAll(resp.Body)

	} else {
		resp, err := http.Get(url)
		if err != nil {

			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Wrong statusCode , %d", resp.StatusCode)
		}
		//读取响应体 并返回
		return ioutil.ReadAll(resp.Body)
	}

}
