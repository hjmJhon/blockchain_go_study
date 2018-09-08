package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("connet fail...")
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for true {
		str, readErr := reader.ReadString('\n')
		if readErr != nil {
			fmt.Println("read fail...")
			return
		}

		input := strings.Trim(str, "\r\n")
		if input == "Q" {
			fmt.Println("马上退出...")
			return
		}
		_, e := conn.Write([]byte(input))
		if e != nil{
			continue
		}
	}
}
