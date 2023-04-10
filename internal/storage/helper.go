package storage

import (
	"reflect"

	"github.com/maze1377/manager-vending-machine/config"
	"github.com/maze1377/manager-vending-machine/internal/storage/entity"
	"github.com/maze1377/manager-vending-machine/pkg/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Migrate() (db *gorm.DB) {
	db, err := GetDataBase(config.Instance.DebugDb)
	if err != nil {
		logrus.WithError(err).Fatal("error while connecting to database")
	}
	logrus.Info("migration started")
	migrate(db, &entity.EventLog{})
	logrus.Info("migration successfully finished")
	return db
}

func migrate(db *gorm.DB, model interface{}) {
	if err := db.AutoMigrate(model); err != nil {
		logrus.WithError(err).WithField("model", reflect.TypeOf(model)).Fatal("error while migrating")
	}
}

func GetDataBase(debug bool) (db *gorm.DB, err error) {
	db, err = sql.GetDatabase(&sql.SqliteConfig{InMemory: true})
	if err != nil {
		logrus.WithError(err).Panic("can not init Db")
	}
	if debug {
		db = db.Debug() // NOTE: shows queries
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(50)
	}
	return db, err
}
