package models

import (
	"encoding/json"
)

// Shop ORM Model
type Shop struct {
	Id      int     `gorm:"column:id;NOT NULL;PRIMARY KEY;"`
	Name    string  `gorm:"column:name;NOT NULL"`
	Price   float32 `gorm:"column:price;NOT NULL"`
	Content string  `gorm:"column:content;NOT NULL"`
}

// ShopContent Model
type ShopContent struct {
	Bandwidth    string
	Expire       string
	Class        string
	Class_expire string
	Reset        string
	Reset_exp    string
	Speedlimit   string
	Connector    string
}

// GetContent converts the content field from string to an obj
func (shop *Shop) GetContent() ShopContent {
	var content ShopContent
	json.Unmarshal([]byte(shop.Content), &content)
	return content
}

// TableName returns the table name
func (Shop) TableName() string {
	return "shop"
}
