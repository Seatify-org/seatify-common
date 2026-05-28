package model

import "time"

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	PhoneNumber  string    `json:"phone_number,omitempty"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RefreshToken struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type AuditLog struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Action    string    `json:"action"`
	IP        string    `json:"ip,omitempty"`
	UserAgent string    `json:"user_agent,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Duration    int       `json:"duration_minutes"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	PosterURL   string    `json:"poster_url,omitempty"`
	BannerURL   string    `json:"banner_url,omitempty"`
	TrailerURL  string    `json:"trailer_url,omitempty"`
	Rating      float64   `json:"rating,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type Cinema struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Address          string    `json:"address,omitempty"`
	City             string    `json:"city,omitempty"`
	Latitude         float64   `json:"latitude,omitempty"`
	Longitude        float64   `json:"longitude,omitempty"`
	Rating           float64   `json:"rating,omitempty"`
	ImageURL         string    `json:"image_url,omitempty"`
	IntegrationLevel int       `json:"integration_level"`
	PhoneNumber      string    `json:"phone_number,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
}

type Hall struct {
	ID          int       `json:"id"`
	CinemaID    int       `json:"cinema_id"`
	Name        string    `json:"name"`
	Rows        int       `json:"rows_count"`
	SeatsPerRow int       `json:"seats_per_row"`
	HallType    string    `json:"hall_type"`
	TotalSeats  int       `json:"total_seats,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type Session struct {
	ID             int       `json:"id"`
	MovieID        int       `json:"movie_id"`
	MovieTitle     string    `json:"movie_title,omitempty"`
	CinemaID       int       `json:"cinema_id"`
	CinemaName     string    `json:"cinema_name,omitempty"`
	CinemaAddress  string    `json:"cinema_address,omitempty"`
	CinemaCity     string    `json:"cinema_city,omitempty"`
	HallID         int       `json:"hall_id"`
	HallName       string    `json:"hall_name,omitempty"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time,omitempty"`
	BasePriceCents int       `json:"base_price_cents"`
	Status         string    `json:"status,omitempty"`
	AvailableSeats int       `json:"available_seats,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Booking struct {
	ID               int        `json:"id"`
	UserID           int        `json:"user_id"`
	SessionID        int        `json:"session_id"`
	TotalAmountCents int        `json:"total_amount_cents"`
	PaymentID        string     `json:"payment_id,omitempty"`
	Status           string     `json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	CancelledAt      *time.Time `json:"cancelled_at,omitempty"`
}
