package model

import (
	"grpc/utils"
	"time"
)

type Category struct {
	Id        int64     `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"size:255" json:"title"`
	ParentId  int64     `json:"parent_id,omitempty"`
	Sort      int64     `json:"sort,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}

func init() {
	db := utils.DB()
	if !db.HasTable(&Category{}) {
		db.CreateTable(&Category{})
	}
}

func GetCategoryList(page int64, limit int64) []Category {
	var categories []Category

	db := utils.DB()
	db.Limit(limit).Offset(limit * (page - 1)).Find(&categories)
	return categories
}

func GetCategoryInfo(id int64) Category {
	var category Category

	db := utils.DB()
	db.Where("id = ?", id).First(&category)
	return category
}

