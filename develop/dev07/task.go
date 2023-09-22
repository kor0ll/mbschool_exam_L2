package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// Функция для совмещения нескольких done каналов в один, который будет закрыт, если один из его составляющих каналов закроется
func or(channels ...<-chan interface{}) <-chan interface{} {

	// если передан один канал, возвращаем его
	if len(channels) == 1 {
		return channels[0]
	}

	// sync.Once позволяет выполнить действие лишь единожды, что подойдет для закрытия канала
	closeOnce := sync.Once{}
	resultChannel := make(chan interface{})

	for _, channel := range channels {
		//для каждого канала вызываем горутину, которая считывает и передает совмещенному каналу данные
		go func(channel <-chan interface{}) {
			for value := range channel {
				resultChannel <- value
			}

			//как только первый канал будет закрыт, закроется и совмещенный канал
			closeOnce.Do(func() {
				close(resultChannel)
			})
		}(channel)
	}

	return resultChannel
}

// Функция возвращает канал, который будет закрыт по истечению времени, переданного в параметре
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
