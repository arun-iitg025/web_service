package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"slotmachine/cache"
	"slotmachine/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ModeHandler struct {
	MongoClient *mongo.Client
	Cache       *cache.Cache
}

func (h *ModeHandler) PopularModesHandler(w http.ResponseWriter, r *http.Request) {
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

	//If not in cache, query MongoDB
	modes, err := models.GetPopularModes(h.MongoClient, areaCode)
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
