package data

import (
	"fmt"
	"net/http"
	"time"
)

// Event является структурой для событий
type Event struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

// Фукнция выполняет валидацию параметров и создает новое событие
func ValidateCreateEventParams(r *http.Request) (*Event, error) {
	title := r.FormValue("title")
	date := r.FormValue("date")
	details := r.FormValue("description")

	// Валидация параметров
	if title == "" || date == "" || details == "" {
		return nil, fmt.Errorf("отсутствуют обязательные параметры")
	}

	// Проверяем формат даты
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("неверный формат даты")
	}

	newEvent := Event{
		Title:       title,
		Date:        date,
		Description: details,
	}

	return &newEvent, nil
}
