package model

type Item struct {
	ID          int      `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Author      string   `json:"author"`
	Price       float32  `json:"price"`
	Rating      float32  `json:"rating"`
	RatingTotal float32  `json:"-"`
	RatingCount int      `json:"-"`
	Description string   `json:"description"`
	CategoryID  int      `json:"category_id"`
	UserID      int      `json:"user_id"`
	Category    Category `json:"-" gorm:"foreignKey:CategoryID"`
	User        User     `json:"-" gorm:"foreignKey:UserID"`
}
