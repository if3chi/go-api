package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Uuid  string
	Email *string
}
