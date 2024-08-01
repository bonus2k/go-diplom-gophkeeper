package model

type User struct {
	ID         int64        `gorm:"primary_key;auto_increment" json:"id"`
	Username   string       `gorm:"size:255;not null" json:"username"`
	Password   string       `gorm:"size:255;not null" json:"password"`
	Email      string       `gorm:"size:255;not null" json:"email"`
	CreatedAt  int64        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  int64        `gorm:"autoUpdateTime" json:"updated_at"`
	HashData   string       `gorm:"size:255" json:"hash_data,omitempty"`
	SecretData []SecretData `json:"secret_data,omitempty"`
}

type SecretData struct {
	ID     int64  `gorm:"primary_key;auto_increment" json:"id"`
	UserID int64  `gorm:"not null" json:"user_id"`
	Type   string `gorm:"size:255;not null" json:"type"`
	Name   string `gorm:"size:255;not null" json:"name"`
	Secret []byte `gorm:"size:255;not null" json:"secret"`
}
