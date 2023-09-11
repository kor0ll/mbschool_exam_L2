package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Фабричный метод используется для инкапсуляции создания структур в отдельном методе,
который позволяет определять структуру исходя из параметров.
Использование фабричного метода позволяет разделить создание объектов от их использования,
что будет полезно в расширении кода

Мой пример: Предположим, на проекте используется 2 базы данных, Mongo для продакшена и Sqlite для разработки
Для быстрого и удобного подключения к ним можно использовать фабричный метод. Для этого определяем интерфейс
Database, который будут реализовывать структуры, предназначенные для разных баз данных. Также реализуем метод
DatabaseFactory, который возвращает экземпляр Database. Этот метод и будет создавать нужную структуру, исходя из
его параметров. Такой метод удобно использовать, т.к. создание структуры скрывается от клиента, и, если мы захотим
использовать еще одну бд в проекте, достаточно будет создать для нее структуру и реализовать в ней интерфейс Database
*/

type SQLite struct {
	db map[string]string
}

func (sql *SQLite) GetData(query string) (string, bool) {
	if v, ok := sql.db[query]; ok {
		fmt.Println("SQLite")
		return v, true
	}
	return "", false
}
func (sql *SQLite) PutData(query string, data string) {
	sql.db[query] = data
}

type MongoDB struct {
	db map[string]string
}

func (mdb *MongoDB) GetData(query string) (string, bool) {
	if v, ok := mdb.db[query]; ok {
		fmt.Println("MongoDB")
		return v, true
	}
	return "", false
}
func (mdb *MongoDB) PutData(query string, data string) {
	mdb.db[query] = data
}

// общий интерфейс для структур
type Database interface {
	GetData(string) (string, bool)
	PutData(string, string)
}

// фабричный метод, исходя из параметра возвращается одну из структур, реализующую интерфейс Database
func DatabaseFactory(env string) Database {
	switch env {
	case "prod":
		return &MongoDB{
			db: make(map[string]string),
		}
	case "dev":
		return &SQLite{
			db: make(map[string]string),
		}
	default:
		return nil
	}
}
