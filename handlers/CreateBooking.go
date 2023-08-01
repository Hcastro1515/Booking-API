package handlers

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/hcastro1515/booking-api/models"
)

func (h handler) CreateBooking(w http.ResponseWriter, r *http.Request) {

	var bookingRequest models.BookingRequest

	err := json.NewDecoder(r.Body).Decode(&bookingRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Fatal(err)
	}
	var service models.Service
	h.DB.First(&service, bookingRequest.ServiceId)
	booking := models.Booking{
		Name:        bookingRequest.Name,
		BookingDate: bookingRequest.BookingDate,
		BookingTime: bookingRequest.BookingTime,
		Price:       bookingRequest.Price,
		PhoneNumber: bookingRequest.PhoneNumber,
		Email:       bookingRequest.Email,
		ServiceId:   uint(service.Id),
		Service:     service,
	}

	if result := h.DB.Create(&booking); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created!")
}
