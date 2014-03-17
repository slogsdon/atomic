package main

import (
	"github.com/codegangsta/martini"
	"github.com/slogsdon/atomic/db"
	"github.com/slogsdon/atomic/handlers"
)

func main() {
	m := martini.Classic()
	m.Use(addHeaders())
	m.Map(&db.DB)

	m.Get("/", handlers.Resources())
	m.Get("/prompts", handlers.Prompts())
	m.Get("/prompts/:prompt_slug", handlers.ShowPrompt())
	m.Get("/prompts/:prompt_slug/responses", handlers.ShowPromptResponses())
	m.Get("/responses", handlers.Responses())
	m.Get("/responses/:response_id", handlers.ShowResponse())

	m.Run()
}
