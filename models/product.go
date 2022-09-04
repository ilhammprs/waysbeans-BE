package models

import "time"

type Product struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type: int"`
	Desc      string    `json:"desc" gorm:"type: varchar(255)"`
	Stock     int       `json:"stock" gorm:"type: int"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	UserID    int       `json:"-" form:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProductTransaction struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (ProductTransaction) TableName() string {
	return "products"
}
