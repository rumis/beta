package entity

import "github.com/rumis/beta/enum"

// ChatCompletionRequest chat completion request
type ChatCompletionRequest struct {
	// A list of messages comprising the conversation so far.
	Messages []ChatCompletionMessage `json:"messages"`
	// ID of the model to use.
	Model string `json:"model"`

	// Optional Defaults to 0
	// Number between -2.0 and 2.0.
	// Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
	// Defaults to null
	// Modify the likelihood of specified tokens appearing in the completion.
	//Accepts a JSON object that maps tokens (specified by their token ID in the tokenizer) to an associated bias value from -100 to 100.
	// Mathematically, the bias is added to the logits generated by the model prior to sampling.
	// The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection;
	// values like -100 or 100 should result in a ban or exclusive selection of the relevant token.
	LogitBias map[string]int `json:"logit_bias,omitempty"`
	// Optional
	// The maximum number of tokens to generate in the chat completion.
	// The total length of input tokens and generated tokens is limited by the model's context length.
	// Example Python code for counting tokens.
	MaxTokens int `json:"max_tokens,omitempty"`

	// Optional Defaults to 1
	// How many chat completion choices to generate for each input message.
	// Note that you will be charged based on the number of generated tokens across all of the choices.
	//  Keep n as 1 to minimize costs.
	N int `json:"n,omitempty"`

	// Optional Defaults to 0
	// Number between -2.0 and 2.0.
	// Positive values penalize new tokens based on whether they appear in the text so far,
	// increasing the model's likelihood to talk about new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// Optional
	// An object specifying the format that the model must output.
	ResponseFormat *ChatCompletionResponseFormat `json:"response_format,omitempty"`

	// Optional
	// This feature is in Beta. If specified, our system will make a best effort to sample deterministically,
	//  such that repeated requests with the same seed and parameters should return the same result.
	// Determinism is not guaranteed, and you should refer to the system_fingerprint response parameter to monitor changes in the backend.
	Seed int `json:"seed,omitempty"`

	// Optional defaults to null
	// Up to 4 sequences where the API will stop generating further tokens.
	Stop []string `json:"stop,omitempty"`

	// Optional defaults to false
	// If set, partial message deltas will be sent, like in ChatGPT.
	// Tokens will be sent as data-only server-sent events as they become available, with the stream terminated by a data: [DONE] message
	Stream bool `json:"stream,omitempty"`

	// Optional defaults to 1.0
	// What sampling temperature to use, between 0 and 2.
	// Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	// We generally recommend altering this or top_p but not both.
	Temperature float64 `json:"temperature,omitempty"`

	// Optional defaults to 1.0
	// An alternative to sampling with temperature, called nucleus sampling,
	// where the model considers the results of the tokens with top_p probability mass.
	// So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// We generally recommend altering this or temperature but not both.
	TopP float64 `json:"top_p,omitempty"`

	// Optional
	// A list of tools the model may call. Currently, only functions are supported as a tool.
	//  Use this to provide a list of functions the model may generate JSON inputs for
	Tools []ChatCompletionTool `json:"tools,omitempty"`

	// Optional This can be either a string or an ChatCompletionToolChoice object.
	// Controls which (if any) function is called by the model.
	// none means the model will not call a function and instead generates a message.
	// auto means the model can pick between generating a message or calling a function.
	// Specifying a particular function via {"type: "function", "function": {"name": "my_function"}} forces the model to call that function.
	ToolChoice any `json:"tool_choice,omitempty"`
	// Optional
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
	User string `json:"user,omitempty"`
}

// ChatCompletionTool
type ChatCompletionTool struct {
	Type     string                           `json:"type"`
	Function ChatCompletionFunctionDefinition `json:"function,omitempty"`
}

// ChatCompletionFunctionDefinition is a function definition object
type ChatCompletionFunctionDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	// Parameters is an object describing the function.
	// You can pass json.RawMessage to describe the schema,
	// or you can pass in a struct which serializes to the proper JSON schema.
	// The jsonschema package is provided for convenience, but you should
	// consider another specialized library if you require more complex schemas.
	Parameters any `json:"parameters"`
}

// ChatCompletionResponseFormat is a response format object
type ChatCompletionResponseFormat struct {
	Type enum.ChatCompletionResponseFormatType `json:"type,omitempty"`
}

// ChatCompletionToolChoice is a tool choice object
type ChatCompletionToolChoice struct {
	Type     string                     `json:"type"`
	Function ChatCompletionToolFunction `json:"function,omitempty"`
}

// ChatCompletionToolFunction is a tool function object
type ChatCompletionToolFunction struct {
	Name string `json:"name"`
}

// ChatCompletionResponse is a response from the openai
type ChatCompletionResponse struct {
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
	// Usage statistics for the completion request.
	Usage ChatCompletionUsage `json:"usage"`
	// This fingerprint represents the backend configuration that the model runs with.
	// Can be used in conjunction with the seed request parameter to understand when backend changes have been made that might impact determinism.
	SystemFingerprint string `json:"system_fingerprint"`
}

// ChatCompletionMessage is a message that is sent to the API to request a completion
type ChatCompletionMessage struct {
	// The role of the author of this message
	Role string `json:"role"`
	// The contents of the message.
	Content string `json:"content"`
	//
	Name string `json:"name,omitempty"`
	// The tool calls generated by the model, such as function calls.
	ToolCalls []ChatCompletionToolCall `json:"tool_calls,omitempty"`
	//  Tool call that this message is responding to.
	//  if message type is tool message ,this field is required
	ToolCallId string `json:"tool_call_id,omitempty"`
}

// ChatCompletionToolCall is a tool call object
type ChatCompletionToolCall struct {
	// The ID of the tool call.
	ID string `json:"id"`
	// The type of the tool. Currently, only function is supported.
	Type string `json:"type"`
	// The function that the model called.
	Function ChatCompletionFunctionCall `json:"function"`
}

// ChatCompletionFunctionCall is a function call object
type ChatCompletionFunctionCall struct {
	// The name of the function to call.
	Name string `json:"name,omitempty"`
	// The arguments to call the function with, as generated by the model in JSON format.
	// Note that the model does not always generate valid JSON,
	// and may hallucinate parameters not defined by your function schema.
	// Validate the arguments in your code before calling your function.
	Arguments string `json:"arguments,omitempty"`
}

// ChatCompletionChoices is a choice object
type ChatCompletionChoices struct {
	// The index of the choice in the list of choices.
	Index int `json:"index"`
	// A chat completion message generated by the model
	Message ChatCompletionMessage `json:"message"`
	// The reason the model stopped generating tokens.
	// This will be stop if the model hit a natural stop point or a provided stop sequence,
	// length if the maximum number of tokens specified in the request was reached,
	// content_filter if content was omitted due to a flag from our content filters,
	// tool_calls if the model called a tool, or function_call (deprecated) if the model called a function.
	FinishReason string `json:"finish_reason"`
}

// ChatCompletionUsage is a usage object that ursed for to see how many tokens are used by an API call
type ChatCompletionUsage struct {
	// Number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens"`
	// Number of tokens in the generated completion.
	CompletionTokens int `json:"completion_tokens"`
	// Total number of tokens used in the request (prompt + completion).
	TotalTokens int `json:"total_tokens"`
}
