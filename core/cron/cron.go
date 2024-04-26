package cron

import (
	"github.com/robfig/cron/v3"
)

var (
	// Cron is a cron instance
	Cron *cron.Cron
)

func InitCron() {
	Cron = cron.New(cron.WithSeconds())
}
