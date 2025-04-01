package model

import "time"

type Chat struct {
	ID        int64
	Usernames []string
}

type Message struct {
	ID        int64
	Text      string
	From      string
	Timestamp time.Time
	ChatID    int64
}

type CreateMessage struct {
	Text      string
	From      string
	Timestamp time.Time
	ChatID    int64
}

type CreateChat struct {
	Usernames []string
}
