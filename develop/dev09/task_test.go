package main

import (
	"testing"
)

func TestDownloadSite(t *testing.T) {
	// Тестирование загрузки существующего сайта.
	existingURL := "https://go.dev/"
	err := DownloadSite(existingURL)
	if err != nil {
		t.Errorf("Ожидалось успешное выполнение, получено: %v", err)
	}

	// Тестирование загрузки несуществующего сайта.
	nonExistentURL := "https://www.thissitedoesnotexist.com"
	err = DownloadSite(nonExistentURL)
	if err == nil {
		t.Error("Ожидалась ошибка при попытке загрузить несуществующий сайт")
	}
}
