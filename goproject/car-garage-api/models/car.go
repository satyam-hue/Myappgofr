package models

type Car struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Model  string `json:"model"`
	Status string `json:"status"`
}
