package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hcastro1515/booking-api/models"
)

func (h handler) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	var bookings []models.Booking
	w.Header().Set("Content-Type", "application/json")
	h.DB.Preload("Service").Find(&bookings)
	json.NewEncoder(w).Encode(bookings)
}
