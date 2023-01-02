package reduce

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 1},
	{input: 5},
	{input: 10},
	{input: 100},
	{input: 1000},
	{input: 74382},
	{input: 382399},
	{input: 764798},
}

func BenchmarkReduceMap(b *testing.B) {
	count := rand.Intn(100000000)
	fmt.Printf("\nrand number: %v\n", count)
	numbers := make([]int, 0, count)
	for i := 1; i < count; i++ {
		numbers = append(numbers, i)
	}

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			Map(b.N, numbers, func(value int) string {
				return strconv.Itoa(value * 2)
			})
		})
	}
}
