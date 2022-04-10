package test

import (
	"errors"
	"fmt"
	"go-tg-bot/helpers"
	"go-tg-bot/internal/utils/cache"
	"sync"
	"testing"
)

func emulateLoad(t *testing.T, c cache.Cacher, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("#{%d}-key", i)
		value := fmt.Sprintf("#{%d}-value", i)

		wg.Add(1)

		go func(k string) {
			err := c.Set(k, value)
			if err != nil {
				t.Error("Set error: ", err)
				t.Fail()
			}
			wg.Done()
		}(key)

		wg.Add(1)

		go func(k, v string) {
			storedValue, err := c.Get(k)

			if !errors.Is(err, helpers.ErrNotFound) {
				if storedValue != v {
					t.Error("storedValue != v", err)
					t.Fail()
				}
			}

			wg.Done()
		}(key, value)

		wg.Add(1)

		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				t.Error("Delete error: ", err)
				t.Fail()
			}
			wg.Done()
		}(key)
	}

	wg.Wait()
}

func Test_Cache(t *testing.T) {
	t.Parallel()

	testCache := cache.NewCache()

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()

		parallelFactor := 100_000
		emulateLoad(t, testCache, parallelFactor)
	})
}
