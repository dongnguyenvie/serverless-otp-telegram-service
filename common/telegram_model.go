package common

import "time"

type TelegramUser struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	UserID      int
	Time        time.Time
	Otp         string
}
