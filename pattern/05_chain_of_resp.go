package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
/*
Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.
Клиенту не нужно знать все звенья цепи, достаточно вызвать метод у первого обработчика в цепи.
Цепочка обязанностей применима в случаях, когда для одного запроса есть несколько обработчиков
Например: существует госпиталь, в нем несколько разных помещений: приемное отделение, доктор, комната медикаментов и касса. Когда пациент приходит,
первым делом он попадает в приемное отделение и далее по цепочке помещений, в которой каждое отправляет его по ней дальше сразу после
выполнения своей функции

Плюсами цепочки обязанностей являются реализация принципа единственной обязанности и принципа открытости/закрытости
Минусом может быть то, что запрос может остаться никем не обработан за всю цепочку
*/

// структура для пациента
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// реализуем общий интерфейс для обработчиков, в нем должен быть метод execute, который и выполняет определенные инструкции и посылает запрос дальше по цепочке
type MedHandler interface {
	execute(*Patient)
	setNext(MedHandler)
}

// далее реализуем конкретные обработчики для данного примера
// у каждого из них есть поле next, указывающее на следующий обработчик, который необходимо вызвать. Так выстраивается цепочка
// обработчик для приемного отделения
type Reception struct {
	next MedHandler
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next MedHandler) {
	r.next = next
}

// обработчик для доктора
type Doctor struct {
	next MedHandler
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next MedHandler) {
	d.next = next
}

// обработчик для комнаты с медикаментами
type Medical struct {
	next MedHandler
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next MedHandler) {
	m.next = next
}

// обработчик для кассы
type Cashier struct {
	next MedHandler
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next MedHandler) {
	c.next = next
}

//таким образом можно выстроить цепочку обработчиков, которые будут взаимодействовать с запросом
