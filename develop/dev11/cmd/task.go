package main

import (
	"dev11/internal/data"
	"dev11/internal/handler"
	"log"
	"net/http"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

// Функция для установки middleware для логов
func middleWareForLogs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func main() {
	// создаем экземпляр LocalDB
	db := data.CreateLocalDB()

	// создаем и настраиваем маршрутизатор
	mux := http.NewServeMux()
	middleware := middleWareForLogs(mux)

	// регистрируем handler'ы на каждый запрос
	mux.HandleFunc("/create_event", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateEventHandler(w, r, db)
	})
	mux.HandleFunc("/update_event", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateEventHandler(w, r, db)
	})
	mux.HandleFunc("/delete_event", func(w http.ResponseWriter, r *http.Request) {
		handler.DeleteEventHandler(w, r, db)
	})
	mux.HandleFunc("/events_for_day", func(w http.ResponseWriter, r *http.Request) {
		handler.GetEventsForDayHandler(w, r, db)
	})
	mux.HandleFunc("/events_for_week", func(w http.ResponseWriter, r *http.Request) {
		handler.GetEventsInDateRangeHandler(w, r, db)
	})
	mux.HandleFunc("/events_for_month", func(w http.ResponseWriter, r *http.Request) {
		handler.GetEventsInDateRangeHandler(w, r, db)
	})

	// указываем порт и запускаем сервер
	port := ":8080"
	log.Printf("Сервер запускается по порту %s\n", port)
	if err := http.ListenAndServe(port, middleware); err != nil {
		log.Fatal("Ошибка сервера:", err)
	}
}
