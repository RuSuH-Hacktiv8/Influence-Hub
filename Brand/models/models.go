package models

type Brand struct {
	ID       uint   `json:"id,omitempty" gorm:"primaryKey"`
	Email    string `json:"email,omitempty" gorm:"unique;"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}
