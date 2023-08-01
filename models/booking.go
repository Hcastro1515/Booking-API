package models

type Booking struct {
	Id          int     `json:"id" gorm:"primary_key;auto_increment;not null"`
	Name        string  `json:"name" gorm:"not null"`
	BookingDate string  `json:"booking_date" gorm:"not null"`
	BookingTime string  `json:"booking_time" gorm:"not null"`
	PhoneNumber string  `json:"phone_number" gorm:"not null"`
	Email       string  `json:"email" gorm:"not null"`
	Price       float64 `json:"price" gorm:"not null"`
	ServiceId   uint    `json:"service_id" gorm:"not null"`
	Service     Service `json:"service" gorm:"foreignkey:ServiceId"`
}

type BookingRequest struct {
	Name        string  `json:"name"`
	BookingDate string  `json:"booking_date"`
	BookingTime string  `json:"booking_time"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Price       float64 `json:"price"`
	ServiceId   uint    `json:"service_id" gorm:"not null"`
}
