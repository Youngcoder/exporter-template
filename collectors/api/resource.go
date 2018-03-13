package api

import (
	"encoding/json"
	"net/http"
)

type Resource struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetResource(w http.ResponseWriter, r *http.Request) {
	resource := Resource{
		Username: "go",
		Password: "123",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource)
}
