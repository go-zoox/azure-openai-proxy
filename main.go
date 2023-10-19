package main

import (
	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "azure-openai-proxy",
		Usage:   "azure-openai-proxy is a portable chatgpt server",
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
				Name:    "chat-completion-resource-for-gpt35",
				Usage:   "Chat-completion Resource (GPT-3.5)",
				EnvVars: []string{"CHAT_COMPLETION_RESOURCE_FOR_GPT35"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-deployment-for-gpt35",
				Usage:   "Chat-completion Deployment (GPT-3.5)",
				EnvVars: []string{"CHAT_COMPLETION_DEPLOYMENT_FOR_GPT35"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-resource-for-gpt4",
				Usage:   "Chat-completion Resource (GPT-4)",
				EnvVars: []string{"CHAT_COMPLETION_RESOURCE_FOR_GPT4"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-deployment-for-gpt4",
				Usage:   "Chat-completion Deployment (GPT-4)",
				EnvVars: []string{"CHAT_COMPLETION_DEPLOYMENT_FOR_GPT4"},
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
			APIs: APIs{
				ChatCompletions: Models{
					"gpt-3.5": ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt35"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt35"),
					},
					"gpt-3.5-turbo": ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt35"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt35"),
					},
					"gpt-3.5-turbo-16k": ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt35"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt35"),
					},
					"gpt-4": ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt4"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt4"),
					},
				},
				Embeddings: ModelResource{
					Resource:   ctx.String("embeddings-resource"),
					Deployment: ctx.String("embeddings-deployment"),
				},
			},
		})
	})

	app.Run()
}
