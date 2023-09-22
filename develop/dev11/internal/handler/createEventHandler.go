package handler

import (
	"dev11/internal/data"
	"dev11/internal/util"
	"fmt"
	"net/http"
)

// Handler функция обрабатывает запросы на создание нового события
func CreateEventHandler(w http.ResponseWriter, r *http.Request, db *data.LocalDB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Недопустимый метод запроса!", http.StatusInternalServerError)
		return
	}

	//создаем новое событие, проверяя данные
	event, err := data.ValidateCreateEventParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// заносим новое событие в хранилище
	eventId := db.AddEvent(event)

	//посылаем ответ
	response := map[string]string{"result": fmt.Sprintf("Событие добавлено в хранилище с id=%d", eventId)}

	util.SendJSONResponse(w, http.StatusOK, response)
}
