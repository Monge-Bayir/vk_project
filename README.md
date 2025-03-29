# Mattermost Voting Bot

## Описание
Этот бот позволяет пользователям создавать голосования, голосовать, просматривать результаты и завершать голосования в Mattermost.

## Требования
- Go 1.18+
- Docker и Docker Compose

## Установка
1. Клонируйте репозиторий:
   ```
   git clone https://github.com/yourusername/mattermost-voting-bot.git
   cd mattermost-voting-bot
   ```

2. Построение Docker контейнера:
   ```
   docker-compose build
   ```

3. Запуск:
   ```
   docker-compose up
   ```

## Использование
- `/createpoll` - создать новое голосование
- `/vote <poll_id> <option>` - проголосовать
- `/results <poll_id>` - посмотреть результаты голосования
- `/endpoll <poll_id>` - завершить голосование
- `/deletepoll <poll_id>` - удалить голосование
