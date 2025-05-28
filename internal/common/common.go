package common

import (
	"time"
)


type User struct {
	Email             string
	Password_hash     string
	Subscribtion_id   int
	Registration_date time.Time
}

type Request struct {
    ID     int
    UserID string
    Time   time.Time
    Input  string
}