package models

import (
	"time"
)

type Hub struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TenantID  uint      `json:"tenant_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
