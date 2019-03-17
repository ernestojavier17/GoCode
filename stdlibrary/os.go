package main

import (
	"fmt"
	"os"
)

func main() {
	var path string = os.Getenv("PATH")
	fmt.Println(path)
	fmt.Println(path.(type))
}