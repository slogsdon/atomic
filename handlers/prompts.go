package handlers

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/slogsdon/atomic/db"
	"github.com/slogsdon/atomic/models"
)

func Prompts() martini.Handler {
	return func() string {
		var prompts []models.Prompt
		db.DB.Find(&prompts)

		resp, err := json.MarshalIndent(prompts, "", "  ")
		if err != nil {
			return ""
		}
		return string(resp)
	}
}

func ShowPrompt() martini.Handler {
	return func(params martini.Params) string {
		var prompt models.Prompt
		db.DB.Where(&models.Prompt{Slug: params["prompt_slug"]}).First(&prompt)

		resp, err := json.MarshalIndent(prompt, "", "  ")
		if err != nil {
			return ""
		}
		return string(resp)
	}
}
