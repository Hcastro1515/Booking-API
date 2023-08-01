package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hcastro1515/booking-api/db"
	"github.com/hcastro1515/booking-api/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/bookings", h.GetAllBookings).Methods(http.MethodGet)
	router.HandleFunc("/bookings", h.CreateBooking).Methods(http.MethodPost)
	router.HandleFunc("/bookings/{id}", h.DeleteBooking).Methods(http.MethodDelete)
	router.HandleFunc("/bookings/{id}", h.UpdateBooking).Methods(http.MethodPut)
	log.Println("API is running! on port :4000")
	http.ListenAndServe(":4000", router)
}
