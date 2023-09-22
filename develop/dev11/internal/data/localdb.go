package data

import (
	"fmt"
	"sync"
)

// Локальное хранилище данных
type LocalDB struct {
	mu        sync.Mutex
	data      map[int]*Event
	idCounter int
}

// Функция возвращает ссылку на экземпляр LocalDB
func CreateLocalDB() *LocalDB {
	data := make(map[int]*Event)
	return &LocalDB{sync.Mutex{}, data, 0}
}

// Функция выполняет поиск события по Id
func (db *LocalDB) GetEventById(id int) (*Event, error) {
	event, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("элемент с Id: %d не найден", id)
	}
	return event, nil
}

// Функция обновляет данные события, указанного по id, в хранилище
func (db *LocalDB) UpdateEvent(id int, newEvent *Event) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.data[id]
	if !ok {
		return fmt.Errorf("элемент с Id: %d не найден", id)
	} else {
		db.data[id] = newEvent
		return nil
	}
}

// Функция удаляет указанное по Id события из хранилища
func (db *LocalDB) DeleteEvent(id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.data[id]
	if !ok {
		return fmt.Errorf("элемент с Id: %d не найден", id)
	}

	delete(db.data, id)
	return nil
}

// Функция добавляет новый Event в хранилище
func (db *LocalDB) AddEvent(event *Event) int {
	db.mu.Lock()
	defer db.mu.Unlock()

	currentId := db.idCounter

	db.data[currentId] = event
	db.idCounter++

	return currentId
}

// Функция возвращает слайс событий с датой указанной в date
func (db *LocalDB) GetEventsForDay(date string) []*Event {
	db.mu.Lock()
	defer db.mu.Unlock()

	events := []*Event{}

	for _, v := range db.data {
		if v.Date == date {
			events = append(events, v)
		}
	}

	return events
}

// Функция возвращает слайс событий в диапазоне указанных дат
func (db *LocalDB) GetEventsInDateRange(startDate string, endDate string) []*Event {
	db.mu.Lock()
	defer db.mu.Unlock()

	events := []*Event{}

	for _, v := range db.data {
		if v.Date >= startDate && endDate >= v.Date {
			events = append(events, v)
		}
	}

	return events
}
