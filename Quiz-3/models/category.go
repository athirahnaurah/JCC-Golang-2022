package models

import "time"

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Created_at  time.Time   `json:"created_at"`
	Update_at time.Time `json:"update_at"`
}
