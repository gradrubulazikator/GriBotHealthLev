package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const telegramAPI = "https://api.telegram.org/bot" // Основной URL API Telegram

// Конфигурация бота
var (
	botToken    = "7215760944:AAF3fveH7KFHOe1HVvxICFGbx27fzQGAkMU" // Прямое значение токена
	chatID      = "5838863003"                                     // Прямое значение ID чата
	remindDelay = time.Hour * 6                                    // Интервал напоминаний
)

// Функция для отправки сообщения
func sendMessage(message string) error {
	endpoint := fmt.Sprintf("%s%s/sendMessage", telegramAPI, botToken)
	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("text", message)

	resp, err := http.PostForm(endpoint, data)
	if err != nil {
		return fmt.Errorf("ошибка отправки сообщения: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка ответа от Telegram: %s", resp.Status)
	}
	return nil
}

// Функция для отправки напоминаний
func healthReminder() {
	messages := []string{
		"Пора сделать разминку!",
		"Не забудьте выпить воды.",
		"Пора проверить уровень стресса.",
	}

	for {
		for _, msg := range messages {
			if err := sendMessage(msg); err != nil {
				log.Printf("Ошибка отправки напоминания: %v", err)
			}
			time.Sleep(remindDelay)
		}
	}
}

func main() {
	log.Println("Запуск GriBotHealthLev...")
	go healthReminder()

	// Сохранение работы бота без блокировки основного потока
	select {}
}

