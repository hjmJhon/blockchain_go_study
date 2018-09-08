package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
)

//将数据写到文件
func writeFile() {
	file, err := os.OpenFile("/Users/hjm/Desktop/test/1.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer file.Close()

	//fmt.Fprintf(file, "hello,world 123 \n")

	fileWriter := bufio.NewWriter(file)
	_, e := fileWriter.WriteString("hhahahaha,huangjianging \n")
	if e == nil {
		fmt.Println("写文件成功")
		fileWriter.Flush()
	}
}

//用缓冲流读文件
func readFile() {
	file, err := os.Open("/Users/hjm/Desktop/test/1.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("err,", err)
		return
	}

	reader := bufio.NewReader(file)
	str, readerErr := reader.ReadString('\n')
	if readerErr != nil {
		fmt.Println("读取文件失败")
		return
	}

	fmt.Println("str:", str)

}

type CharStruct struct {
	chCount    int
	numCount   int
	spaceCount int
	otherCount int
}

//读文件,并统计其中字符,数字,空格的个数
func readFileAndCacul() {
	file, err := os.Open("/Users/hjm/Desktop/test/1.txt")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()

	var charStruct CharStruct

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读文件失败")
			return
		}
		bytes := []rune(str)
		for _, v := range bytes {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				charStruct.chCount++

			case v >= '0' && v <= '9':
				charStruct.numCount++
			case v == ' ' || v == '\t':
				charStruct.spaceCount++
			default:
				charStruct.otherCount++
			}
		}
	}
	fmt.Println("charStruct=", charStruct)

}

//用缓冲流读标准输入中的数据
func readBuffer() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read failed")
		return
	}
	fmt.Println("str=", str)
}

//文件的拷贝
func copyFile() {
	inputFile := "/Users/hjm/Desktop/test/1.txt"
	outputFile := "/Users/hjm/Desktop/test/2.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}

	writeFileErr := ioutil.WriteFile(outputFile, bytes, os.ModePerm)
	if writeFileErr != nil{
		fmt.Println("写文件失败")
	}

}

func main() {
	//writeFile()
	//
	//readFile()

	//readFileAndCacul()

	//copyFile()

	//readBuffer()

}
