package models

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	CreatedAt time.Time
}

type NABHistory struct {
	ID        int     `gorm:"primaryKey"`
	NAB       float64 `gorm:"not null"`
	CreatedAt time.Time
}

type UserInvestment struct {
	ID        int     `gorm:"primaryKey"`
	UserID    int     `gorm:"not null"`
	Unit      float64 `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Transaction struct {
	ID        int     `gorm:"primaryKey"`
	UserID    int     `gorm:"not null"`
	Type      string  `gorm:"type:enum('TOPUP','WITHDRAW');not null"`
	Amount    float64 `gorm:"not null"`
	Unit      float64 `gorm:"not null"`
	NABAt     float64 `gorm:"not null"`
	CreatedAt time.Time
}
