package models

import "time"

type UserData struct {
	ID        int64
	Username  string
	FirstName string
	LastName  string
	CockSize  int
	Time      time.Time
}
