# Appointment Booking API

This is an appointment booking API developed in Go. It allows users to create, view, update, and delete bookings for appointments.

## Endpoints
---

The API has the following endpoints:

- `GET /bookings`: Retrieves a list of all bookings.
- `POST /bookings`: Creates a new booking.
- `DELETE /bookings/{id}`: Deletes a booking with the specified ID.
- `PUT /bookings/{id}`: Updates a booking with the specified ID.

## Usage

---

To use the API, send HTTP requests to the appropriate endpoint with the required parameters and data. For example, to create a new booking, send a `POST` request to the `/bookings` endpoint with the booking data in the request body.

## Example

---

Here's an example of how you can use the API to create a new booking:

```
curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "booking_date": "2023-08-01", "booking_time": "14:00", "phone_number": "+1-555-123-4567", "email": "john.doe@example.com", "price": 100}' http://localhost:8080/bookings

```

Copy

This will send a `POST` request to the `/bookings` endpoint with the specified booking data in the request body. The API will create a new booking with this data and return a response indicating that the operation was successful.

I hope this helps you get started with using the appointment booking API! 😊
