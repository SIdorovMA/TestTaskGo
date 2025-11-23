package models

import "time"

type Answer struct {
    ID         uint      `gorm:"primaryKey"`
    QuestionID uint      `gorm:"not null"`
    UserID     string    `gorm:"type:uuid;not null"`
    Text       string    `gorm:"type:text;not null"`
    CreatedAt  time.Time `gorm:"autoCreateTime"`
}
