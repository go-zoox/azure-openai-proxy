package main

import (
	"github.com/go-zoox/azure-openai-proxy/config"
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
			// &cli.StringFlag{
			// 	Name:     "api-key",
			// 	Usage:    "OpenAI API Key",
			// 	EnvVars:  []string{"API_KEY"},
			// 	Required: true,
			// },
			// &cli.StringFlag{
			// 	Name:    "api-version",
			// 	Usage:   "OpenAI API Version",
			// 	EnvVars: []string{"API_VERSION"},
			// 	Value:   "2023-03-15-preview",
			// },
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
				Name:    "chat-completion-api-version-for-gpt35",
				Usage:   "Chat-completion API Version (GPT-3.5)",
				EnvVars: []string{"CHAT_COMPLETION_API_VERSION_FOR_GPT35"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-api-key-for-gpt35",
				Usage:   "Chat-completion API Key (GPT-3.5)",
				EnvVars: []string{"CHAT_COMPLETION_API_KEY_FOR_GPT35"},
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
				Name:    "chat-completion-api-version-for-gpt4",
				Usage:   "Chat-completion API Version (GPT-4)",
				EnvVars: []string{"CHAT_COMPLETION_API_VERSION_FOR_GPT4"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-api-key-for-gpt4",
				Usage:   "Chat-completion API Key (GPT-4)",
				EnvVars: []string{"CHAT_COMPLETION_API_KEY_FOR_GPT4"},
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
			&cli.StringFlag{
				Name:    "embeddings-api-version",
				Usage:   "Embeddings API Version",
				EnvVars: []string{"EMBEDDING_API_VERSION"},
			},
			&cli.StringFlag{
				Name:    "embeddings-api-key",
				Usage:   "Embeddings API Key",
				EnvVars: []string{"EMBEDDING_API_KEY"},
			},
			&cli.StringFlag{
				Name:    "images-generations-resource",
				Usage:   "Images Generations Resource",
				EnvVars: []string{"IMAGES_GENERATIONS_RESOURCE"},
			},
			&cli.StringFlag{
				Name:    "images-generations-deployment",
				Usage:   "Images Generations Deployment",
				EnvVars: []string{"IMAGES_GENERATIONS_DEPLOYMENT"},
			},
			&cli.StringFlag{
				Name:    "image-generations-api-version",
				Usage:   "Images Generations API Version",
				EnvVars: []string{"IMAGES_GENERATIONS_API_VERSION"},
			},
			&cli.StringFlag{
				Name:    "image-generations-api-key",
				Usage:   "Images Generations API Key",
				EnvVars: []string{"IMAGES_GENERATIONS_API_KEY"},
			},
			&cli.StringFlag{
				Name:    "images-edits-resource",
				Usage:   "Images Edits Resource",
				EnvVars: []string{"IMAGES_EDITS_RESOURCE"},
			},
			&cli.StringFlag{
				Name:    "images-edits-deployment",
				Usage:   "Images Edits Deployment",
				EnvVars: []string{"IMAGES_EDITS_DEPLOYMENT"},
			},
			&cli.StringFlag{
				Name:    "image-edits-api-version",
				Usage:   "Images Edits API Version",
				EnvVars: []string{"IMAGES_GENERATIONS_API_VERSION"},
			},
			&cli.StringFlag{
				Name:    "image-edits-api-key",
				Usage:   "Images Edits API Key",
				EnvVars: []string{"IMAGES_GENERATIONS_API_KEY"},
			},
		},
	})

	app.Command(func(ctx *cli.Context) (err error) {
		cfg := &config.Config{
			Port:      ctx.Int64("port"),
			BasePath:  ctx.String("base-path"),
			AuthToken: ctx.String("auth-token"),
			APIs: config.APIs{
				ChatCompletions: config.Models{
					"gpt-3.5": config.ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt35"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt35"),
						APIVersion: ctx.String("chat-completion-api-version-for-gpt35"),
						APIKey:     ctx.String("chat-completion-api-key-for-gpt35"),
					},
					"gpt-3.5-turbo": config.ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt35"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt35"),
						APIVersion: ctx.String("chat-completion-api-version-for-gpt35"),
						APIKey:     ctx.String("chat-completion-api-key-for-gpt35"),
					},
					"gpt-3.5-turbo-16k": config.ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt35"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt35"),
						APIVersion: ctx.String("chat-completion-api-version-for-gpt35"),
						APIKey:     ctx.String("chat-completion-api-key-for-gpt35"),
					},
					"gpt-4": config.ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt4"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt4"),
						APIVersion: ctx.String("chat-completion-api-version-for-gpt4"),
						APIKey:     ctx.String("chat-completion-api-key-for-gpt4"),
					},
				},
				Embeddings: config.ModelResource{
					Resource:   ctx.String("embeddings-resource"),
					Deployment: ctx.String("embeddings-deployment"),
					APIVersion: ctx.String("embeddings-api-version"),
					APIKey:     ctx.String("embeddings-api-key"),
				},
				ImagesGenerations: config.ModelResource{
					Resource:   ctx.String("images-generations-resource"),
					Deployment: ctx.String("images-generations-deployment"),
					APIVersion: ctx.String("image-generations-api-version"),
					APIKey:     ctx.String("image-generations-api-key"),
				},
				ImagesEdits: config.ModelResource{
					Resource:   ctx.String("images-edits-resource"),
					Deployment: ctx.String("images-edits-deployment"),
					APIVersion: ctx.String("image-edits-api-version"),
					APIKey:     ctx.String("image-edits-api-key"),
				},
			},
		}

		return Server(cfg)
	})

	app.Run()
}
