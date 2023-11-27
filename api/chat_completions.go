package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-zoox/azure-openai-proxy/config"
	"github.com/go-zoox/headers"
	"github.com/go-zoox/proxy"
	"github.com/go-zoox/zoox"
)

func ChatCompletions(path string, cfg *config.Config) zoox.HandlerFunc {
	return func(ctx *zoox.Context) {
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

		model, ok := cfg.APIs.ChatCompletions[config.ModelName(modelName)]
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
				originQuery.Set("api-version", model.APIVersion)
				req.URL.RawQuery = originQuery.Encode()

				req.Header.Del(headers.Authorization)
				req.Header.Set(headers.Host, req.URL.Host)
				req.Header.Set(headers.Origin, fmt.Sprintf("%s://%s", req.URL.Scheme, req.URL.Host))
				req.Header.Set("api-key", model.APIKey)

				ctx.Logger.Infof("[chat/completions][proxy] %s -> %s", originReq.URL, req.URL.String())

				return nil
			},
		}))(ctx)

		// issue: https://github.com/Chanzhaoyu/chatgpt-web/issues/831
		if ctx.Writer.Header().Get(headers.ContentType) == "text/event-stream" {
			ctx.Write([]byte{'\n'})
		}
	}
}
