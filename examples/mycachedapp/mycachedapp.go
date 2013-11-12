package main

import (
	"github.com/rif/cache2go"
	"fmt"
	"time"
)

type myStruct struct {
	text     string
	moreData []byte
}

func main() {
	// Accessing a new cache table for the first time will create it
	cache := cache2go.Cache("myCache")

	// We will put a new item in the cache. It will expire in 5 seconds
	val := myStruct{"This is a test!", []byte{}}
	cache.Cache("someKey", 5*time.Second, &val)

	// Let's retrieve the item from the cache
	res, err := cache.Value("someKey")
	if err == nil {
		fmt.Println("Found value in cache:", res.Data().(*myStruct).text)
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	// Wait for the item to expire in cache
	time.Sleep(6 * time.Second)
	res, err = cache.Value("someKey")
	if err != nil {
		fmt.Println("Item is not cached (anymore).")
	}

	// Add another item that never expires
	cache.Cache("someKey", 0, &val)

	// cache2go supports a few handy callbacks and loading mechanisms
	cache.SetAboutToDeleteItemCallback(func(e *cache2go.CacheItem) {
		fmt.Println("Deleting:", e.Key(), e.Data().(*myStruct).text, e.CreatedOn())
	})

	// Remove the item from the cache
	cache.Delete("someKey")

	// And wipe the entire cache table
	cache.Flush()
}