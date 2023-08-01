package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hcastro1515/booking-api/models"
)

func (h handler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var updatedBooking models.Booking
	json.Unmarshal(body, &updatedBooking)

	var booking models.Booking

	if result := h.DB.First(&booking, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	var service models.Service
	h.DB.First(&service, updatedBooking.ServiceId)

	booking.Name = updatedBooking.Name
	booking.Email = updatedBooking.Email
	booking.PhoneNumber = updatedBooking.PhoneNumber
	booking.Price = updatedBooking.Price
	booking.BookingDate = updatedBooking.BookingDate
	booking.BookingTime = updatedBooking.BookingTime
	booking.ServiceId = updatedBooking.ServiceId
	booking.Service = service

	h.DB.Save(&booking)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}
