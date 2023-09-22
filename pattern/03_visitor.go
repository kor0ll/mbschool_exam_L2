package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*

Посетитель - поведенческий паттерн, который позволяет добавлять в программу новые операции,
не изменяя классы объектов, над которыми эти операции могут выполняться
Например: есть 3 структуры разных фигур - квадрат, треугольник и круг. Все они реализуют интерфейс Фигуры
Нужно добавить в структуры фукнцию, возвращающую площадь фигуры. Простым решением может быть добавление новой функции
getArea в интерфейс фигуры и реализовать его в каждой структуре. Но так мы рискуем наделать ошибок в коде каждый раз, когда
необходимо добавить новое поведение. Поэтому, можно решить эту проблему с помощью паттерна посетитель

Плюсами посетителя являются упрощение добавления операций, работающих со сложными структурами объектов и объединение родственных операций в одном классе
Минусами является возможность нарушения инкапсуляции элементов, в случае если новое поведение затрагивает приватные поля

*/

// реализуем интерфейс фигуры и 3 структуры. Для работы посетителя необходимо включить во все структуры метод accept
// это может показаться странным, ведь так или иначе пришлось изменить все структуры, но в случае с посетителем это
// пришлось сделать лишь однажды
type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForrectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

// определяем интерфейс посетителя. Он должен содержать методы с нужным поведением для каждого типа фигуры
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForrectangle(*Rectangle)
}

// теперь реализуем конкретного посетителя
type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Расчет площади для квадрата")
}
func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Расчет площади для круга")
}
func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
	fmt.Println("Расчет площади для треугольника")
}

//этих конкретных посетителей может быть сколько угодно, и менять структуры фигур уже не нужно
