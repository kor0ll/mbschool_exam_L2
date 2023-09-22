package handler

import (
	"dev11/internal/data"
	"dev11/internal/util"
	"fmt"
	"net/http"
	"strconv"
)

// Handler функция обрабатывает запросы на создание нового события
func UpdateEventHandler(w http.ResponseWriter, r *http.Request, db *data.LocalDB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Недопустимый метод запроса!", http.StatusInternalServerError)
		return
	}

	// создаем обновленное событие, проверяя данные
	event, err := data.ValidateCreateEventParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// получаем id удаляемого события с обработкой ошибок
	eventID := r.FormValue("id")
	if eventID == "" {
		http.Error(w, "Отсутствует идентификатор события", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, "Неверный формат идентификатора события", http.StatusBadRequest)
		return
	}

	// выполняем обновление данных
	err = db.UpdateEvent(id, event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// посылаем ответ
	response := map[string]string{"result": fmt.Sprintf("Обновлено событие с id=%d", id)}

	util.SendJSONResponse(w, http.StatusOK, response)
}
