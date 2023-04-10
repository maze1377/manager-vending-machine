package entity

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
)

var (
	EventLogConflict         = "id"
	EventLogOnConflictUpdate = []string{
		"success", "message",
		"updated_at", "deleted_at",
	}
)

type EventLog struct {
	gorm.Model
	Status  Event  `gorm:"index" json:"status"`
	Message string `json:"message"`
	Success bool   `gorm:"index" json:"success"`
}

type Event string

func (p *Event) Scan(value interface{}) error {
	val := fmt.Sprint(value)
	switch val {
	case EventPayment.String():
		*p = EventPayment
	case EventDispensed.String():
		*p = EventDispensed
	default:
		*p = EventUndefined
	}
	return nil
}

func (p Event) Value() (driver.Value, error) {
	return string(p), nil
}

func (p Event) String() string {
	return string(p)
}

const (
	EventPayment   Event = "payment"
	EventDispensed Event = "dispensed"
	EventUndefined Event = "undefined"
)
