package voting

import (
    "fmt"
    "mattermost-voting-bot/pkg/logging"
)

// Обработчик команды для создания голосования
func HandleCreatePollCommand(options []string, creatorID string) string {
    poll := CreatePoll(options, creatorID)

    // Логирование действия
    logging.LogAction(fmt.Sprintf("Создано новое голосование с ID: %s", poll.ID))

    return fmt.Sprintf("Голосование создано! ID: %s, варианты: %v", poll.ID, poll.Options)
}

// Обработчик команды для голосования
func HandleVoteCommand(pollID string, option string) string {
    // Логика голосования, добавление голоса за выбранный вариант
    logging.LogAction(fmt.Sprintf("Пользователь проголосовал за вариант: %s в голосовании: %s", option, pollID))
    return fmt.Sprintf("Вы проголосовали за вариант: %s", option)
}
