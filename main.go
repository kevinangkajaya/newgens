package main

import (
	"fmt"
	"newgens/src"
)

func main() {
	path := src.ReadConsole()
	data, err := src.ReadLines(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
