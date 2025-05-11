package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"alex.com/application-bot/internal/domain/entities"
)

type ApplicationRepository struct {
	Connection *sql.DB
}

func (repo ApplicationRepository) GetLastByUserId(userId int) (*entities.Application, error) {
	query := `
		select 
			id, 
			user_id, 
			country,
			mark_or_conditions,
			budget,
			steering_wheel_type,
			city,
			person_name,
			person_phone,
			submitted, 
			step,
			created_at, 
			updated_at
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
		&appl.UserId,
		&appl.Country,
		&appl.MarkOrConditions,
		&appl.Budget,
		&appl.SteeringWheelType,
		&appl.City,
		&appl.PersonName,
		&appl.PersonPhone,
		&appl.Submitted,
		&appl.Step,
		&appl.CreatedAt,
		&appl.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &appl, nil
}

func (repo ApplicationRepository) GetByUserId(userId int) ([]*entities.Application, error) {
	sql := `
		select 
			id, 
			user_id, 
			country,
			mark_or_conditions,
			budget,
			steering_wheel_type,
			city,
			person_name,
			person_phone,
			submitted, 
			step,
			created_at, 
			updated_at 
		from 
			applications 
		where user_id = ?
	`

	rows, err := repo.Connection.Query(sql, userId)

	if err != nil {
		defer rows.Close()

		return nil, err
	}

	var applications []*entities.Application

	for rows.Next() {
		var appl entities.Application
		err := rows.Scan(
			&appl.ID,
			&appl.UserId,
			&appl.Country,
			&appl.MarkOrConditions,
			&appl.Budget,
			&appl.SteeringWheelType,
			&appl.City,
			&appl.PersonName,
			&appl.PersonPhone,
			&appl.Submitted,
			&appl.Step,
			&appl.CreatedAt,
			&appl.UpdatedAt,
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

func (repo ApplicationRepository) CreateEmpty(userdId int) error {
	sql := `insert into applications(user_id, created_at, updated_at) values (?, datetime('now'), datetime('now'))`

	_, err := repo.Connection.Exec(sql, userdId)

	if err == nil {
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
            submitted = ?,
            step = ?,
            updated_at = datetime('now')
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
		appl.Submitted,
		appl.Step,
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

func NewApplicationRepository(conn *sql.DB) *ApplicationRepository {
	return &ApplicationRepository{Connection: conn}
}
