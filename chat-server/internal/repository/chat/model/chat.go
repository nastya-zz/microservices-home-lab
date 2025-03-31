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
	To        string
	Timestamp time.Time
}

type CreateMessage struct {
	Text      string
	From      string
	To        string
	Timestamp time.Time
}

type CreateChat struct {
	Usernames []string
}
