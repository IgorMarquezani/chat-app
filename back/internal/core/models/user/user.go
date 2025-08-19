package user

import "time"

type User struct {
  ID        uint32   `gorm:"primaryKey" json:"id"`
  Name      UserName `gorm:"not null" json:"name"`
  Email     Email    `gorm:"unique" json:"email"`
  Password  Password `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
