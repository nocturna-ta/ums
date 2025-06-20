package response

import "time"

type UserLogResponse struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Username     string    `json:"username"`
	Name         string    `json:"name"`
	Role         string    `json:"role"`
	Time         time.Time `json:"time"`
	Activity     string    `json:"activity"`
	ActivityType string    `json:"activity_type"`
}
