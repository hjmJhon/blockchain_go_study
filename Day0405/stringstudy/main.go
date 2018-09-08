package main

import (
	"fmt"
	"strings"
)

func appendStr(str string,str2 string) string {
	return str+ str2
}



func main() {
	var str = "hello world"
	if strings.HasPrefix(str, "hell") {
		fmt.Printf("str is not begin with %s",str)
	}

	fmt.Println()
	var append_method  = appendStr
	fmt.Printf("append_method result = %s",append_method("hello","world"))
}
