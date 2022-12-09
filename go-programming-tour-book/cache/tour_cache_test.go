package cache_test

import (
	"log"
	"sync"
	"testing"

	"github.com/lvdbing/go-programming-tour-book/cache"
	"github.com/lvdbing/go-programming-tour-book/cache/lru"
	"github.com/matryer/is"
)

func TestTourCacheGet(t *testing.T) {
	db := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	}
	getter := cache.GetFunc(func(key string) interface{} {
		log.Println("[From DB] find key", key)
		if val, ok := db[key]; ok {
			return val
		}
		return nil
	})
	tourChche := cache.NewTourCache(getter, lru.New(0, nil))

	is := is.New(t)
	var wg sync.WaitGroup

	for k, v := range db {
		wg.Add(1)
		go func(k, v string) {
			defer wg.Done()
			is.Equal(tourChche.Get(k), v)
			is.Equal(tourChche.Get(k), v)
			is.Equal(tourChche.Get(k), v)
		}(k, v)
	}
	wg.Wait()

	is.Equal(tourChche.Get("unknown"), nil)
	is.Equal(tourChche.Get("unknown"), nil)

	is.Equal(tourChche.Stat().NGet, 14)
	is.Equal(tourChche.Stat().NHit, 8)
}
