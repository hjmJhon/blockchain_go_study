package parser

import (
	"regexp"
	"study.com/0504-parser-14/engine"
)

const regexpString = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>].+>([^<].+)</a>`

// 城市列表解析器
func CityListParser(b []byte) (r engine.RequestResult){
	// ^[a-z]
	// [^[a-z]]
	// [a-z0-9]+
	//
	// class=""
	// >
	//[^<].+ 汉字城市名字
	match := regexp.MustCompile(regexpString)

	// 查找所有的匹配的字符串
	bytes := match.FindAllSubmatch(b, -1)

	counter := 2
	for _, m := range bytes {
		//fmt.Printf("City:%s  URL:%s\n", m[2], m[1])
		//fmt.Println(r)
		r.Items = append(r.Items,m[2])
		r.R = append(r.R,engine.Request{
			string(m[1]),
			CityParser,
		})
		counter--
		if counter == 0 {
			break
		}
	}
	return
}
