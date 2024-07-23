package translate

import (
	openai "github.com/zijiren233/openai-translator"
)

const (
    fixedURL   = "https://gateway.ai.cloudflare.com/v1/c7301c245fab3e5e60a72e7bd911a64a/aiproxy/openai"
    fixedModel = "gpt-4o-mini"
)

func OpenaiTranslate(q, source, target, key string) (result string, err error) {
	return openai.Translate(q, target, key,
        openai.WithFrom(source),
        openai.WithUrl(fixedURL),
        openai.WithModel(fixedModel),
    )
}
