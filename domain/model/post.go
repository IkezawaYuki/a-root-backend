package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	CustomerID       int    `gorm:"column:customer_id"`
	InstagramMediaID string `gorm:"column:instagram_media_id"`
	InstagramLink    string `gorm:"column:instagram_link"`
	WordpressMediaID string `gorm:"column:wordpress_media_id"`
	WordpressLink    string `gorm:"column:wordpress_link"`
}
