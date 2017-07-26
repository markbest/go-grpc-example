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

func init() {
	db := utils.DB()
	if !db.HasTable(&Article{}) {
		db.CreateTable(&Article{})
	}
}

func (Article) TableName() string {
	return "articles"
}

func GetArticleList(page int64, limit int64) []Article {
	var articles []Article

	db := utils.DB()
	db.Limit(limit).Offset(limit * (page - 1)).Find(&articles)
	return articles
}

func GetArticleInfo(id int64) Article {
	var article Article

	db := utils.DB()
	db.Where("id = ?", id).First(&article)
	return article
}
