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