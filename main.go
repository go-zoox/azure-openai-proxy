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
				Usage:   "Chat Completion Resource (GPT-3.5)",
				EnvVars: []string{"CHAT_COMPLETION_RESOURCE_FOR_GPT35"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-deployment-for-gpt35",
				Usage:   "Chat Completion Deployment (GPT-3.5)",
				EnvVars: []string{"CHAT_COMPLETION_DEPLOYMENT_FOR_GPT35"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-api-version-for-gpt35",
				Usage:   "Chat Completion API Version (GPT-3.5)",
				EnvVars: []string{"CHAT_COMPLETION_API_VERSION_FOR_GPT35"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-api-key-for-gpt35",
				Usage:   "Chat Completion API Key (GPT-3.5)",
				EnvVars: []string{"CHAT_COMPLETION_API_KEY_FOR_GPT35"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-resource-for-gpt4",
				Usage:   "Chat Completion Resource (GPT-4)",
				EnvVars: []string{"CHAT_COMPLETION_RESOURCE_FOR_GPT4"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-deployment-for-gpt4",
				Usage:   "Chat Completion Deployment (GPT-4)",
				EnvVars: []string{"CHAT_COMPLETION_DEPLOYMENT_FOR_GPT4"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-api-version-for-gpt4",
				Usage:   "Chat Completion API Version (GPT-4)",
				EnvVars: []string{"CHAT_COMPLETION_API_VERSION_FOR_GPT4"},
			},
			&cli.StringFlag{
				Name:    "chat-completion-api-key-for-gpt4",
				Usage:   "Chat Completion API Key (GPT-4)",
				EnvVars: []string{"CHAT_COMPLETION_API_KEY_FOR_GPT4"},
			},
			&cli.StringFlag{
				Name:    "embedding-resource",
				Usage:   "Embedding Resource",
				EnvVars: []string{"EMBEDDING_RESOURCE"},
			},
			&cli.StringFlag{
				Name:    "embedding-deployment",
				Usage:   "Embedding Deployment",
				EnvVars: []string{"EMBEDDING_DEPLOYMENT"},
			},
			&cli.StringFlag{
				Name:    "embedding-api-version",
				Usage:   "Embedding API Version",
				EnvVars: []string{"EMBEDDING_API_VERSION"},
			},
			&cli.StringFlag{
				Name:    "embedding-api-key",
				Usage:   "Embedding API Key",
				EnvVars: []string{"EMBEDDING_API_KEY"},
			},
			&cli.StringFlag{
				Name:    "image-generation-resource",
				Usage:   "Image Generation Resource",
				EnvVars: []string{"IMAGE_GENERATION_RESOURCE"},
			},
			&cli.StringFlag{
				Name:    "image-generation-deployment",
				Usage:   "Image Generation Deployment",
				EnvVars: []string{"IMAGE_GENERATION_DEPLOYMENT"},
			},
			&cli.StringFlag{
				Name:    "image-generation-api-version",
				Usage:   "Image Generation API Version",
				EnvVars: []string{"IMAGE_GENERATION_API_VERSION"},
			},
			&cli.StringFlag{
				Name:    "image-generation-api-key",
				Usage:   "Image Generation API Key",
				EnvVars: []string{"IMAGE_GENERATION_API_KEY"},
			},
			&cli.StringFlag{
				Name:    "image-edit-resource",
				Usage:   "Image Edits Resource",
				EnvVars: []string{"IMAGE_EDIT_RESOURCE"},
			},
			&cli.StringFlag{
				Name:    "image-edit-deployment",
				Usage:   "Image Edits Deployment",
				EnvVars: []string{"IMAGE_EDIT_DEPLOYMENT"},
			},
			&cli.StringFlag{
				Name:    "image-edit-api-version",
				Usage:   "Image Edits API Version",
				EnvVars: []string{"IMAGE_EDIT_API_VERSION"},
			},
			&cli.StringFlag{
				Name:    "image-edit-api-key",
				Usage:   "Image Edits API Key",
				EnvVars: []string{"IMAGE_EDIT_API_KEY"},
			},
		},
	})

	app.Command(func(ctx *cli.Context) (err error) {
		cfg := &config.Config{
			Port:      ctx.Int64("port"),
			BasePath:  ctx.String("base-path"),
			AuthToken: ctx.String("auth-token"),
			APIs: config.APIs{
				ChatCompletion: config.Models{
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
					"gpt-4-turbo": config.ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt4"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt4"),
						APIVersion: ctx.String("chat-completion-api-version-for-gpt4"),
						APIKey:     ctx.String("chat-completion-api-key-for-gpt4"),
					},
					"gpt-4-1106-preview": config.ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt4"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt4"),
						APIVersion: ctx.String("chat-completion-api-version-for-gpt4"),
						APIKey:     ctx.String("chat-completion-api-key-for-gpt4"),
					},
					"gpt-4-32k": config.ModelResource{
						Resource:   ctx.String("chat-completion-resource-for-gpt4"),
						Deployment: ctx.String("chat-completion-deployment-for-gpt4"),
						APIVersion: ctx.String("chat-completion-api-version-for-gpt4"),
						APIKey:     ctx.String("chat-completion-api-key-for-gpt4"),
					},
				},
				Embedding: config.ModelResource{
					Resource:   ctx.String("embedding-resource"),
					Deployment: ctx.String("embedding-deployment"),
					APIVersion: ctx.String("embedding-api-version"),
					APIKey:     ctx.String("embedding-api-key"),
				},
				ImageGeneration: config.ModelResource{
					Resource:   ctx.String("image-generation-resource"),
					Deployment: ctx.String("image-generation-deployment"),
					APIVersion: ctx.String("image-generation-api-version"),
					APIKey:     ctx.String("image-generation-api-key"),
				},
				ImageEdit: config.ModelResource{
					Resource:   ctx.String("image-edit-resource"),
					Deployment: ctx.String("image-edit-deployment"),
					APIVersion: ctx.String("image-edit-api-version"),
					APIKey:     ctx.String("image-edit-api-key"),
				},
			},
		}

		return Server(cfg)
	})

	app.Run()
}
