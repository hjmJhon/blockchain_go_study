package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

/*
	网络请求获取数据
 */
func FetchData(url string) ([]byte, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", "response code is", response.StatusCode)
	}

	encode := determinEncoding(response.Body)

	reader := transform.NewReader(response.Body, encode.NewDecoder())

	return ioutil.ReadAll(reader)
}


func determinEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
