package task

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uuid.UUID      `gorm:primaryKey" json:id`
	Name        string         `json:name`
	Description string         `json:description`
	UpdatedAt   time.Time      `json:updated_at`
	CreatedAt   time.Time      `json:created_at`
	//DeletedAt   gorm.DeletedAt `json:deleted_at gorm:"index"`

}

func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	task.ID = uuid.New()
	return
}
