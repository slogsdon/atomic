package models

import (
	"time"
)

type Response struct {
	Id        int64     `json:"id"`
	Response  string    `json:"response"`
	Url       string    `json:"url"`
	Prompt    Prompt    `json:"-"`
	PromptId  int64     `json:"prompt_id"`
	CreatedAt time.Time `json:"created_at"`
}
