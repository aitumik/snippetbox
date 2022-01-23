package models

import (
	"errors"
	"time"
	_ "time"
)

var ErrNoRecord = errors.New("error: no record found")

type Snippet struct {
	ID      int       `json:"id" gorm:"Column:id"`
	Title   string    `json:"title" gorm:"Column:title"`
	Content string    `json:"content" gorm:"Column:content"`
	Created time.Time `json:"created_at" gorm:"Column:created"`
	Expires time.Time `json:"expires_at" gorm:"Column:expires"`
}
