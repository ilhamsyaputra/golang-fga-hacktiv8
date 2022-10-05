package models

import "time"

type Order struct {
	OrderID      uint      `gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items"`
	OrderedAt    time.Time `json:"orderedAt"`
}
