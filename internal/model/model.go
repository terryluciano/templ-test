package model

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"not null;uniqueIndex"`
	Password  string `gorm:"not null"`
	FirstName string
	LastName  string
}
