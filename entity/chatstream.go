package entity

// ChatCompletionChunkResponse chat completion stream chunk response
type ChatCompletionChunkResponse struct {
	// A unique identifier for the chat completion.
	ID string `json:"id"`
	// The object type, which is always chat.completion.
	Object string `json:"object"`
	// The Unix timestamp (in seconds) of when the chat completion was created.
	Created int `json:"created"`
	// The model used for the chat completion.
	Model string `json:"model"`
	// A list of chat completion choices. Can be more than one if n is greater than 1.
	Choices []ChatCompletionChoices `json:"choices"`
	// This fingerprint represents the backend configuration that the model runs with.
	// Can be used in conjunction with the seed request parameter to understand when backend changes have been made that might impact determinism.
	SystemFingerprint string `json:"system_fingerprint"`
}
