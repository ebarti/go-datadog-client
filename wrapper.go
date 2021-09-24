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
	CreateIncidentTeam(ctx context.Context, body IncidentTeamCreateRequest) (IncidentTeamResponse, *http.Response, error)
	DeleteIncidentTeam(ctx context.Context, teamId string) (*http.Response, error)
	GetIncidentTeam(ctx context.Context, teamId string, o ...GetIncidentTeamOptionalParameters) (IncidentTeamResponse, *http.Response, error)
	ListIncidentTeams(ctx context.Context, o ...ListIncidentTeamsOptionalParameters) (IncidentTeamsResponse, *http.Response, error)
	UpdateIncidentTeam(ctx context.Context, teamId string, body IncidentTeamUpdateRequest) (IncidentTeamResponse, *http.Response, error)
}

type IncidentsAPI interface {
	CreateIncident(ctx context.Context, body IncidentCreateRequest) (IncidentResponse, *http.Response, error)
	DeleteIncident(ctx context.Context, incidentId string) (*http.Response, error)
	GetIncident(ctx context.Context, incidentId string, o ...GetIncidentOptionalParameters) (IncidentResponse, *http.Response, error)
	ListIncidents(ctx context.Context, o ...ListIncidentsOptionalParameters) (IncidentsResponse, *http.Response, error)
	UpdateIncident(ctx context.Context, incidentId string, body IncidentUpdateRequest) (IncidentResponse, *http.Response, error)
}

type KeyManagementAPI interface {
	CreateAPIKey(ctx context.Context, body APIKeyCreateRequest) (APIKeyResponse, *http.Response, error)
	CreateCurrentUserApplicationKey(ctx context.Context, body ApplicationKeyCreateRequest) (ApplicationKeyResponse, *http.Response, error)
	DeleteAPIKey(ctx context.Context, apiKeyId string) (*http.Response, error)
	DeleteApplicationKey(ctx context.Context, appKeyId string) (*http.Response, error)
	DeleteCurrentUserApplicationKey(ctx context.Context, appKeyId string) (*http.Response, error)
	GetAPIKey(ctx context.Context, apiKeyId string, o ...GetAPIKeyOptionalParameters) (APIKeyResponse, *http.Response, error)
	GetApplicationKey(ctx context.Context, appKeyId string, o ...GetApplicationKeyOptionalParameters) (ApplicationKeyResponse, *http.Response, error)
	GetCurrentUserApplicationKey(ctx context.Context, appKeyId string) (ApplicationKeyResponse, *http.Response, error)
	ListAPIKeys(ctx context.Context, o ...ListAPIKeysOptionalParameters) (APIKeysResponse, *http.Response, error)
	ListApplicationKeys(ctx context.Context, o ...ListApplicationKeysOptionalParameters) (ListApplicationKeysResponse, *http.Response, error)
	ListCurrentUserApplicationKeys(ctx context.Context, o ...ListCurrentUserApplicationKeysOptionalParameters) (ListApplicationKeysResponse, *http.Response, error)
	UpdateAPIKey(ctx context.Context, apiKeyId string, body APIKeyUpdateRequest) (APIKeyResponse, *http.Response, error)
	UpdateApplicationKey(ctx context.Context, appKeyId string, body ApplicationKeyUpdateRequest) (ApplicationKeyResponse, *http.Response, error)
	UpdateCurrentUserApplicationKey(ctx context.Context, appKeyId string, body ApplicationKeyUpdateRequest) (ApplicationKeyResponse, *http.Response, error)
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
	CreateLogsMetric(ctx context.Context, body LogsMetricCreateRequest) (LogsMetricResponse, *http.Response, error)
	DeleteLogsMetric(ctx context.Context, metricId string) (*http.Response, error)
	GetLogsMetric(ctx context.Context, metricId string) (LogsMetricResponse, *http.Response, error)
	ListLogsMetrics(ctx context.Context) (LogsMetricsResponse, *http.Response, error)
	UpdateLogsMetric(ctx context.Context, metricId string, body LogsMetricUpdateRequest) (LogsMetricResponse, *http.Response, error)
}

type MetricsAPI interface {
	CreateTagConfiguration(ctx context.Context, metricName string, body MetricTagConfigurationCreateRequest) (MetricTagConfigurationResponse, *http.Response, error)
	DeleteTagConfiguration(ctx context.Context, metricName string) (*http.Response, error)
	ListTagConfigurationByName(ctx context.Context, metricName string) (MetricTagConfigurationResponse, *http.Response, error)
	ListTagConfigurations(ctx context.Context, o ...ListTagConfigurationsOptionalParameters) (MetricsAndMetricTagConfigurationsResponse, *http.Response, error)
	ListTagsByMetricName(ctx context.Context, metricName string) (MetricAllTagsResponse, *http.Response, error)
	ListVolumesByMetricName(ctx context.Context, metricName string) (MetricVolumesResponse, *http.Response, error)
	UpdateTagConfiguration(ctx context.Context, metricName string, body MetricTagConfigurationUpdateRequest) (MetricTagConfigurationResponse, *http.Response, error)
}

type ProcessesAPI interface {
}

type RolesAPI interface {
}

type SecurityMonitoringAPI interface {
}

type ServiceAccountsAPI interface {
}

type UsersAPI interface {
}
