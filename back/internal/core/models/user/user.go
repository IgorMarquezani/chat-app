package user

import "time"

type User struct {
	ID        uint32   `gorm:"primaryKey"`
	Name      UserName `gorm:"not null"`
	Email     Email    `gorm:"unique"`
	Password  Password `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
