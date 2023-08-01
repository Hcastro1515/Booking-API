package models

type Service struct {
	Id          int    `json:"id" gorm:"primary_key;auto_increment;not null"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}
