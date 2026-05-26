package model

import "time"

// User представляет пользователя системы
type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password_hash"` // Скрыт в JSON
	Name      string    `json:"name" db:"full_name"`
	FirstName string    `json:"first_name" db:"first_name"` // Для совместимости, если нужно
	LastName  string    `json:"last_name" db:"last_name"`   // Для совместимости, если нужно
	Phone     string    `json:"phone" db:"phone"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Movie представляет фильм
type Movie struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Duration    int       `json:"duration_minutes" db:"duration_minutes"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	PosterURL   string    `json:"poster_url" db:"poster_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Cinema представляет кинотеатр
type Cinema struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	City    string `json:"city" db:"city"`
}

// Hall представляет зал в кинотеатре
type Hall struct {
	ID           int    `json:"id" db:"id"`
	CinemaID     int    `json:"cinema_id" db:"cinema_id"`
	Name         string `json:"name" db:"name"`
	Rows         int    `json:"rows" db:"rows"`
	SeatsPerRow  int    `json:"seats_per_row" db:"seats_per_row"`
}

// Session представляет киносеанс
type Session struct {
	ID            int       `json:"id" db:"id"`
	MovieID       int       `json:"movie_id" db:"movie_id"`
	HallID        int       `json:"hall_id" db:"hall_id"`
	CinemaID      int       `json:"cinema_id" db:"cinema_id"` // Денормализация для удобства
	StartTime     time.Time `json:"start_time" db:"start_time"`
	BasePriceCents int      `json:"base_price_cents" db:"base_price_cents"`
	
	// Дополнительные поля для ответа API (заполняются при получении)
	MovieTitle    string    `json:"movie_title,omitempty" db:"movie_title"`
	CinemaName    string    `json:"cinema_name,omitempty" db:"cinema_name"`
	CinemaAddress string    `json:"cinema_address,omitempty" db:"cinema_address"`
	CinemaCity    string    `json:"cinema_city,omitempty" db:"cinema_city"`
	HallName      string    `json:"hall_name,omitempty" db:"hall_name"`
}

// Booking представляет бронирование
type Booking struct {
	ID             int        `json:"id" db:"id"`
	UserID         int        `json:"user_id" db:"user_id"`
	SessionID      int        `json:"session_id" db:"session_id"`
	TotalAmount    int        `json:"total_amount_cents" db:"total_amount_cents"`
	PaymentID      string     `json:"payment_id,omitempty" db:"payment_id"`
	Status         string     `json:"status" db:"status"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	CancelledAt    *time.Time `json:"cancelled_at,omitempty" db:"cancelled_at"`
}

// Seat представляет место в зале
type Seat struct {
	ID          int    `json:"id" db:"id"`
	HallID      int    `json:"hall_id" db:"hall_id"`
	Row         int    `json:"row" db:"row"`
	SeatNumber  int    `json:"seat_number" db:"seat_number"`
	SeatType    string `json:"seat_type" db:"seat_type"`
	PriceCents  int    `json:"price_cents" db:"price_cents"`
}

// BookingSeat связывает бронирование с местами
type BookingSeat struct {
	ID        int `json:"id" db:"id"`
	BookingID int `json:"booking_id" db:"booking_id"`
	SeatID    int `json:"seat_id" db:"seat_id"`
}
