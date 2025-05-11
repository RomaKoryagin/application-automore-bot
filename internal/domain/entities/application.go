package entities

import "database/sql"

type Application struct {
	ID                int            `json:"id"`
	UserId            int            `json:"user_id"`
	Country           sql.NullString `json:"country"`
	MarkOrConditions  sql.NullString `json:"mark_or_conditions"`
	Budget            sql.NullString `json:"budget"`
	SteeringWheelType sql.NullString `json:"steering_wheel_type"`
	City              sql.NullString `json:"city"`
	PersonName        sql.NullString `json:"person_name"`
	PersonPhone       sql.NullString `json:"person_phone"`
	Submitted         sql.NullBool   `json:"submitted"`
	Step              int            `json:"step"`
	UpdatedAt         string         `json:"updated_at"`
	CreatedAt         string         `json:"created_at"`
}
