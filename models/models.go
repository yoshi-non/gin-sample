package models

import "time"

type Memo struct {
	ID         int       `json:"id" gorm:"index,primary_key"`
	Title      string    `json:"title"`
	Message    string    `json:"message"`
	Updatetime time.Time `json:"updatetime"`
}
