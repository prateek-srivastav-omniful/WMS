package models

import "time"

// Hub represents the hub model
type Inventory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Sku_id    uint      `json:"sku_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
