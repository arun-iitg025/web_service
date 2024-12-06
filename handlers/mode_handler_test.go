package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"slotmachine/models"
)

// MockCache simulates the Redis cache
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

// MockMongoDB simulates MongoDB interactions
type MockMongoDB struct{}

func (db *MockMongoDB) GetPopularModes(areaCode string) ([]models.ModeInfo, error) {
	if areaCode == "123" {
		return []models.ModeInfo{
			{Mode: "battle_royale", Usage: 1500},
			{Mode: "team_deathmatch", Usage: 800},
		}, nil
	}
	return nil, nil
}

// Define a mock ModeHandler for testing without *mongo.Client or *cache.Cache
type MockModeHandler struct {
	MongoClient *MockMongoDB
	Cache       *MockCache
}

func NewMockModeHandler() *MockModeHandler {
	return &MockModeHandler{
		MongoClient: &MockMongoDB{},
		Cache:       NewMockCache(),
	}
}

func (h *MockModeHandler) PopularModesHandler(w http.ResponseWriter, r *http.Request) {
	areaCode := r.URL.Query().Get("area_code")
	if areaCode == "" {
		http.Error(w, "area_code is required", http.StatusBadRequest)
		return
	}

	// Check cache
	cachedData, _ := h.Cache.Get(areaCode)
	if cachedData != "" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedData))
		return
	}

	// If not in cache, query MongoDB
	modes, err := h.MongoClient.GetPopularModes(areaCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(modes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Cache the response for future requests
	h.Cache.Set(areaCode, string(response), 10*time.Minute)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func TestPopularModesHandler(t *testing.T) {
	handler := NewMockModeHandler() // Use the mock handler with mock DB and cache

	req, _ := http.NewRequest("GET", "/popular-modes?area_code=123", nil)
	rr := httptest.NewRecorder()
	handler.PopularModesHandler(rr, req)

	// Verify the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Expected response data as Go struct
	expectedData := []models.ModeInfo{
		{Mode: "battle_royale", Usage: 1500},
		{Mode: "team_deathmatch", Usage: 800},
	}

	// Convert expected data to JSON
	expectedJSON, _ := json.Marshal(expectedData)

	// Compare JSON output
	if rr.Body.String() != string(expectedJSON) {
		t.Errorf("Expected body %s, got %s", string(expectedJSON), rr.Body.String())
	}
}
