package engine

// url string
// 解析器 方法

//func (b []byte) 类型
//CityListParser 变量名

// Request
// Url
// ParserFunc
// Request{} 创建一个对象，对象需要空间存储
// 存储空间是客观存在的
// 经纬度{123.9 45.8}
//r := Request{}
//r = Request{}
//
//pr := &Request{}
//


//RequestResult{
//	["惠儿","风中的蒲公英","幽诺"],
//	[Request{
//		"http://album.zhenai.com/u/108415017",
//		NilParser,
//	},Request{
//		"http://album.zhenai.com/u/1314495053",
//		NilParser,
//	},Request{
//		"http://album.zhenai.com/u/110171680",
//		NilParser,
//	}],
//}

type Request struct {
	Url string  // 每一个请求对应的Url
	ParserFunc func (b []byte) (r RequestResult) // 每一个页面取到的数据都传给解析器进行解析
}

// 每一个请求返回回来的结构体对象
type RequestResult struct {
	Items []interface{} //存储解析器解析出来的数据
	R []Request //数据请求数组
}

// 空的解析器
func NilParser(b []byte) (r RequestResult) {
	return r
}