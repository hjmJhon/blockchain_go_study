package main

// 动态判断编码
// go get golang.org/x/net/html

import (

	"study.com/0504-parser-14/engine"
	"study.com/0504-parser-14/zhenai/parser"
)

func main() {
	//无限深入的去爬虫
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parser.CityListParser,
	})

}


