package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Строитель - порождающий паттерн, позволяющий создавать объекты пошагово
С помощью него возможно создавать продукты с разными свойствами, используя один и тот же процесс строительства
Когда нужный продукт (структура) сложный и требует нескольких шагов для построения, в таком случае
несколько конструкторных методов подойдут лучше, чем один громадный конструктор, т.к он будет выглядеть монструозно, а некоторые
параметры не будут использоваться часто. Строитель предлагает выносить конструирование объекта за пределы его класса,
поручив это дело отдельным объектом, которые называются строителями
Например: нам нужно создавать структуру House, и, так как дома бывают разных конструкций и материалов, можно воспользоваться строителями
В моем примере есть возможность создать 2 типа конструкций - нормальный дом и иглу

Плюсами этого паттерна являются изоляция сложного кода сборки продукта от его бизнес логики, а также использование одного и того же кода
для создания продуктов различной конфигурации
Минусами является усложнение кода программы из за введения дополнительных классов
*/

type House struct {
	windowType string
	doorType   string
	floor      int
}

// билдеры должны иметь общий интерфейс
type IHouseBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IHouseBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

// реализация билдера для нормального дома
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// реализация билдера для иглу
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// можно выделить вызовы метода строителя в отдельную структуру, она называется Директор
// и задает порядок шагов строительства
type Director struct {
	builder IHouseBuilder
}

func newDirector(b IHouseBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IHouseBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
