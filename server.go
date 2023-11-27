package main

import (
	"strings"

	"github.com/go-zoox/azure-openai-proxy/api"
	"github.com/go-zoox/azure-openai-proxy/config"
	"github.com/go-zoox/core-utils/fmt"

	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"
	"github.com/go-zoox/zoox/middleware"
)

func Server(cfg *config.Config) error {
	app := defaults.Application()

	fmt.PrintJSON(cfg)

	if cfg.AuthToken != "" {
		app.Use(middleware.BearerToken(strings.Split(cfg.AuthToken, ",")))
	}

	app.Group(cfg.BasePath, func(r *zoox.RouterGroup) {
		// 1. Chat Completions
		{
			path := "/chat/completions"
			r.Any(path, api.ChatCompletions(path, cfg))
		}

		// 2. Embeddings
		{
			path := "/embeddings"
			r.Any(path, api.Embeddings(path, cfg))
		}

		// 3. Images
		{
			{
				path := "/images/generations"
				r.Any(path, api.ImagesGenerations(path, cfg))
			}
			{
				path := "/images/edits"
				r.Any(path, api.ImagesEdits(path, cfg))
			}
		}
	})

	return app.Run(fmt.Sprintf(":%d", cfg.Port))
}
