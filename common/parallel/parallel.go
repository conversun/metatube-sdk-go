package parallel

import (
	"sync"

	"github.com/metatube-community/metatube-sdk-go/common/recovery"
)

func Parallel[T any, R any](fn func(T) R, args ...T) []R {
	var wg sync.WaitGroup
	results := make([]R, len(args))

	for i, v := range args {
		wg.Add(1)
		go func(i int, v T) {
			defer wg.Done()
			defer recovery.Recover("Parallel")
			results[i] = fn(v)
		}(i, v)
	}

	wg.Wait()
	return results
}
