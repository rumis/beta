package entity

type ChatRequest struct {
	Prompt string `json:"prompt"`
}

type ChatResponse struct {
	ID string `json:"id"`
}

type ChatResponseChunk struct {
	ID           string `json:"id"`
	Chunk        string `json:"chunk"`
	FinishReason string `json:"finish_reason"`
}

// ChatCompletionRequest chat completion request chunk metadata
type ChatResponseChunkMetadata struct {
	ID        string
	ErrString string
}
