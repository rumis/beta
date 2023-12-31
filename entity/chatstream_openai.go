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
	Choices []ChatCompletionChunkChoices `json:"choices"`
	// This fingerprint represents the backend configuration that the model runs with.
	// Can be used in conjunction with the seed request parameter to understand when backend changes have been made that might impact determinism.
	SystemFingerprint string `json:"system_fingerprint"`
}

// ChatCompletionChunkChoices is a chunk choice object
type ChatCompletionChunkChoices struct {
	// The index of the choice in the list of choices.
	Index int `json:"index"`
	// A chat completion message generated by the model
	Delta ChatCompletionChunkDelta `json:"delta"`
	// Log probability information for the choice
	Logprobs ChatCompletionChunkLogprobs `json:"logprobs"`
	// The reason the model stopped generating tokens.
	// This will be stop if the model hit a natural stop point or a provided stop sequence,
	// length if the maximum number of tokens specified in the request was reached,
	// content_filter if content was omitted due to a flag from our content filters,
	// tool_calls if the model called a tool, or function_call (deprecated) if the model called a function.
	FinishReason string `json:"finish_reason"`
}

// ChatCompletionChunkMessage is a chunk message that is sent to the API to stream request a completion
type ChatCompletionChunkDelta struct {
	// The contents of the message.
	Content string `json:"content"`
}

type ChatCompletionChunkLogprobs struct {
	Content []ChatCompletionChunkLogprobsContent `json:"content"`
}

type ChatCompletionChunkLogprobsContent struct {
	Token string `json:"token"`
	// The log probabilities of tokens, selected by the model.
	Logprob float64 `json:"logprob"`
	// The log probabilities of the finish tokens, selected by the model.
	Bytes []byte `json:"bytes"`
	// The log probabilities of the chosen tokens, selected by the model.
	TopLogprobs []ChatCompletionChunkLogprobsContent1 `json:"top_logprobs"`
}

type ChatCompletionChunkLogprobsContent1 struct {
	Token string `json:"token"`
	// The log probabilities of tokens, selected by the model.
	Logprob float64 `json:"logprob"`
	// The log probabilities of the finish tokens, selected by the model.
	Bytes []byte `json:"bytes"`
}
