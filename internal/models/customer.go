package models

type Customer struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
