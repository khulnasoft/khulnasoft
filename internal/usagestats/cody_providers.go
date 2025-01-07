package usagestats

import (
	"github.com/khulnasoft/khulnasoft/internal/conf"
	"github.com/khulnasoft/khulnasoft/internal/conf/conftypes"
	"github.com/khulnasoft/khulnasoft/internal/types"
)

func GetCodyProviders() (*types.CodyProviders, error) {
	siteConfig := conf.SiteConfig()
	completionsConfig := conf.GetCompletionsConfig(siteConfig)
	embeddingsConfig := conf.GetEmbeddingsConfig(siteConfig)
	providers := types.CodyProviders{}
	if completionsConfig != nil {
		providers.Completions = &types.CodyCompletionProvider{
			Provider: completionsConfig.Provider,
		}
		if completionsConfig.Provider == conftypes.CompletionsProviderNameKhulnasoft {
			providers.Completions.ChatModel = completionsConfig.ChatModel
			providers.Completions.CompletionModel = completionsConfig.CompletionModel
			providers.Completions.FastChatModel = completionsConfig.FastChatModel
		}
	}
	if embeddingsConfig != nil {
		providers.Embeddings = &types.CodyEmbeddingsProvider{
			Provider: embeddingsConfig.Provider,
		}
		if embeddingsConfig.Provider == conftypes.EmbeddingsProviderNameKhulnasoft {
			providers.Embeddings.Model = embeddingsConfig.Model
		}
	}
	return &providers, nil
}
