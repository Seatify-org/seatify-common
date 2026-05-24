package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Не выводится в JSON
	Name      string    `json:"name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Booking struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	SessionID   int       `json:"session_id"`
	PaymentID   string    `json:"payment_id,omitempty"`
	Status      string    `json:"status"`
	TotalAmount int       `json:"total_amount_cents"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Seat struct {
	ID         int    `json:"id"`
	SessionID  int    `json:"session_id"`
	RowNumber  int    `json:"row_number"`
	SeatNumber int    `json:"seat_number"`
	SeatType   string `json:"seat_type"`
	Status     string `json:"status"`
}

type Session struct {
	ID         int       `json:"id"`
	MovieTitle string    `json:"movie_title"`
	CinemaName string    `json:"cinema_name"`
	HallName   string    `json:"hall_name"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	PriceCents int       `json:"price_cents"`
}
