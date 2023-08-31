package awesomeProject1

import "time"

type UserUpdate struct {
	UserSegments []string `json:"user_segments" db:"user_segments"`
}

type User struct {
	UserId       int      `json:"id" db:"id"`
	UserSegments []string `json:"user_segments" db:"user_segments"`
}
type UpdateUserSegments struct {
	UpdateId           int       `json:"id" db:"id"`
	UpdateUserSegments []string  `json:"update_user_segments" db:"update_user_segments""`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}
