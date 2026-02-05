package models

import "time"

type AttedanceLogs struct {
	ID          int64
	UserID      int64
	TokenID     int64
	Status      string
	CapturedIp  *string
	ClockInTime time.Time

	User  Users           `gorm:"foreignKey:UserID;references:ID"`
	Token AttedanceTokens `gorm:"foreignKey:TokenID;references:ID"`
}
