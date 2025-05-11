package repositories

import (
	"database/sql"

	"alex.com/application-bot/internal/domain/entities"
)

type UserRepository struct {
	Connection *sql.DB
}

func (userRepository UserRepository) GetByChatId(id int64) (*entities.User, error) {
	row := userRepository.Connection.QueryRow(`select id, chat_id, created_at, updated_at from users where chat_id = ?`, id)
	var u entities.User
	err := row.Scan(&u.ID, &u.ChatID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &entities.User{}, err
	}

	return &u, nil
}

func (userRepository UserRepository) Create(chatId int64) (*int, error) {
	sql := `insert into users (chat_id, created_at, updated_at) values (?, datetime('now'), datetime('now')) returning id`
	row := userRepository.Connection.QueryRow(sql, chatId)
	var id int

	err := row.Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func NewUserRepository(connection *sql.DB) *UserRepository {
	return &UserRepository{Connection: connection}
}
