package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hcastro1515/booking-api/models"
)

func (h handler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingId, _ := strconv.Atoi(vars["id"])

	if result := h.DB.Where("id = ?", bookingId).Delete(&models.Booking{}); result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted!")

}
