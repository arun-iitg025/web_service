package cache

import (
	"testing"
	"time"
)

// MockCache for testing cache operations
type MockCache struct {
	data map[string]string
}

func NewMockCache() *MockCache {
	return &MockCache{data: make(map[string]string)}
}

func (c *MockCache) Get(key string) (string, error) {
	if value, exists := c.data[key]; exists {
		return value, nil
	}
	return "", nil
}

func (c *MockCache) Set(key string, value string, ttl time.Duration) error {
	c.data[key] = value
	return nil
}

func TestCache_SetAndGet(t *testing.T) {
	cache := NewMockCache()
	key := "testKey"
	value := "testValue"
	ttl := 10 * time.Second

	if err := cache.Set(key, value, ttl); err != nil {
		t.Fatalf("Failed to set cache: %v", err)
	}

	got, err := cache.Get(key)
	if err != nil {
		t.Fatalf("Failed to get cache: %v", err)
	}

	if got != value {
		t.Errorf("Expected value %s, got %s", value, got)
	}
}
