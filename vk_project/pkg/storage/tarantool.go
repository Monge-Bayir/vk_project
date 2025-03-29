package storage

import (
    "github.com/tarantool/go-tarantool"
    "log"
)

// Подключение к базе данных Tarantool
func ConnectTarantool() (*tarantool.Connection, error) {
    conn, err := tarantool.Connect("localhost:3301")
    if err != nil {
        log.Fatalf("Не удалось подключиться к Tarantool: %v", err)
        return nil, err
    }
    return conn, nil
}

// Сохранение нового голосования в Tarantool
func SavePoll(conn *tarantool.Connection, poll Poll) error {
    _, err := conn.Insert("polls", []interface{}{poll.ID, poll.Options, poll.Votes, poll.CreatorID, poll.Active})
    return err
}

// Получение голосования по ID
func GetPollByID(conn *tarantool.Connection, id string) (Poll, error) {
    result, err := conn.Select("polls", "primary", 0, 1, tarantool.IterEq, []interface{}{id})
    if err != nil || len(result) == 0 {
        return Poll{}, err
    }
    return result[0].Tuple, nil
}
