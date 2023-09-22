package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	t.Run("Single channel", func(t *testing.T) {
		doneChan := sig(2 * time.Second)
		select {
		case <-or(doneChan):
			// Функция or не должна закрывать doneChan до истечения 2 секунд.
		case <-time.After(3 * time.Second):
			t.Fatal("or(doneChan) должна закрыться не раньше, чем через 2 секунды")
		}
	})

	t.Run("Multiple channels", func(t *testing.T) {
		doneChan1 := sig(2 * time.Second)
		doneChan2 := sig(3 * time.Second)
		doneChan3 := sig(1 * time.Second)

		mergedChan := or(doneChan1, doneChan2, doneChan3)

		select {
		case <-mergedChan:
			// По крайней мере один из каналов должен быть закрыт.
		case <-time.After(4 * time.Second):
			t.Fatal("or(doneChan1, doneChan2, doneChan3) должна закрыться не раньше, чем через 3 секунды")
		}
	})

}
