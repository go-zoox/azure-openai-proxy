package config

type Config struct {
	Port      int64
	BasePath  string
	AuthToken string
	APIs      APIs
}

type APIs struct {
	ChatCompletions   Models
	Embeddings        ModelResource
	ImagesGenerations ModelResource
	ImagesEdits       ModelResource
}

type Models map[ModelName]ModelResource

type ModelName string

type ModelResource struct {
	Resource   string
	Deployment string
	APIVersion string
	APIKey     string
}
