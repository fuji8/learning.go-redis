package redis_test

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type Object struct {
	Str string
	Num int
}

func Example_twoCaches() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"localhost": ":6379",
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	ctx := context.TODO()
	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   time.Hour,
	}); err != nil {
		panic(err)
	}

	anothercache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	var wanted Object
	if err := anothercache.Get(ctx, key, &wanted); err == nil {
		fmt.Println(wanted)
	}

	// Output: {mystring 42}
}

func Example_twoCachesLocalCache() {
	redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"localhost": ":6379",
		},
	})

	mycache := cache.New(&cache.Options{
		// Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	ctx := context.TODO()
	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   time.Hour,
	}); err != nil {
		panic(err)
	}

	anothercache := cache.New(&cache.Options{
		// Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	var wanted Object
	if err := anothercache.Get(ctx, key, &wanted); err == nil {
		fmt.Println(wanted)
	}

	// Output:
}

func Example_twoCachestwoDB() {
	ring0 := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"localhost": ":6379",
		},
		DB: 0,
	})
	ring1 := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"localhost": ":6379",
		},
		DB: 1,
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring0,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	ctx := context.TODO()
	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   time.Hour,
	}); err != nil {
		panic(err)
	}

	anothercache := cache.New(&cache.Options{
		Redis:      ring1,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	var wanted Object
	if err := anothercache.Get(ctx, key, &wanted); err == nil {
		fmt.Println(wanted)
	}

	// Output:
}
