package handlers

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/slogsdon/atomic/models"
)

var resources = []models.Resource{
	models.Resource{
		Label: "Resource Collection",
		Url:   "http://atomic.logsdon.io/",
	},
	models.Resource{
		Label: "Prompts Collection",
		Url:   "http://atomic.logsdon.io/prompts",
	},
	models.Resource{
		Label: "Retrieve a Prompt",
		Url:   "http://atomic.logsdon.io/prompts/{slug}",
	},
	models.Resource{
		Label: "Prompt's Responses Collection",
		Url:   "http://atomic.logsdon.io/prompts/{slug}/responses",
	},
}

func Resources() martini.Handler {
	return func() string {
		resp, err := json.MarshalIndent(resources, "", "  ")
		if err != nil {
			return ""
		}
		return string(resp)
	}
}
