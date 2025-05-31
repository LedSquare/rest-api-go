package utils

import(
	"time"
)

type Account struct {
    ID      uint    `gorm:"primaryKey"`
    OwnerID string  `gorm:"unique"`
    Balance float64 `gorm:"default:0"`
}

type Invoice struct {
    ID          uint      `gorm:"primaryKey"`
    ReferenceID string    `gorm:"unique"`
    UserID      uint      `gorm:"index"`
    Amount      float64   `gorm:"not null"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
    Status      string    `gorm:"default:'PENDING'"`
    IsTmp       *bool     `gorm:"default:null"`
}	

type Message struct {
    ID         uint      `gorm:"primaryKey"`
    UserID     uint      `gorm:"index"`
    Model      string    `gorm:"not null"`
    TokenCount int       `gorm:"default:0"`
    Type       string    `gorm:"not null"`
    CreatedAt  time.Time `gorm:"autoCreateTime"`
}

type MidjourneyMessage struct {
    ID        uint      `gorm:"primaryKey"`
    ChatID    uint      `gorm:"index"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Offer struct {
    ID string `gorm:"primaryKey"`
}

type Price struct {
    ID           uint    `gorm:"primaryKey"`
    ServiceID    uint    `gorm:"index"`
    PriceTypeID  uint    `gorm:"index"`
    Amount       float64 `gorm:"not null"`
}

type Service struct {
    ID    uint    `gorm:"primaryKey"`
    Name  string  `gorm:"not null"`
    LLM   string  `gorm:"not null"`
    Prices []Price `gorm:"foreignKey:ServiceID"`
}