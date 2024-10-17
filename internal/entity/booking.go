package entity

type Booking struct {
	ID        int    `json:"id"`
	EventID   int    `json:"event_id"`
	UserID    int    `json:"user_id"`
	Quantity  int    `json:"quantity"`
	Reference string `json:"reference"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
