package main

import (
	"fmt"
	"mapreduce/internal/reduce"
	"os"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	res, err := reduce.Map(6, numbers, func(value int) int {
		return value * 2
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}
