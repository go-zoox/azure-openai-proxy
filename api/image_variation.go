package api

import (
	"fmt"
	"net/http"

	"github.com/go-zoox/azure-openai-proxy/config"
	"github.com/go-zoox/headers"
	"github.com/go-zoox/logger"
	"github.com/go-zoox/proxy"
	"github.com/go-zoox/zoox"
)

func ImagesVariations(path string, cfg *config.Config) zoox.HandlerFunc {
	return zoox.WrapH(proxy.New(&proxy.Config{
		OnRequest: func(req, originReq *http.Request) error {
			req.URL.Scheme = "https"
			req.URL.Host = fmt.Sprintf("%s.openai.azure.com", cfg.APIs.ImageVariation.Resource)
			req.URL.Path = fmt.Sprintf("/openai/deployments/%s%s", cfg.APIs.ImageVariation.Deployment, path)
			req.Host = req.URL.Host

			originQuery := req.URL.Query()
			originQuery.Set("api-version", cfg.APIs.ImageVariation.APIVersion)
			req.URL.RawQuery = originQuery.Encode()

			req.Header.Del(headers.Authorization)
			// req.Header.Set(headers.Host, req.URL.Host)
			req.Header.Set(headers.Origin, fmt.Sprintf("%s://%s", req.URL.Scheme, req.URL.Host))
			req.Header.Set("api-key", cfg.APIs.ImageVariation.APIKey)

			logger.Infof("[image_variation][proxy] %s -> %s", originReq.URL, req.URL.String())

			return nil
		},
	}))
}
