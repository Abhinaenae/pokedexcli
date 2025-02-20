package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(5 * time.Second)
	if cache.cache == nil {
		t.Errorf("Cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("%s was not found", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("got %s want %s", val, c.val)
				return
			}

		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("cache should exist have been reaped")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("should have been reaped")
		return
	}
}
