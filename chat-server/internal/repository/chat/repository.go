package chat

import (
	"chat-server/internal/client/db"
	"chat-server/internal/model"
	"chat-server/internal/repository"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"log"
)

const (
	tableChatName       = "chat"
	idChatColumn        = "id"
	usernamesChatColumn = "usernames"

	tableMsgName    = "message"
	fromColumn      = "from"
	toColumn        = "to"
	timestampColumn = "timestamp"
	textColumn      = "text"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r repo) SendMessage(ctx context.Context, message *model.CreateMessage) error {
	const op = "chat.SendMessage"

	builder := sq.Insert(tableMsgName).
		PlaceholderFormat(sq.Dollar).
		Columns(fromColumn, toColumn, timestampColumn, textColumn).
		Values(message.From, message.To, message.Timestamp, message.Text).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error in send msg %w, %d", err, id)
	}

	log.Println("id:", id)
	return nil
}

func (r repo) Create(ctx context.Context, chat *model.CreateChat) (int64, error) {
	const op = "chat.Create"

	builder := sq.Insert(tableChatName).
		PlaceholderFormat(sq.Dollar).
		Columns(usernamesChatColumn).
		Values(chat.Usernames).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("error in create chat %w", err)
	}

	log.Println("id:", id)
	return id, nil
}

func (r repo) Delete(ctx context.Context, id int64) error {
	const op = "chat.Delete"

	builder := sq.Delete(tableChatName).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idChatColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	row := r.db.DB().QueryRowContext(ctx, q, args...)
	if row == nil {
		return fmt.Errorf("cannot delete chat with id: %d", id)
	}

	return nil
}
