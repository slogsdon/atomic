package handlers

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/slogsdon/atomic/db"
	"github.com/slogsdon/atomic/models"
)

func Responses() martini.Handler {
	return func() string {
		var responses []models.Response
		db.DB.Order("id").Find(&responses)

		resp, err := json.MarshalIndent(responses, "", "  ")
		if err != nil {
			return ""
		}
		return string(resp)
	}
}

func ShowResponse() martini.Handler {
	return func(params martini.Params) string {
		var response models.Response
		db.DB.Find(&response, params["response_id"])

		resp, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return ""
		}
		return string(resp)
	}
}

func ShowPromptResponses() martini.Handler {
	return func(params martini.Params) string {
		var prompt models.Prompt
		var responses []models.Response
		db.DB.Where(&models.Prompt{Slug: params["prompt_slug"]}).First(&prompt)
		db.DB.Model(&prompt).Order("id").Related(&responses)

		resp, err := json.MarshalIndent(responses, "", "  ")
		if err != nil {
			return ""
		}
		return string(resp)
	}
}
