package main

import (
	"MyNet/core/frame/tensor"
	"fmt"
)

func test(args ...int) {
	if len(args) == 0 {
		println(0)
	} else {
		fmt.Println(args)
	}
}

func main() {
	t := tensor.NewTensor[int](3, 5)
	println(t.Offset(2, 1))
}
