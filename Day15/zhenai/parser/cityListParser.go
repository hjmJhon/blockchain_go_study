package parser

import (
	"regexp"
	"fmt"
	"study.com/Day15/zhenai/engine"
	)

const regexpRule = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>].+>([^<].+)</a>`

/*
	城市列表解析器
 */
func CityListParse(str string) (r engine.RequestResult) {
	compile, e := regexp.Compile(regexpRule)
	if e != nil {
		fmt.Println(" fail,error=", e)
		return r
	}

	allStr := compile.FindAllStringSubmatch(str, -1)

	count := 2
	for _, v := range allStr {
		//fmt.Println(v[2], v[1]) //城市名 城市链接
		r.Items = append(r.Items, v[2])
		r.R = append(r.R, engine.Request{
			Url:   v[1],
			Parse: CityParser,
		})

		count --
		if count < 0 {
			break
		}
	}

	return r
}
