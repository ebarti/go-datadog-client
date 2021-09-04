package client

import (
	"context"
	"time"
)

// LogsSort Sort parameters when querying logs.
type LogsSort string

// List of LogsSort
const (
	LOGSSORT_TIMESTAMP_ASCENDING  LogsSort = "timestamp"
	LOGSSORT_TIMESTAMP_DESCENDING LogsSort = "-timestamp"
)

type DatadogClient interface {
}

type ListLogsGetOptionalParameters struct {
	FilterQuery        *string
	FilterIndex        *string
	FilterFrom         *time.Time
	FilterTo           *time.Time
	SortTimeDescending bool
	PageCursor         *string
	PageLimit          *int32
}

type LogsListResponse struct {
	// Array of logs matching the request.
	Data  *[]Log                 `json:"data,omitempty"`
	Links *LogsListResponseLinks `json:"links,omitempty"`
	Meta  *LogsResponseMetadata  `json:"meta,omitempty"`
}

// Log Object description of a log after being processed and stored by Datadog.
type Log struct {
	Attributes *LogAttributes `json:"attributes,omitempty"`
	// Unique ID of the Log.
	Id   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// LogAttributes JSON object containing all log attributes and their associated values.
type LogAttributes struct {
	// JSON object of attributes from your log.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Name of the machine from where the logs are being sent.
	Host *string `json:"host,omitempty"`
	// The message [reserved attribute](https://docs.datadoghq.com/logs/log_collection/#reserved-attributes) of your log. By default, Datadog ingests the value of the message attribute as the body of the log entry. That value is then highlighted and displayed in the Logstream, where it is indexed for full text search.
	Message *string `json:"message,omitempty"`
	// The name of the application or service generating the log events. It is used to switch from Logs to APM, so make sure you define the same value when you use both products.
	Service *string `json:"service,omitempty"`
	// Status of the message associated with your log.
	Status *string `json:"status,omitempty"`
	// Array of tags associated with your log.
	Tags *[]string `json:"tags,omitempty"`
	// Timestamp of your log.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type LogsListResponseLinks struct {
	// Link for the next set of results. Note that the request can also be made using the POST endpoint.
	Next *string `json:"next,omitempty"`
}

// LogsResponseMetadata The metadata associated with a request
type LogsResponseMetadata struct {
	// The time elapsed in milliseconds
	Elapsed *int64                    `json:"elapsed,omitempty"`
	Page    *LogsResponseMetadataPage `json:"page,omitempty"`
	// The identifier of the request
	RequestId *string `json:"request_id,omitempty"`
	Status    *string `json:"status,omitempty"` // done, timeout
	// A list of warnings (non fatal errors) encountered, partial results might be returned if warnings are present in the response.
	Warnings *[]LogsWarning `json:"warnings,omitempty"`
}

// LogsWarning A warning message indicating something that went wrong with the query
type LogsWarning struct {
	// A unique code for this type of warning
	Code *string `json:"code,omitempty"`
	// A detailed explanation of this specific warning
	Detail *string `json:"detail,omitempty"`
	// A short human-readable summary of the warning
	Title *string `json:"title,omitempty"`
}

// LogsResponseMetadataPage Paging attributes.
type LogsResponseMetadataPage struct {
	// The cursor to use to get the next results, if any. To make the next request, use the same. parameters with the addition of the `page[cursor]`.
	After *string `json:"after,omitempty"`
}

type DDLogsClient interface {
	ListLogs(context.Context, ...ListLogsGetOptionalParameters)
}
