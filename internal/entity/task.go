package entity

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	CronSpec string `json:"cron_spec"`
	Payload  string `json:"payload"`
	Type     string `json:"type"`
}
