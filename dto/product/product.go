package productdto

type CreateProduct struct {
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
	Desc  string `json:"desc" form:"desc" gorm:"type: varchar(255)"`
	Stock int    `json:"stock" form:"stock" gorm:"type: int"`
	Image string `json:"image" form:"id" gorm:"type: varchar(255)"`
}

type UpdateProduct struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price"`
	Desc  string `json:"desc" from:desc"`
	Stock int    `json:"stock" from:stock"`
	Image string `json:"image" form:"id"`
}

type ProductResponse struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price"`
	Desc  string `json:"desc" from:desc"`
	Stock int    `json:"stock" from:stock"`
	Image string `json:"image" form:"image"`
}
