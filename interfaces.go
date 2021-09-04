package client

import "context"

type DatadogClient interface {

}

type DDLogsClient interface {
	ListLogs(ctx context.Context)
}
