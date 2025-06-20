package model

import "time"

type UserLogs struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	Username     string    `db:"username"`
	Name         string    `db:"name"`
	Role         string    `db:"role"`
	Time         time.Time `db:"time"`
	Activity     string    `db:"activity"`
	ActivityType string    `db:"activity_type"`
}
