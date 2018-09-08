package engine

import (
	"study.com/Day15/zhenai/fetcher"
	"fmt"
)

func Run(r ...Request) {
	for len(r) > 0 {
		firstRequest := r[0]
		fmt.Println("正在爬取的网页是:", firstRequest.Url)
		r = r[1:]
		bytes, err := fetcher.FetchData(firstRequest.Url)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}

		result := firstRequest.Parse(string(bytes))
		r = append(r, result.R...)
	}

}
