package models

import "github.com/robfig/cron/v3"

type Chat struct {
	ID               int64
	Title            string
	NotificationCron string
	EntryId          cron.EntryID
	OnlyMyChat       bool
}
