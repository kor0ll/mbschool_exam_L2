package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
 Фасад - структурный паттерн, суть которого в предоставлении простого интерфейса к сложной системе, содержащей много структур и связей
Например: при заказе пиццы с помощью кредитной карты, в процессе участвуют многие подсистемы, такие как
проверка баланса карты, проверка аккаунта, отправка оповещений и т.д. Фасад в этом случае позволит клиенту работать с
десятками компонентов, использую простой интерфейс

Плюсом этого паттерна является изоляция клиента от элементов сложной подсистемы
Минусом является то, что что фасад может стать "божественным объектом", делающим слишком много и привязанным ко всем классам программы
*/

// описываю примерную систему (упрощенно)
type Account struct {
	name string
}

func newAccount(acName string) *Account {
	return &Account{
		name: acName,
	}
}
func (a *Account) checkAccount(acName string) error {
	if a.name != acName {
		return errors.New("[Аккаунт] Некорректное имя аккаунта")
	}
	fmt.Println("[Аккаунт] Аккаунт верифицирован!")
	return nil
}

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}
func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return errors.New("[Безопасность] Неверный код безопасности")
	}
	fmt.Println("[Безопасность] Код безопасности верифицировн!")
	return nil
}

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("[Кошелек] Баланс успешно изменен!")
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return errors.New("[Кошелек] Недостаточно средств")
	}
	fmt.Println("[Кошелек] Баланс успешно изменен!")
	w.balance = w.balance - amount
	return nil
}

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
}

// сама реализация фасада
func newWalletFacade(acName string, code int) *WalletFacade {
	fmt.Println("Создание нового аккаунта...")
	walletFacade := &WalletFacade{
		account:      newAccount(acName),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
	}
	fmt.Println("Аккаунт успешно создан!")
	return walletFacade
}
func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Пополнение баланса кошелька...")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Снятие денег с кошелька...")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	return nil
}
