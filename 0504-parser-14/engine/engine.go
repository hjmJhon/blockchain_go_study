package engine

import (
	"fmt"
	"study.com/0504-parser-14/fetcher"
	"log"
)

// 通过引擎去调用,参数

// 种子
//Request{"http://www.zhenai.com/zhenghun",CityListParser}

// HTML
// iOS Android Java C++

func Run(r ...Request) {

	for len(r) > 0 {
		firstRequest := r[0]
		r = r[1:]
		fmt.Println("即将请求：",firstRequest.Url)
		all, err := fetcher.FetchData(firstRequest.Url)

		if err != nil {
			log.Printf("%v\n", err)
		}
		//fmt.Printf("%s\n", all)

		result := firstRequest.ParserFunc(all)
		//fmt.Printf("%s", result)
		r = append(r, result.R...)
	}

}
