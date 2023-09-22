package util

import (
	"dev11/internal/data"
	"encoding/json"
	"net/http"
)

// Функция для отправки JSON-ответа
func SendJSONResponse[T string | []*data.Event](w http.ResponseWriter, status int, data map[string]T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Не удалось закодировать JSON", http.StatusInternalServerError)
	}
}
