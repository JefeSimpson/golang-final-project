package model

type Purchase struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	ItemID  int    `json:"item_id" gorm:"primaryKey"`
	UserID  int    `json:"user_id" gorm:"primaryKey"`
	Address string `json:"address"`
	Status  string `json:"status"`
	Item    Item   `json:"-" gorm:"foreignKey:ItemID"`
	User    User   `json:"-" gorm:"foreignKey:UserID"`
}
