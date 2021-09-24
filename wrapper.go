package client

import (
	_context "context"
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"golang.org/x/net/context"
	_nethttp "net/http"
)

// LogsSort Sort parameters when querying logs.
type LogsSort string

// List of LogsSort
const (
	LOGSSORT_TIMESTAMP_ASCENDING  LogsSort = "timestamp"
	LOGSSORT_TIMESTAMP_DESCENDING LogsSort = "-timestamp"
)

type DashboardListsAPI interface {
	CreateDashboardListItems(ctx _context.Context, dashboardListId int64, body DashboardListAddItemsRequest) (DashboardListAddItemsResponse, *_nethttp.Response, error)
	DeleteDashboardListItems(ctx _context.Context, dashboardListId int64, body DashboardListDeleteItemsRequest) (DashboardListDeleteItemsResponse, *_nethttp.Response, error)
	GetDashboardListItems(ctx _context.Context, dashboardListId int64) (DashboardListItems, *_nethttp.Response, error)
	UpdateDashboardListItems(ctx _context.Context, dashboardListId int64, body DashboardListUpdateItemsRequest) (DashboardListUpdateItemsResponse, *_nethttp.Response, error)
}

type IncidentServicesAPI interface {
	CreateIncidentService(ctx context.Context, body IncidentServiceCreateRequest) (IncidentServiceResponse, *http.Response, error)
	DeleteIncidentService(ctx context.Context, serviceId string) (*http.Response, error)
	GetIncidentService(ctx context.Context, serviceId string, o ...GetIncidentServiceOptionalParameters) (IncidentServiceResponse, *http.Response, error)
	ListIncidentServices(ctx context.Context, o ...ListIncidentServicesOptionalParameters) (IncidentServicesResponse, *http.Response, error)
	UpdateIncidentService(ctx context.Context, serviceId string, body IncidentServiceUpdateRequest) (IncidentServiceResponse, *http.Response, error)
}

type IncidentTeamsAPI interface {
}

type IncidentsAPI interface {
}

type KeyManagementAPI interface {
}

type LogsAPI interface {
	AggregateLogs(ctx _context.Context, body datadog.LogsAggregateRequest) (datadog.LogsAggregateResponse, *_nethttp.Response, error)
	ListLogs(ctx _context.Context, o ...datadog.ListLogsOptionalParameters) (datadog.LogsListResponse, *_nethttp.Response, error)
	ListLogsGet(ctx _context.Context, o ...datadog.ListLogsGetOptionalParameters) (datadog.LogsListResponse, *_nethttp.Response, error)
}

type LogsArchivesAPI interface {
	AddReadRoleToArchive(ctx context.Context, archiveId string, body RelationshipToRole) (*http.Response, error)
	CreateLogsArchive(ctx context.Context, body LogsArchiveCreateRequest) (LogsArchive, *http.Response, error)
	DeleteLogsArchive(ctx context.Context, archiveId string) (*http.Response, error)
	GetLogsArchive(ctx context.Context, archiveId string) (LogsArchive, *http.Response, error)
	GetLogsArchiveOrder(ctx context.Context) (LogsArchiveOrder, *http.Response, error)
	ListArchiveReadRoles(ctx context.Context, archiveId string) (RolesResponse, *http.Response, error)
	ListLogsArchives(ctx context.Context) (LogsArchives, *http.Response, error)
	RemoveRoleFromArchive(ctx context.Context, archiveId string, body RelationshipToRole) (*http.Response, error)
	UpdateLogsArchive(ctx context.Context, archiveId string, body LogsArchiveCreateRequest) (LogsArchive, *http.Response, error)
	UpdateLogsArchiveOrder(ctx context.Context, body LogsArchiveOrder) (LogsArchiveOrder, *http.Response, error)
}

type LogsMetricsAPI interface {
}

type MetricsAPI interface {
}

type ProcessesAPI interface {
}

type a3 interface {
}

type a3 interface {
}

type a3 interface {
}

type a3 interface {
}
