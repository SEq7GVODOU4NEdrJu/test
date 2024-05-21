package main

import (
	"fmt"
	"os"
)

func main() {
	test := os.Getenv("APOLLO")
	fmt.Printf("hello,world:%s", test)
}
