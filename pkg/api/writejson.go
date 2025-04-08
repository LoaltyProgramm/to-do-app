package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJson(w http.ResponseWriter, data any) {
	responce, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("Error marshal data: %v", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(responce)
} 