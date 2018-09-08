package parser

import (
	"study.com/Day15/zhenai/engine"
	"regexp"
	"fmt"
)

const cityRegex string = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a></th>`

/*
	城市解析器
 */
func CityParser(str string) (r engine.RequestResult) {
	compile, err := regexp.Compile(cityRegex)
	if err != nil {
		fmt.Println("err:", err)
	}
	data := compile.FindAllStringSubmatch(str, -1)
	for _, v := range data {
		fmt.Println("用户名",v[2],",Url:",v[1])
		r.Items = append(r.Items, v[2])
		r.R = append(r.R, engine.Request{
			Url:   v[1],
			Parse: UserInfoParser,
		})
	}
	return r
}
