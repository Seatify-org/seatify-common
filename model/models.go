package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"` // или FullName
	Role      string    `json:"role"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Movie struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Duration    int      `json:"duration_minutes"`
	ReleaseDate string   `json:"release_date"` // YYYY-MM-DD
	PosterURL   string   `json:"poster_url"`
	BannerURL   string   `json:"banner_url,omitempty"`
	TrailerURL  string   `json:"trailer_url,omitempty"`
	Genres      []string `json:"genres"`
	Cast        []string `json:"cast,omitempty"`
	Director    string   `json:"director,omitempty"`
	Rating      float64  `json:"rating,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type Cinema struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	Address            string   `json:"address"`
	City               string   `json:"city"`
	Latitude           float64  `json:"latitude,omitempty"`
	Longitude          float64  `json:"longitude,omitempty"`
	Rating             float64  `json:"rating,omitempty"`
	Facilities         []string `json:"facilities,omitempty"`
	InfrastructureTags []string `json:"infrastructure_tags,omitempty"`
	ImageURL           string   `json:"image_url,omitempty"`
	IntegrationLevel   int      `json:"integration_level"` // 1, 2, 3
	PhoneNumber        string   `json:"phone_number,omitempty"`
	TotalHalls         int      `json:"total_halls,omitempty"`
	CreatedAt          time.Time `json:"created_at"`
}

type Hall struct {
	ID           int    `json:"id"`
	CinemaID     int    `json:"cinema_id"`
	Name         string `json:"name"`
	Rows         int    `json:"rows"`
	TotalSeats int   `json:"total_seats,omitempty"`
	SeatsPerRow  int    `json:"seats_per_row"`
	CreatedAt time.Time `json:"created_at"`
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
