package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"` // или FullName
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    int       `json:"duration_minutes"`
	ReleaseDate time.Time `json:"release_date"`
	PosterURL   string    `json:"poster_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type Cinema struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
}

type Hall struct {
	ID           int    `json:"id"`
	CinemaID     int    `json:"cinema_id"`
	Name         string `json:"name"`
	Rows         int    `json:"rows"`
	SeatsPerRow  int    `json:"seats_per_row"`
}

type Session struct {
	ID             int       `json:"id"`
	MovieID        int       `json:"movie_id"`
	HallID         int       `json:"hall_id"`
	CinemaID       int       `json:"cinema_id"`
	CinemaAddress  string    `json:"cinema_address"`
	CinemaCity     string    `json:"cinema_city"`
	HallName       string    `json:"hall_name"`       // <--- ДОБАВИТЬ
	CinemaName     string    `json:"cinema_name"`     // <--- ДОБАВИТЬ
	StartTime      time.Time `json:"start_time"`
	BasePriceCents int       `json:"base_price_cents"`
	CreatedAt      time.Time `json:"created_at"`
}

type Seat struct {
	ID          int    `json:"id"`
	HallID      int    `json:"hall_id"`
	Row         int    `json:"row"`
	SeatNumber  int    `json:"seat_number"`
	SeatType    string `json:"seat_type"`
	PriceCents  int    `json:"price_cents"`
}

type Booking struct {
	ID             int        `json:"id"`
	UserID         int        `json:"user_id"`
	SessionID      int        `json:"session_id"`
	TotalAmountCents int      `json:"total_amount_cents"` // Проверьте имя поля
	PaymentID      string     `json:"payment_id,omitempty"`
	Status         string     `json:"status"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	CancelledAt    *time.Time `json:"cancelled_at,omitempty"`
}
