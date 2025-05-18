package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"alex.com/application-bot/internal/application/constants"
	"alex.com/application-bot/internal/domain/entities"
)

type ApplicationRepository struct {
	Connection *sql.DB
}

func (repo ApplicationRepository) GetLastByUserId(userId int) (*entities.Application, error) {
	query := `
		select 
			id,
			chat_id,
			telegram_id,
			user_id, 
			country,
			mark_or_conditions,
			budget,
			steering_wheel_type,
			city,
			person_name,
			person_phone,
			step,
			created_at, 
			updated_at,
			sended_telegram,
			sended_bitrix
		from
			applications 
		where user_id = ?
		order by id desc
		limit 1
	`

	row := repo.Connection.QueryRow(query, userId)

	var appl entities.Application

	err := row.Scan(
		&appl.ID,
		&appl.ChatId,
		&appl.TelegramId,
		&appl.UserId,
		&appl.Country,
		&appl.MarkOrConditions,
		&appl.Budget,
		&appl.SteeringWheelType,
		&appl.City,
		&appl.PersonName,
		&appl.PersonPhone,
		&appl.Step,
		&appl.CreatedAt,
		&appl.UpdatedAt,
		&appl.SendedToTelegram,
		&appl.SendedToBitfix,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &appl, nil
}

func (repo ApplicationRepository) CreateEmpty(userdId int, chatId int64, telegramId string) error {
	sql := `insert into applications(user_id, chat_id, telegram_id, created_at, updated_at) values (?, ?, ?, datetime('now'), datetime('now'))`

	_, err := repo.Connection.Exec(sql, userdId, chatId, telegramId)

	if err != nil {
		log.Printf("error while trying to create empty application, more: %s", err)
		return err
	}

	return nil
}

func (repo ApplicationRepository) Update(appl *entities.Application) error {
	sql := `
        update applications 
        set 
            country = ?,
            mark_or_conditions = ?,
            budget = ?,
            steering_wheel_type = ?,
            city = ?,
            person_name = ?,
            person_phone = ?,
            step = ?,
            updated_at = datetime('now'),
			sended_telegram = ?,
			sended_bitrix = ?,
			chat_id = ?,
			telegram_id = ?
        where 
            id = ?
    `

	result, err := repo.Connection.Exec(sql,
		appl.Country,
		appl.MarkOrConditions,
		appl.Budget,
		appl.SteeringWheelType,
		appl.City,
		appl.PersonName,
		appl.PersonPhone,
		appl.Step,
		appl.SendedToTelegram,
		appl.SendedToBitfix,
		appl.ChatId,
		appl.TelegramId,
		appl.ID,
	)

	if err != nil {
		log.Printf("error while executing update, more: %s", err)
		return fmt.Errorf("error while executing update, more: %s", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error while tryning to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("there is no rows affected for application: %d", appl.ID)
	}

	return nil
}

func (repo ApplicationRepository) GetSubbmited() ([]*entities.Application, error) {
	sql := `
		select 
			id, 
			user_id, 
			chat_id,
			telegram_id,
			country,
			mark_or_conditions,
			budget,
			steering_wheel_type,
			city,
			person_name,
			person_phone,
			step,
			created_at, 
			updated_at,
			sended_telegram,
			sended_bitrix
		from 
			applications 
		where step = ? and not sended_telegram
	`

	rows, err := repo.Connection.Query(sql, constants.MaxStepType)

	if err != nil {
		if rows != nil {
			defer rows.Close()
		}

		return nil, err
	}

	var applications []*entities.Application

	for rows.Next() {
		var appl entities.Application
		err := rows.Scan(
			&appl.ID,
			&appl.UserId,
			&appl.ChatId,
			&appl.TelegramId,
			&appl.Country,
			&appl.MarkOrConditions,
			&appl.Budget,
			&appl.SteeringWheelType,
			&appl.City,
			&appl.PersonName,
			&appl.PersonPhone,
			&appl.Step,
			&appl.CreatedAt,
			&appl.UpdatedAt,
			&appl.SendedToTelegram,
			&appl.SendedToBitfix,
		)

		if err != nil {
			return nil, err
		}

		applications = append(applications, &appl)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return applications, nil
}

func NewApplicationRepository(conn *sql.DB) *ApplicationRepository {
	return &ApplicationRepository{Connection: conn}
}
