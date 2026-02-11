package fetcher

import (
	"encoding/json"
	"sync"
	"testing"
)

func TestSharedCache_GetSet(t *testing.T) {
	cache := NewSharedCache()

	// Set a value
	data := json.RawMessage(`{"type": "string"}`)
	cache.Set("http://example.com/schema.json", data)

	// Get should return the value
	got, ok := cache.Get("http://example.com/schema.json")
	if !ok {
		t.Fatal("expected to find value in cache")
	}
	if string(got) != string(data) {
		t.Errorf("got %s, want %s", got, data)
	}

	// Has should return true
	if !cache.Has("http://example.com/schema.json") {
		t.Error("expected Has to return true")
	}

	// Non-existent key
	_, ok = cache.Get("http://example.com/other.json")
	if ok {
		t.Error("expected not to find non-existent key")
	}
}

func TestSharedCache_NormalizesURI(t *testing.T) {
	cache := NewSharedCache()

	// Set with fragment
	data := json.RawMessage(`{"type": "string"}`)
	cache.Set("http://example.com/schema.json#/defs/Foo", data)

	// Get without fragment should find it (normalized)
	got, ok := cache.Get("http://example.com/schema.json")
	if !ok {
		t.Fatal("expected to find value (URI should be normalized)")
	}
	if string(got) != string(data) {
		t.Errorf("got %s, want %s", got, data)
	}

	// Get with different fragment should also find it
	got2, ok := cache.Get("http://example.com/schema.json#/defs/Bar")
	if !ok {
		t.Fatal("expected to find value with different fragment")
	}
	if string(got2) != string(data) {
		t.Errorf("got %s, want %s", got2, data)
	}
}

func TestSharedCache_ToCache(t *testing.T) {
	cache := NewSharedCache()

	cache.Set("http://a.com/1.json", json.RawMessage(`{"a": 1}`))
	cache.Set("http://b.com/2.json", json.RawMessage(`{"b": 2}`))

	snapshot := cache.ToCache()

	if len(snapshot) != 2 {
		t.Fatalf("expected 2 items in snapshot, got %d", len(snapshot))
	}

	// Verify items
	if string(snapshot["http://a.com/1.json"]) != `{"a": 1}` {
		t.Error("missing or wrong value for key a")
	}
	if string(snapshot["http://b.com/2.json"]) != `{"b": 2}` {
		t.Error("missing or wrong value for key b")
	}

	// Modifying snapshot should not affect original cache
	snapshot["http://c.com/3.json"] = json.RawMessage(`{"c": 3}`)
	if cache.Len() != 2 {
		t.Error("modifying snapshot affected original cache")
	}
}

func TestSharedCache_ConcurrentAccess(t *testing.T) {
	cache := NewSharedCache()
	const goroutines = 100
	const iterations = 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				uri := "http://example.com/schema.json"
				data := json.RawMessage(`{"id": "test"}`)

				// Mix of operations
				cache.Set(uri, data)
				cache.Get(uri)
				cache.Has(uri)
				cache.Len()
				cache.ToCache()
			}
		}(i)
	}

	wg.Wait()

	// If we get here without race detector complaints, the test passes
	if cache.Len() != 1 {
		t.Errorf("expected 1 item in cache, got %d", cache.Len())
	}
}
