package logging

import (
    "log"
    "os"
)

var Logger *log.Logger

// Инициализация логирования
func InitLogger() {
    Logger = log.New(os.Stdout, "VOTING_BOT: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Логирование действий
func LogAction(action string) {
    Logger.Println(action)
}
