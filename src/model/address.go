package model

type Address struct {
	Id      int `gorm:"primaryKey;autoIncrement"`
	Address string
	City    string
	Country string
}
