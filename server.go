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
			r.Post(path, zoox.WrapH(proxy.New(&proxy.Config{
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

					logger.Infof("proxying request [%s] %s -> %s", path, originReq.URL, req.URL.String())

					return nil
				},
				OnResponse: func(res *http.Response, originReq *http.Request) error {
					// // issue: https://github.com/Chanzhaoyu/chatgpt-web/issues/831
					// if res.Header.Get("Content-Type") == "text/event-stream" {
					// 	res.Body = &NewBody{Body: res.Body}
					// }

					return nil
				},
			})))
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

// type NewBody struct {
// 	Body io.ReadCloser
// }

// func (b *NewBody) Close() error {
// 	return b.Body.Close()
// }

// func (b *NewBody) Read(p []byte) (n int, err error) {
// 	n, err = b.Body.Read(p)
// 	if err != nil {
// 		return n, err
// 	}

// 	if n != len(p) {
// 		fmt.Println("read", n, len(p))
// 		copy(p[n:], []byte{'\n'})
// 		return n + 1, nil
// 	}

// 	return n, nil
// }
