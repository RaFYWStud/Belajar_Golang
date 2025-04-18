package dto

type AccountRequest struct {
	ID        int    `json:"id gorm:"primaryKey"`
	Email     string    `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AccountData struct {
	ID        int     `json:"id gorm:"primaryKey"`
	Email     string   `json:"email"`
	Password  string  `json:"password,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type AccountResponse struct {
	StatusCode int		   `json:"status_code"`
	Massage	   string      `json:"message"`
	Data       AccountData `json:"data,omitempty"`
}

func (i *AccountResponse) Error() string {
	panic("unimplemented")
}