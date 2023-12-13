package enum

type FinishReason string

const (
	// API returned complete message, or a message terminated by one of the stop sequences provided via the stop parameter
	FinishReasonStop FinishReason = "stop"
	// Incomplete model output due to max_tokens parameter or token limit
	FinishReasonLength FinishReason = "length"
	//  The model decided to call a function
	FinishReasonFunctionCall FinishReason = "function_call"
	// Omitted content due to a flag from our content filters
	FinishReasonContentFilter FinishReason = "content_filter"
	//  API response still in progress or incomplete
	FinishReasonNull FinishReason = "null"
)

const (
	// GPT432K0613           = "gpt-4-32k-0613"
	// GPT432K0314           = "gpt-4-32k-0314"
	GPT432K = "gpt-4-32k"
	// GPT40613              = "gpt-4-0613"
	// GPT40314              = "gpt-4-0314"
	GPT4TurboPreview  = "gpt-4-1106-preview"
	GPT4VisionPreview = "gpt-4-vision-preview"
	GPT4              = "gpt-4"
	// GPT3Dot5Turbo1106     = "gpt-3.5-turbo-1106"
	// GPT3Dot5Turbo0613     = "gpt-3.5-turbo-0613"
	// GPT3Dot5Turbo0301     = "gpt-3.5-turbo-0301"
	GPT3Dot5Turbo16K = "gpt-3.5-turbo-16k"
	// GPT3Dot5Turbo16K0613  = "gpt-3.5-turbo-16k-0613"
	GPT3Dot5Turbo = "gpt-3.5-turbo"
	// GPT3Dot5TurboInstruct = "gpt-3.5-turbo-instruct"
)

type ChatCompletionResponseFormatType string

const (
	ChatCompletionResponseFormatTypeJSONObject ChatCompletionResponseFormatType = "json_object"
	ChatCompletionResponseFormatTypeText       ChatCompletionResponseFormatType = "text"
)
