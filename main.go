package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func getOrSet(client *memcache.Client) error {

	log.Printf("before Get")
	item, err := client.Get("key1")
	if err == memcache.ErrCacheMiss {
		log.Printf("Get cache miss")
	} else if err != nil {
		return fmt.Errorf("client.Get error: %w", err)
	}

	log.Printf("after Get item = %+v", item)
	if item != nil {
		log.Printf("Get returned item.key = %q item.value = %q", item.Key, string(item.Value))
		return nil
	}

	item = &memcache.Item{
		Key:        "key1",
		Value:      []byte("value1"),
		Expiration: 10,
	}
	log.Printf("before Set item = %+v", item)
	err = client.Set(item)
	if err != nil {
		return fmt.Errorf("client.Set error %w", err)
	}
	log.Printf("after Set")

	return nil
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	log.Printf("begin main")

	servers := []string{"localhost:11211"}
	if len(os.Args) > 1 {
		servers = os.Args[1:]
	}

	log.Printf("servers = %v", servers)

	client := memcache.New(servers...)

	log.Printf("created client")

	for {
		err := getOrSet(client)
		if err != nil {
			log.Printf("getOrSet error: %v", err)
		}

		time.Sleep(time.Second)
	}
}
