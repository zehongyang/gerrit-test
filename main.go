package main

import "fmt"

type MyInt interface {
	int | int8 | int16 | int32 | int64
}

func GetMaxNum[T MyInt](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(1)
	fmt.Println("test")
	fmt.Println("1111")
	GetMaxNum(1, 2)
}
