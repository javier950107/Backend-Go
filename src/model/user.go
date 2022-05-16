package model

type Users struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	User      string
	Password  string
	IdAddress string
}
