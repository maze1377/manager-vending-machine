package dbrepository

import (
	"github.com/maze1377/manager-vending-machine/internal/metrics"
	"github.com/maze1377/manager-vending-machine/internal/storage/entity"
	"gorm.io/gorm"
)

type EventLog struct {
	*Gorm[*entity.EventLog]
}

func NewEventLog(client *gorm.DB, watcher metrics.Communicator) EventLogRepository {
	return &EventLog{
		Gorm: &Gorm[*entity.EventLog]{
			TableSetting: TableSetting{
				ConflictColumn:   entity.EventLogConflict,
				OnConflictUpdate: entity.EventLogOnConflictUpdate,
			},
			Metrics: watcher,
			Client:  client,
		},
	}
}
