package user

import "time"

type User struct {
	UserID    int
	Username  string
	Password  string
	FullName  string
	CreatedAt time.Time
	UpdateAt  time.Time
}
