package handler

import (
	"dev11/internal/data"
	"dev11/internal/util"
	"net/http"
	"time"
)

// Handler фукнция обрабатывает запросы на получение событий в определенном диапазоне
func GetEventsInDateRangeHandler(w http.ResponseWriter, r *http.Request, db *data.LocalDB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// получаем начальную и конечную даты
	startDate := r.FormValue("startDate")
	endDate := r.FormValue("endDate")
	if startDate == "" || endDate == "" {
		http.Error(w, "Отсутствует параметр даты", http.StatusBadRequest)
		return
	}

	// Проверяем формат дат
	_, err1 := time.Parse("2006-01-02", startDate)
	_, err2 := time.Parse("2006-01-02", endDate)
	if err1 != nil || err2 != nil {
		http.Error(w, "Неверный формат даты", http.StatusBadRequest)
		return
	}

	// Найдите события в диапазоне указанных дат
	events := db.GetEventsInDateRange(startDate, endDate)

	//посылаем ответ
	response := map[string][]*data.Event{"result": events}

	util.SendJSONResponse(w, http.StatusOK, response)
}
