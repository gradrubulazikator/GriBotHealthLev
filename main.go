package main

import (
    "log"
    "time"
    "GriBotHealthLev/internal" // Импортируем пакет internal, чтобы использовать его функции и константы
)

func healthReminder() {
    messages := []string{
        "Пора сделать разминку!",
        "Не забудьте выпить воды.",
        "Пора проверить уровень стресса.",
    }

    for {
        for _, msg := range messages {
            if err := internal.SendMessage(msg); err != nil { // Используем internal.SendMessage
                log.Printf("Ошибка отправки напоминания: %v", err)
            }
            time.Sleep(internal.RemindDelay) // Используем internal.RemindDelay
        }
    }
}

func main() {
    log.Println("Запуск GriBotHealthLev...")
    go healthReminder()
    select {}
}

