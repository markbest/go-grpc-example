package model

import (
	"grpc/utils"
	"time"
)

type Article struct {
	Id        int64      `gorm:"primary_key" json:"id"`
	Title     string     `gorm:"size:255" json:"title"`
	Slug      string     `gorm:"size:255" json:"slug,omitempty"`
	Summary   string     `gorm:"size:255" json:"summary,omitempty"`
	Body      string     `gorm:"type:text" json:"body,omitempty"`
	Image     string     `gorm:"size:255" json:"image"`
	Views     int64      `gorm:"size:255" json:"views"`
	UserId    int64      `json:"userId"`
	CatId     int64      `json:"catId"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

func (Article) TableName() string {
	return "articles"
}

// e.GET("/article/:id", GetArticleInfo)
func GetArticleInfo(id int64) Article {
	var article Article

	db := utils.DB()
	db.Where("id = ?", id).First(&article)
	return article
}
