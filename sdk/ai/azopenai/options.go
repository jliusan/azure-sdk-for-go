//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenai

// AddUploadPartOptions contains the optional parameters for the Client.AddUploadPart method.
type AddUploadPartOptions struct {
	// placeholder for future optional parameters
}

// CancelBatchOptions contains the optional parameters for the Client.CancelBatch method.
type CancelBatchOptions struct {
	// placeholder for future optional parameters
}

// CancelUploadOptions contains the optional parameters for the Client.CancelUpload method.
type CancelUploadOptions struct {
	// placeholder for future optional parameters
}

// CompleteUploadOptions contains the optional parameters for the Client.CompleteUpload method.
type CompleteUploadOptions struct {
	// placeholder for future optional parameters
}

// CreateBatchOptions contains the optional parameters for the Client.CreateBatch method.
type CreateBatchOptions struct {
	// placeholder for future optional parameters
}

// CreateUploadOptions contains the optional parameters for the Client.CreateUpload method.
type CreateUploadOptions struct {
	// placeholder for future optional parameters
}

// DeleteFileOptions contains the optional parameters for the Client.DeleteFile method.
type DeleteFileOptions struct {
	// placeholder for future optional parameters
}

// GenerateSpeechFromTextOptions contains the optional parameters for the Client.GenerateSpeechFromText method.
type GenerateSpeechFromTextOptions struct {
	// placeholder for future optional parameters
}

// getAudioTranscriptionInternalOptions contains the optional parameters for the Client.getAudioTranscriptionInternal
// method.
type getAudioTranscriptionInternalOptions struct {
	// The optional filename or descriptive identifier to associate with with the audio data.
	Filename *string

	// The primary spoken language of the audio data to be transcribed, supplied as a two-letter ISO-639-1 language code such
	// as 'en' or 'fr'. Providing this known input language is optional but may improve
	// the accuracy and/or latency of transcription.
	Language *string

	// The model to use for this transcription request.
	DeploymentName *string

	// An optional hint to guide the model's style or continue from a prior audio segment. The written language of the prompt
	// should match the primary spoken language of the audio data.
	Prompt *string

	// The requested format of the transcription response data, which will influence the content and detail of the result.
	ResponseFormat *AudioTranscriptionFormat

	// The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values
	// like 0.2 will make it more focused and deterministic. If set to 0, the model will
	// use log probability to automatically increase the temperature until certain thresholds are hit.
	Temperature *float32

	// The timestamp granularities to populate for this transcription.response_format must be set verbose_json to use timestamp
	// granularities. Either or both of these options are supported: word, or segment.
	// Note: There is no additional latency for segment timestamps, but generating word timestamps incurs additional latency.
	TimestampGranularities []AudioTranscriptionTimestampGranularity
}

// getAudioTranslationInternalOptions contains the optional parameters for the Client.getAudioTranslationInternal method.
type getAudioTranslationInternalOptions struct {
	// The optional filename or descriptive identifier to associate with with the audio data.
	Filename *string

	// The model to use for this translation request.
	DeploymentName *string

	// An optional hint to guide the model's style or continue from a prior audio segment. The written language of the prompt
	// should match the primary spoken language of the audio data.
	Prompt *string

	// The requested format of the translation response data, which will influence the content and detail of the result.
	ResponseFormat *AudioTranslationFormat

	// The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values
	// like 0.2 will make it more focused and deterministic. If set to 0, the model will
	// use log probability to automatically increase the temperature until certain thresholds are hit.
	Temperature *float32
}

// GetBatchOptions contains the optional parameters for the Client.GetBatch method.
type GetBatchOptions struct {
	// placeholder for future optional parameters
}

// GetChatCompletionsOptions contains the optional parameters for the Client.GetChatCompletions method.
type GetChatCompletionsOptions struct {
	// placeholder for future optional parameters
}

// GetCompletionsOptions contains the optional parameters for the Client.GetCompletions method.
type GetCompletionsOptions struct {
	// placeholder for future optional parameters
}

// GetEmbeddingsOptions contains the optional parameters for the Client.GetEmbeddings method.
type GetEmbeddingsOptions struct {
	// placeholder for future optional parameters
}

// GetFileContentOptions contains the optional parameters for the Client.GetFileContent method.
type GetFileContentOptions struct {
	// placeholder for future optional parameters
}

// GetFileOptions contains the optional parameters for the Client.GetFile method.
type GetFileOptions struct {
	// placeholder for future optional parameters
}

// GetImageGenerationsOptions contains the optional parameters for the Client.GetImageGenerations method.
type GetImageGenerationsOptions struct {
	// placeholder for future optional parameters
}

// ListBatchesOptions contains the optional parameters for the Client.ListBatches method.
type ListBatchesOptions struct {
	// Identifier for the last event from the previous pagination request.
	After *string

	// Number of batches to retrieve. Defaults to 20.
	Limit *int32
}

// ListFilesOptions contains the optional parameters for the Client.ListFiles method.
type ListFilesOptions struct {
	// A value that, when provided, limits list results to files matching the corresponding purpose.
	Purpose *FilePurpose
}

// UploadFileOptions contains the optional parameters for the Client.UploadFile method.
type UploadFileOptions struct {
	// A filename to associate with the uploaded data.
	Filename *string
}
