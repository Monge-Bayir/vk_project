package main

import (
    "fmt"
    "log"
    "mattermost-voting-bot/pkg/voting"
    "mattermost-voting-bot/pkg/storage"
    "mattermost-voting-bot/pkg/logging"
)

func main() {
    // Инициализация логирования
    logging.InitLogger()

    // Подключение к базе данных Tarantool
    conn, err := storage.ConnectTarantool()
    if err != nil {
        log.Fatalf("Ошибка подключения к Tarantool: %v", err)
    }

    // Создание и обработка команд бота (например, в Mattermost API)
    fmt.Println("Бот запущен и слушает команды...")
    // Здесь будет код для интеграции с Mattermost API и команд бота
}
