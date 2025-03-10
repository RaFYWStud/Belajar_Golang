package dto

type IntroRequest struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Nama          string `json:"nama"`
	NamaPanggilan string `json:"nama_panggilan"`
	FunFact       string `json:"fun_fact"`
	KeinginanBE   string `json:"keinginan_be"`
	UpdatedAt     string `json:"updated_at"`
	CreatedAt     string `json:"created_at"`
}

type IntroData struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	NamaPanggilan string `json:"nama_panggilan"`
	FunFact       string `json:"fun_fact"`
	KeinginanBE   string `json:"keinginan_be"`
}

type IntroResponse struct {
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
	Data       IntroData `json:"data"`
}

// Error implements error.
func (i *IntroResponse) Error() string {
	panic("unimplemented")
}
