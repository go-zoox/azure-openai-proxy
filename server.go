package main

import (
	"encoding/json"
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
	APIs       APIs
}

type APIs struct {
	ChatCompletions Models
	Embeddings      ModelResource
}

type Models map[ModelName]ModelResource

type ModelName string

type ModelResource struct {
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
				body, err := ctx.CloneBody()
				if err != nil {
					ctx.Logger.Errorf("[chat/completions] failed to clone body: %s", err)
					ctx.Error(500, fmt.Sprintf("failed to clone body: %s", err))
					return
				}

				// data := &openaiclient.CreateChatCompletionRequest{}
				data := map[string]any{}
				if err := json.NewDecoder(body).Decode(&data); err != nil {
					ctx.Logger.Errorf("[chat/completions] failed to parse body: %s", err)
					ctx.Error(500, fmt.Sprintf("failed to parse body: %s", err))
					return
				}
				ctx.Logger.Debugf("[chat/completions] raw: %s", toJSON(data))

				modelNameX, ok := data["model"]
				if !ok {
					ctx.Logger.Errorf("[chat/completions] missing model(data: %s)", toJSON(data))
					ctx.Error(500, "missing model")
					return
				}

				modelName, ok := modelNameX.(string)
				if !ok {
					ctx.Logger.Errorf("[chat/completions] modelName is not string (data: %s)", toJSON(data))
					ctx.Error(500, "modelName is not string")
					return
				}
				ctx.Logger.Infof("[chat/completions] model: %s", modelName)

				model, ok := cfg.APIs.ChatCompletions[ModelName(modelName)]
				if !ok {
					ctx.Logger.Errorf("[chat/completions] unsupport model: %s", modelName)
					ctx.Error(500, fmt.Sprintf("unsupport model: %s", modelName))
					return
				}

				zoox.WrapH(proxy.New(&proxy.Config{
					OnRequest: func(req, originReq *http.Request) error {
						req.URL.Scheme = "https"
						req.URL.Host = fmt.Sprintf("%s.openai.azure.com", model.Resource)
						req.URL.Path = fmt.Sprintf("/openai/deployments/%s%s", model.Deployment, path)
						req.Host = req.URL.Host

						originQuery := req.URL.Query()
						originQuery.Set("api-version", cfg.APIVersion)
						req.URL.RawQuery = originQuery.Encode()

						req.Header.Del(headers.Authorization)
						req.Header.Set(headers.Host, req.URL.Host)
						req.Header.Set(headers.Origin, fmt.Sprintf("%s://%s", req.URL.Scheme, req.URL.Host))
						req.Header.Set("api-key", cfg.APIKey)

						ctx.Logger.Infof("[chat/completions][proxy] %s -> %s", originReq.URL, req.URL.String())

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
					req.URL.Host = fmt.Sprintf("%s.openai.azure.com", cfg.APIs.Embeddings.Resource)
					req.URL.Path = fmt.Sprintf("/openai/deployments/%s%s", cfg.APIs.Embeddings.Deployment, path)
					req.Host = req.URL.Host

					originQuery := req.URL.Query()
					originQuery.Set("api-version", cfg.APIVersion)
					req.URL.RawQuery = originQuery.Encode()

					req.Header.Del(headers.Authorization)
					// req.Header.Set(headers.Host, req.URL.Host)
					req.Header.Set(headers.Origin, fmt.Sprintf("%s://%s", req.URL.Scheme, req.URL.Host))
					req.Header.Set("api-key", cfg.APIKey)

					logger.Infof("[embeddings][proxy] %s -> %s", originReq.URL, req.URL.String())

					return nil
				},
			})))
		}
	})

	return app.Run(fmt.Sprintf(":%d", cfg.Port))
}

// @TODO
func toJSON(v any) string {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("failed to toJSON: %#v", v)
	}

	return string(j)
}
