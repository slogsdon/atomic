package models

import "time"

type Prompt struct {
	Id        int64     `json:"id"`
	Prompt    string    `json:"prompt"`
	Slug      string    `json:"slug"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
