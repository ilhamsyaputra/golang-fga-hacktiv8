package models

import "time"

type Order struct {
	OrderID      uint      `gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:ItemID" json:"items"`
}
