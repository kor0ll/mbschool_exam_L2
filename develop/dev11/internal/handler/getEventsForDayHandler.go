package handler

import (
	"dev11/internal/data"
	"dev11/internal/util"
	"net/http"
	"time"
)

// Handler фукнция обрабатывает запросы на получение событий на указанную дату
func GetEventsForDayHandler(w http.ResponseWriter, r *http.Request, db *data.LocalDB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// получаем дату
	date := r.FormValue("date")
	if date == "" {
		http.Error(w, "Отсутствует параметр даты", http.StatusBadRequest)
		return
	}

	// Проверяем формат даты
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		http.Error(w, "Неверный формат даты", http.StatusBadRequest)
		return
	}

	// Найдите события на указанную дату
	events := db.GetEventsForDay(date)

	//посылаем ответ
	response := map[string][]*data.Event{"result": events}

	util.SendJSONResponse(w, http.StatusOK, response)
}
