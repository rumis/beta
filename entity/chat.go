package entity

type ChatResponse struct {
	ID string `json:"id"`
}

// ChatCompletionRequest chat completion request chunk metadata
type ChatResponseChunkMetadata struct {
	ID        string
	ErrString string
}
