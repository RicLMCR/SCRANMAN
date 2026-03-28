package main

import (
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

type restaurantResponse struct {
	Message string `json:"message"`
}

func (app *Config) Health(w http.ResponseWriter, r *http.Request) {
	resp := healthResponse{
		Status: "ok",
	}

	out, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// get resturants based on cuisine selection
func (app *Config) Restaurants(w http.ResponseWriter, r *http.Request) {

	cuisine := r.URL.Query().Get("cuisine")

	// default value if not provided
	if cuisine == "" {
		http.Error(w, "no cuisine provided", http.StatusBadRequest)
	}

	resp := restaurantResponse{
		Message: "received cuisine selection: " + cuisine,
	}

	out, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
