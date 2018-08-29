package main

import (
	"fmt"
	memo "gopl/ch9/memo1"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	memo := memo.New(httpGetBody)
	var n sync.WaitGroup
	incommingURLs := []string{"https://www.douban.com",
		"https://www.baidu.com",
		"https://www.bilibili.com",
		"https://www.bilibili.com",
		"https://www.bilibili.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.bilibili.com",
		"https://www.bilibili.com",
		"https://www.bilibili.com",
		"https://www.baidu.com"}

	for _, url := range incommingURLs {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := memo.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
