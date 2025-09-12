package main

import (
	"fmt"
	"github.com/GrayMan124/pokedexcli/internal/pokecache"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "jakPanhesusPowiedzial.com",
			val: []byte("Tak jak pan jezus powiedzial"),
		},
		{
			key: "Jacob.Galuska",
			val: []byte("To mala pietruszka"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Running case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Incorrect val")
				return
			}
		})

	}

}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime / 2

	cache := pokecache.NewCache(baseTime)
	cache.Add("Jacob.Galuska", []byte("To mala pietruszka"))

	_, ok := cache.Get("Jacob.Galuska")
	if !ok {
		t.Errorf("Expected to find Jacob")
	}

	time.Sleep(waitTime + 2*time.Millisecond)

	cache.Add("Jak.Pan", []byte("Tak jak Pan"))

	time.Sleep(waitTime)

	_, ok = cache.Get("Jacob.Galuska")
	if ok {
		t.Errorf("Expected to not find Jacob")
		return
	}

	_, ok = cache.Get("Jak.Pan")
	if !ok {
		t.Errorf("Expected to find Jak Pan")
		return
	}
}
