package model

type NoteDto struct {
	Name       string       `json:"name"`
	Type       TypeNote     `json:"type"`
	SecretData []SecretData `json:"secret_data,omitempty"`
}

type UserDto struct {
	Username string `gorm:"size:255;not null" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
	Email    string `gorm:"size:255;not null" json:"email"`
}
