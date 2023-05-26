package main

import (
	"net/http"
	"strings"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/headers"
	"github.com/go-zoox/logger"

	"github.com/go-zoox/proxy"
	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"
	"github.com/go-zoox/zoox/middleware"
)

type Config struct {
	Port       int64
	BasePath   string
	AuthToken  string
	APIKey     string
	APIVersion string
	Models     Models
}

type Models struct {
	ChatCompletions Model
	Embeddings      Model
}

type Model struct {
	Resource   string
	Deployment string
}

func Server(cfg *Config) error {
	app := defaults.Application()

	fmt.PrintJSON(cfg)

	if cfg.AuthToken != "" {
		app.Use(middleware.BearerToken(strings.Split(cfg.AuthToken, ",")))
	}

	app.Group(cfg.BasePath, func(r *zoox.RouterGroup) {
		// 1. Chat Completions
		{
			path := "/chat/completions"
			r.Post(path, func(ctx *zoox.Context) {
				zoox.WrapH(proxy.New(&proxy.Config{
					OnRequest: func(req, originReq *http.Request) error {
						req.URL.Scheme = "https"
						req.URL.Host = fmt.Sprintf("%s.openai.azure.com", cfg.Models.ChatCompletions.Resource)
						req.URL.Path = fmt.Sprintf("/openai/deployments/%s%s", cfg.Models.ChatCompletions.Deployment, path)
						req.Host = req.URL.Host

						originQuery := req.URL.Query()
						originQuery.Set("api-version", cfg.APIVersion)
						req.URL.RawQuery = originQuery.Encode()

						req.Header.Del(headers.Authorization)
						req.Header.Set(headers.Host, req.URL.Host)
						req.Header.Set(headers.Origin, fmt.Sprintf("%s://%s", req.URL.Scheme, req.URL.Host))
						req.Header.Set("api-key", cfg.APIKey)

						logger.Infof("[proxy] %s -> %s", originReq.URL, req.URL.String())

						return nil
					},
				}))(ctx)

				// issue: https://github.com/Chanzhaoyu/chatgpt-web/issues/831
				if ctx.Writer.Header().Get(headers.ContentType) == "text/event-stream" {
					ctx.Write([]byte{'\n'})
				}
			})
		}

		// 2. Embeddings
		{
			path := "/embeddings"
			r.Post(path, zoox.WrapH(proxy.New(&proxy.Config{
				OnRequest: func(req, originReq *http.Request) error {
					req.URL.Scheme = "https"
					req.URL.Host = fmt.Sprintf("%s.openai.azure.com", cfg.Models.Embeddings.Resource)
					req.URL.Path = fmt.Sprintf("/openai/deployments/%s%s", cfg.Models.Embeddings.Deployment, path)
					req.Host = req.URL.Host

					originQuery := req.URL.Query()
					originQuery.Set("api-version", cfg.APIVersion)
					req.URL.RawQuery = originQuery.Encode()

					req.Header.Del(headers.Authorization)
					// req.Header.Set(headers.Host, req.URL.Host)
					req.Header.Set(headers.Origin, fmt.Sprintf("%s://%s", req.URL.Scheme, req.URL.Host))
					req.Header.Set("api-key", cfg.APIKey)

					logger.Infof("[proxy] %s -> %s", originReq.URL, req.URL.String())

					return nil
				},
			})))
		}
	})

	return app.Run(fmt.Sprintf(":%d", cfg.Port))
}
