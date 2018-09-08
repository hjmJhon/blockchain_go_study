package main

import "fmt"

func reverse(str string)  string{
	result := ""
	for i := 0; i < len(str); i++  {
		result+= fmt.Sprintf("%c",str[len(str)-i-1])

	}
	return result
}

func reverse2(str string) string {
	var result  []byte
	temp := []byte(str)
	strLen := len(str)
	for i:= 0; i < strLen; i++ {
		result = append(result,temp[strLen-i-1])
	}
	return string(result)

}

func main() {
	s := reverse("hello")
	fmt.Println("s = ",s)

	fmt.Println("------------------")

	result := reverse2("hello world !")
	fmt.Println("result = ",result)
}
