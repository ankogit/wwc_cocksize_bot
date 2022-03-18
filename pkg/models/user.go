package models

import "time"

type UserData struct {
	ID       int64
	CockSize int
	Time     time.Time
}
