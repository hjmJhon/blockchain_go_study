package main

import (
	"fmt"
	"study.com/Day18/block"
)

func main() {

	firstBlock := block.GenerateFirstBlock()
	fmt.Println(firstBlock.String())

	block.Run()

}
