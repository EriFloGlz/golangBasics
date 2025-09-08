package main

import (
	"fmt"
	"os"
)

func main() {
	f, e := os.Create("hello.txt")

	if e != nil {
		fmt.Print(e)
		return
	}
	defer f.Close()
	_, err := f.Write([]byte("Hello"))

	if err != nil {
		fmt.Print(err)
		return
	}
}
