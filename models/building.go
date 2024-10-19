package models

import "gorm.io/gorm"

type Building struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"not null"`
    City      string `gorm:"not null"`
    YearBuilt int    `gorm:"not null"`
    Floors    int    `gorm:"not null"`
}