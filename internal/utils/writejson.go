package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data any) {
	responce, err := json.Marshal(data)
	if err != nil {
		http.Error(w, `{"error":"Failed to serialize response"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(responce)
} 