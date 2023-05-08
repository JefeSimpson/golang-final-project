package model

type Comment struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Text   string `json:"text"`
	ItemID int    `json:"item_id"`
	UserID int    `json:"user_id"`
	Item   Item   `json:"-" gorm:"foreignKey:ItemID"`
	User   User   `json:"-" gorm:"foreignKey:UserID"`
}
