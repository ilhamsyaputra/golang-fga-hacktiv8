package models

type Item struct {
	ItemID      uint   `gorm:"primaryKey"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `gorm:"foreignKey:OrderID"`
}
