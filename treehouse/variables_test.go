package main

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	slice := []int{1, 2, 3, 5, 8}

	v, _ := First(slice)
	fmt.Println(v)
}
