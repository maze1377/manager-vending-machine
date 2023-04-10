package sql

import "gorm.io/gorm"

type DbConfig interface {
	Connection() gorm.Dialector
}
