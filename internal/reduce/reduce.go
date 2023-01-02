package reduce

import (
	"errors"
	"fmt"
	"sync"
)

var (
	dataNil         = errors.New("data is nil")
	goroutinesIsNil = errors.New("goroutines is nil")
)

func Map[T any, R any](treads int, data []T, f func(T) R) ([]R, error) {
	if treads == 0 {
		return nil, fmt.Errorf("Map: %w", goroutinesIsNil)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("Map: %w", dataNil)
	}

	wg := sync.WaitGroup{}
	wg.Add(treads)

	res := make([]R, len(data))
	for i := 0; i < treads; i++ {
		go func(start, step int, datSize int) {
			defer wg.Done()

			for index := start; index < datSize; index += step {
				res[index] = f(data[index])
			}
		}(i, treads, len(data))
	}

	wg.Wait()
	return res, nil
}
