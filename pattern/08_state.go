package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/
/*
Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.
Такой паттерн применим в случаях, когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
причём типов состояний много, и их код часто меняется, или когда код класса содержит множество больших условных операторов, которые выбирают
поведение в зависимости от текущих значений полей класса
Пример: предположим, существует торговый автомат который выдает 1 товар. Он может пребывать в 4 состояниях:
hasItem (имеетПредмет), noItem (неИмеетПредмет), itemRequested (выдаётПредмет), hasMoney (получилДеньги), а также 4 различных действия:
выбрать предмет, добавить предмет, ввести деньги, выдать предмет.
В таком примере состояние торгового автомата постоянно будет меняться, и в зависимости от них он по разному будет отвечать на одни и те же запросы.
Такой торговый автомат можно реализовать с помощью паттерна состояние.

Плюсами данного паттерна являются упрощение кода контекста (в данном примере VendingMachine) и избавление от множества больших условных операторов машины состояний
Минусом является неоправданное усложнение кода, если этих состояний будет мало, или же если они меняются редко

*/

// определяем интерфейс для состояний
type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

// определяем структуру автомата, он хранит все возможные состояния, а также текущее состояние
type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State

	itemCount int
	itemPrice int
}

// все 4 заявленных действия автомат делегирует на структуру текущего состояния
func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

// также реализуем функцию смены состояния
func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

// реализуем фукнцию добавления товара
func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

//Далее реализуем различные состояния автомата. У каждого из них есть поле с ссылкой на структуру автомата, т.к разные состояния должны иметь
//возможность менять текущее состояние автомата

// состояние "не имеет предметов"
type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

// состояние "имеет предмет"
type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}
func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}

// состояние "выдает предмет"
type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

// состояние "получил деньги"
type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

// так как структура автомата получилась сложной, реализуем конструктор
func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

//таким образом состояния будут взаимодействовать друг с другом и менять друг друга, обеспечивая нормальную работу автомата
