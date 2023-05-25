package main

import (
	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "chatgpt-for-chatbot-feishu",
		Usage:   "chatgpt-for-chatbot-feishu is a portable chatgpt server",
		Version: Version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Usage:   "server port",
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Value:   8080,
			},
			&cli.StringFlag{
				Name:    "base-path",
				Usage:   "custom api path, default: /",
				EnvVars: []string{"BASE_PATH"},
				Value:   "/",
			},
			&cli.StringFlag{
				Name:    "auth-token",
				Usage:   "auth token",
				EnvVars: []string{"AUTH_TOKEN"},
			},
			&cli.StringFlag{
				Name:     "api-key",
				Usage:    "OpenAI API Key",
				EnvVars:  []string{"API_KEY"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "api-version",
				Usage:   "OpenAI API Version",
				EnvVars: []string{"API_VERSION"},
				Value:   "2023-03-15-preview",
			},
			&cli.StringFlag{
				Name:    "chat-completion-resource",
				Usage:   "Chat-completion Resource",
				EnvVars: []string{"CHAT_COMPLETION_RESOURCE"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-deployment",
				Usage:   "Chat-completion Deployment",
				EnvVars: []string{"CHAT_COMPLETION_DEPLOYMENT"},
			},
			&cli.StringFlag{
				Name:    "embeddings-resource",
				Usage:   "Embeddings Resource",
				EnvVars: []string{"EMBEDDING_RESOURCE"},
			},
			&cli.StringFlag{
				Name:    "embeddings-deployment",
				Usage:   "Embeddings Deployment",
				EnvVars: []string{"EMBEDDING_DEPLOYMENT"},
			},
		},
	})

	app.Command(func(ctx *cli.Context) (err error) {
		return Server(&Config{
			Port:       ctx.Int64("port"),
			BasePath:   ctx.String("base-path"),
			AuthToken:  ctx.String("auth-token"),
			APIKey:     ctx.String("api-key"),
			APIVersion: ctx.String("api-version"),
			Models: Models{
				ChatCompletions: Model{
					Resource:   ctx.String("chat-completion-resource"),
					Deployment: ctx.String("chat-completion-deployment"),
				},
				Embeddings: Model{
					Resource:   ctx.String("embeddings-resource"),
					Deployment: ctx.String("embeddings-deployment"),
				},
			},
		})
	})

	app.Run()
}
