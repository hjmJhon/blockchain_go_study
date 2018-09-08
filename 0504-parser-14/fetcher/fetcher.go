package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/transform"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

func FetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error:Status Code, ",resp.StatusCode)
		return nil, fmt.Errorf("Error:Status Code:%d", resp.StatusCode)
	}

	// 动态判断编码
	e := determinEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body,
		e.NewDecoder())
	// ([]byte, error)
	return ioutil.ReadAll(utf8Reader)

}

// 将response.Body 作为参数传入到函数中
// 会自动返回编码格式
func determinEncoding(r io.Reader) encoding.Encoding {

	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		//panic(err)
		return unicode.UTF8
	}
	// 查找编码格式
	e, _, _ := charset.DetermineEncoding(bytes, "")
	// 返回编码格式
	return e
}
