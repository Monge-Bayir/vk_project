package voting

import "fmt"

// Структура для голосования
type Poll struct {
    ID        string            `json:"id"`
    Options   []string          `json:"options"`
    Votes     map[string]int    `json:"votes"` // ключ - вариант, значение - количество голосов
    CreatorID string            `json:"creator_id"`
    Active    bool              `json:"active"`
}

// Функция для создания нового голосования
func CreatePoll(options []string, creatorID string) Poll {
    poll := Poll{
        ID:        generatePollID(),
        Options:   options,
        Votes:     make(map[string]int),
        CreatorID: creatorID,
        Active:    true,
    }

    return poll
}

// Генерация ID для нового голосования (простой пример)
func generatePollID() string {
    return fmt.Sprintf("poll-%d", len(polls) + 1) // Пример генерации ID
}
