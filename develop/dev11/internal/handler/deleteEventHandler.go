package handler

import (
	"dev11/internal/data"
	"dev11/internal/util"
	"fmt"
	"net/http"
	"strconv"
)

// Handler фукнция обрабатывает запросы на удаление события
func DeleteEventHandler(w http.ResponseWriter, r *http.Request, db *data.LocalDB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Недопустимый метод запроса", http.StatusInternalServerError)
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

	// удаляем событие из хранилища
	err = db.DeleteEvent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	//посылаем ответ
	response := map[string]string{"result": fmt.Sprintf("Событие с id=%d удалено из хранилища", id)}

	util.SendJSONResponse(w, http.StatusOK, response)
}
