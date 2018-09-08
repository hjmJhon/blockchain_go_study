package main

import (
	"study.com/Day15/zhenai/engine"
	"study.com/Day15/zhenai/parser"
)

const cityListUrl = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		Url:   cityListUrl,
		Parse: parser.CityListParse,
	}, )
}
