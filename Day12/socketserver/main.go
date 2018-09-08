package main

import (
	"net"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net fail...")
		return
	}

	fmt.Println("server is start...")

	for true {
		conn, acceptErr := listener.Accept()
		if acceptErr != nil{
			fmt.Println("Accept 失败...")
			continue
		}

		go processConn(conn)
	}


}
func processConn(conn net.Conn) {
	defer conn.Close()
	for {
		buffer := make([]byte,1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("read fail...")
			return
		}

		str := string(buffer[0:n])
		fmt.Println("str:",str)
	}
}
