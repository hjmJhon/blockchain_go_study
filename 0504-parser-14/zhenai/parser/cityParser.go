package parser

import (
	"regexp"
	"study.com/0504-parser-14/engine"
	"fmt"
)


// <th><a href="http://album.zhenai.com/u/1314495053" target="_blank">风中的蒲公英</a></th>
const regexpCity = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a></th>`

// 城市列表解析器
func CityParser(b []byte) (r engine.RequestResult){

	match := regexp.MustCompile(regexpCity)

	// 查找所有的匹配的字符串
	bytes := match.FindAllSubmatch(b, -1)


	for _, m := range bytes {
		fmt.Printf("用户的名字:%s  URL:%s\n", m[2], m[1])
		//fmt.Println(r)
		r.Items = append(r.Items,string(m[2]))
		r.R = append(r.R,engine.Request{
			string(m[1]),
			UserInfoParser,
		})
	}
	return
}
