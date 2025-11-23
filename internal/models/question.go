package models

import "time"

type Question struct {
    ID        uint      `gorm:"primaryKey"`
    Text      string    `gorm:"type:text;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`

    Answers []Answer `gorm:"constraint:OnDelete:CASCADE;"`
}
