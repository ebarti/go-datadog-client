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
	CreateDashboardListItems(ctx _context.Context, dashboardListId int64, body datadog.DashboardListAddItemsRequest) (datadog.DashboardListAddItemsResponse, *_nethttp.Response, error)
	DeleteDashboardListItems(ctx _context.Context, dashboardListId int64, body datadog.DashboardListDeleteItemsRequest) (datadog.DashboardListDeleteItemsResponse, *_nethttp.Response, error)
	GetDashboardListItems(ctx _context.Context, dashboardListId int64) (datadog.DashboardListItems, *_nethttp.Response, error)
	UpdateDashboardListItems(ctx _context.Context, dashboardListId int64, body datadog.DashboardListUpdateItemsRequest) (datadog.DashboardListUpdateItemsResponse, *_nethttp.Response, error)
}

type IncidentServicesAPI interface {
	CreateIncidentService(ctx context.Context, body datadog.IncidentServiceCreateRequest) (datadog.IncidentServiceResponse, *_nethttp.Response, error)
	DeleteIncidentService(ctx context.Context, serviceId string) (*_nethttp.Response, error)
	GetIncidentService(ctx context.Context, serviceId string, o ...datadog.GetIncidentServiceOptionalParameters) (datadog.IncidentServiceResponse, *_nethttp.Response, error)
	ListIncidentServices(ctx context.Context, o ...datadog.ListIncidentServicesOptionalParameters) (datadog.IncidentServicesResponse, *_nethttp.Response, error)
	UpdateIncidentService(ctx context.Context, serviceId string, body datadog.IncidentServiceUpdateRequest) (datadog.IncidentServiceResponse, *_nethttp.Response, error)
}

type IncidentTeamsAPI interface {
	CreateIncidentTeam(ctx context.Context, body datadog.IncidentTeamCreateRequest) (datadog.IncidentTeamResponse, *_nethttp.Response, error)
	DeleteIncidentTeam(ctx context.Context, teamId string) (*_nethttp.Response, error)
	GetIncidentTeam(ctx context.Context, teamId string, o ...datadog.GetIncidentTeamOptionalParameters) (datadog.IncidentTeamResponse, *_nethttp.Response, error)
	ListIncidentTeams(ctx context.Context, o ...datadog.ListIncidentTeamsOptionalParameters) (datadog.IncidentTeamsResponse, *_nethttp.Response, error)
	UpdateIncidentTeam(ctx context.Context, teamId string, body datadog.IncidentTeamUpdateRequest) (datadog.IncidentTeamResponse, *_nethttp.Response, error)
}

type IncidentsAPI interface {
	CreateIncident(ctx context.Context, body datadog.IncidentCreateRequest) (datadog.IncidentResponse, *_nethttp.Response, error)
	DeleteIncident(ctx context.Context, incidentId string) (*_nethttp.Response, error)
	GetIncident(ctx context.Context, incidentId string, o ...datadog.GetIncidentOptionalParameters) (datadog.IncidentResponse, *_nethttp.Response, error)
	ListIncidents(ctx context.Context, o ...datadog.ListIncidentsOptionalParameters) (datadog.IncidentsResponse, *_nethttp.Response, error)
	UpdateIncident(ctx context.Context, incidentId string, body datadog.IncidentUpdateRequest) (datadog.IncidentResponse, *_nethttp.Response, error)
}

type KeyManagementAPI interface {
	CreateAPIKey(ctx context.Context, body datadog.APIKeyCreateRequest) (datadog.APIKeyResponse, *_nethttp.Response, error)
	CreateCurrentUserApplicationKey(ctx context.Context, body datadog.ApplicationKeyCreateRequest) (datadog.ApplicationKeyResponse, *_nethttp.Response, error)
	DeleteAPIKey(ctx context.Context, apiKeyId string) (*_nethttp.Response, error)
	DeleteApplicationKey(ctx context.Context, appKeyId string) (*_nethttp.Response, error)
	DeleteCurrentUserApplicationKey(ctx context.Context, appKeyId string) (*_nethttp.Response, error)
	GetAPIKey(ctx context.Context, apiKeyId string, o ...datadog.GetAPIKeyOptionalParameters) (datadog.APIKeyResponse, *_nethttp.Response, error)
	GetApplicationKey(ctx context.Context, appKeyId string, o ...datadog.GetApplicationKeyOptionalParameters) (datadog.ApplicationKeyResponse, *_nethttp.Response, error)
	GetCurrentUserApplicationKey(ctx context.Context, appKeyId string) (datadog.ApplicationKeyResponse, *_nethttp.Response, error)
	ListAPIKeys(ctx context.Context, o ...datadog.ListAPIKeysOptionalParameters) (datadog.APIKeysResponse, *_nethttp.Response, error)
	ListApplicationKeys(ctx context.Context, o ...datadog.ListApplicationKeysOptionalParameters) (datadog.ListApplicationKeysResponse, *_nethttp.Response, error)
	ListCurrentUserApplicationKeys(ctx context.Context, o ...datadog.ListCurrentUserApplicationKeysOptionalParameters) (datadog.ListApplicationKeysResponse, *_nethttp.Response, error)
	UpdateAPIKey(ctx context.Context, apiKeyId string, body datadog.APIKeyUpdateRequest) (datadog.APIKeyResponse, *_nethttp.Response, error)
	UpdateApplicationKey(ctx context.Context, appKeyId string, body datadog.ApplicationKeyUpdateRequest) (datadog.ApplicationKeyResponse, *_nethttp.Response, error)
	UpdateCurrentUserApplicationKey(ctx context.Context, appKeyId string, body datadog.ApplicationKeyUpdateRequest) (datadog.ApplicationKeyResponse, *_nethttp.Response, error)
}

type LogsAPI interface {
	AggregateLogs(ctx _context.Context, body datadog.LogsAggregateRequest) (datadog.LogsAggregateResponse, *_nethttp.Response, error)
	ListLogs(ctx _context.Context, o ...datadog.ListLogsOptionalParameters) (datadog.LogsListResponse, *_nethttp.Response, error)
	ListLogsGet(ctx _context.Context, o ...datadog.ListLogsGetOptionalParameters) (datadog.LogsListResponse, *_nethttp.Response, error)
}

type LogsArchivesAPI interface {
	AddReadRoleToArchive(ctx context.Context, archiveId string, body datadog.RelationshipToRole) (*_nethttp.Response, error)
	CreateLogsArchive(ctx context.Context, body datadog.LogsArchiveCreateRequest) (datadog.LogsArchive, *_nethttp.Response, error)
	DeleteLogsArchive(ctx context.Context, archiveId string) (*_nethttp.Response, error)
	GetLogsArchive(ctx context.Context, archiveId string) (datadog.LogsArchive, *_nethttp.Response, error)
	GetLogsArchiveOrder(ctx context.Context) (datadog.LogsArchiveOrder, *_nethttp.Response, error)
	ListArchiveReadRoles(ctx context.Context, archiveId string) (datadog.RolesResponse, *_nethttp.Response, error)
	ListLogsArchives(ctx context.Context) (datadog.LogsArchives, *_nethttp.Response, error)
	RemoveRoleFromArchive(ctx context.Context, archiveId string, body datadog.RelationshipToRole) (*_nethttp.Response, error)
	UpdateLogsArchive(ctx context.Context, archiveId string, body datadog.LogsArchiveCreateRequest) (datadog.LogsArchive, *_nethttp.Response, error)
	UpdateLogsArchiveOrder(ctx context.Context, body datadog.LogsArchiveOrder) (datadog.LogsArchiveOrder, *_nethttp.Response, error)
}

type LogsMetricsAPI interface {
	CreateLogsMetric(ctx context.Context, body datadog.LogsMetricCreateRequest) (datadog.LogsMetricResponse, *_nethttp.Response, error)
	DeleteLogsMetric(ctx context.Context, metricId string) (*_nethttp.Response, error)
	GetLogsMetric(ctx context.Context, metricId string) (datadog.LogsMetricResponse, *_nethttp.Response, error)
	ListLogsMetrics(ctx context.Context) (datadog.LogsMetricsResponse, *_nethttp.Response, error)
	UpdateLogsMetric(ctx context.Context, metricId string, body datadog.LogsMetricUpdateRequest) (datadog.LogsMetricResponse, *_nethttp.Response, error)
}

type MetricsAPI interface {
	CreateTagConfiguration(ctx context.Context, metricName string, body datadog.MetricTagConfigurationCreateRequest) (datadog.MetricTagConfigurationResponse, *_nethttp.Response, error)
	DeleteTagConfiguration(ctx context.Context, metricName string) (*_nethttp.Response, error)
	ListTagConfigurationByName(ctx context.Context, metricName string) (datadog.MetricTagConfigurationResponse, *_nethttp.Response, error)
	ListTagConfigurations(ctx context.Context, o ...datadog.ListTagConfigurationsOptionalParameters) (datadog.MetricsAndMetricTagConfigurationsResponse, *_nethttp.Response, error)
	ListTagsByMetricName(ctx context.Context, metricName string) (datadog.MetricAllTagsResponse, *_nethttp.Response, error)
	ListVolumesByMetricName(ctx context.Context, metricName string) (datadog.MetricVolumesResponse, *_nethttp.Response, error)
	UpdateTagConfiguration(ctx context.Context, metricName string, body datadog.MetricTagConfigurationUpdateRequest) (datadog.MetricTagConfigurationResponse, *_nethttp.Response, error)
}

type ProcessesAPI interface {
	ListProcesses(ctx context.Context, o ...datadog.ListProcessesOptionalParameters) (datadog.ProcessSummariesResponse, *_nethttp.Response, error)
}

type RolesAPI interface {
	AddPermissionToRole(ctx context.Context, roleId string, body datadog.RelationshipToPermission) (datadog.PermissionsResponse, *_nethttp.Response, error)
	AddUserToRole(ctx context.Context, roleId string, body datadog.RelationshipToUser) (datadog.UsersResponse, *_nethttp.Response, error)
	CreateRole(ctx context.Context, body datadog.RoleCreateRequest) (datadog.RoleCreateResponse, *_nethttp.Response, error)
	DeleteRole(ctx context.Context, roleId string) (*_nethttp.Response, error)
	GetRole(ctx context.Context, roleId string) (datadog.RoleResponse, *_nethttp.Response, error)
	ListPermissions(ctx context.Context) (datadog.PermissionsResponse, *_nethttp.Response, error)
	ListRolePermissions(ctx context.Context, roleId string) (datadog.PermissionsResponse, *_nethttp.Response, error)
	ListRoleUsers(ctx context.Context, roleId string, o ...datadog.ListRoleUsersOptionalParameters) (datadog.UsersResponse, *_nethttp.Response, error)
	ListRoles(ctx context.Context, o ...datadog.ListRolesOptionalParameters) (datadog.RolesResponse, *_nethttp.Response, error)
	RemovePermissionFromRole(ctx context.Context, roleId string, body datadog.RelationshipToPermission) (datadog.PermissionsResponse, *_nethttp.Response, error)
	RemoveUserFromRole(ctx context.Context, roleId string, body datadog.RelationshipToUser) (datadog.UsersResponse, *_nethttp.Response, error)
	UpdateRole(ctx context.Context, roleId string, body datadog.RoleUpdateRequest) (datadog.RoleUpdateResponse, *_nethttp.Response, error)
}

type SecurityMonitoringAPI interface {
	CreateSecurityFilter(ctx context.Context, body datadog.SecurityFilterCreateRequest) (datadog.SecurityFilterResponse, *_nethttp.Response, error)
	CreateSecurityMonitoringRule(ctx context.Context, body datadog.SecurityMonitoringRuleCreatePayload) (datadog.SecurityMonitoringRuleResponse, *_nethttp.Response, error)
	DeleteSecurityFilter(ctx context.Context, securityFilterId string) (*_nethttp.Response, error)
	DeleteSecurityMonitoringRule(ctx context.Context, ruleId string) (*_nethttp.Response, error)
	GetSecurityFilter(ctx context.Context, securityFilterId string) (datadog.SecurityFilterResponse, *_nethttp.Response, error)
	GetSecurityMonitoringRule(ctx context.Context, ruleId string) (datadog.SecurityMonitoringRuleResponse, *_nethttp.Response, error)
	ListSecurityFilters(ctx context.Context) (datadog.SecurityFiltersResponse, *_nethttp.Response, error)
	ListSecurityMonitoringRules(ctx context.Context, o ...datadog.ListSecurityMonitoringRulesOptionalParameters) (datadog.SecurityMonitoringListRulesResponse, *_nethttp.Response, error)
	ListSecurityMonitoringSignals(ctx context.Context, o ...datadog.ListSecurityMonitoringSignalsOptionalParameters) (datadog.SecurityMonitoringSignalsListResponse, *_nethttp.Response, error)
	SearchSecurityMonitoringSignals(ctx context.Context, o ...datadog.SearchSecurityMonitoringSignalsOptionalParameters) (datadog.SecurityMonitoringSignalsListResponse, *_nethttp.Response, error)
	UpdateSecurityFilter(ctx context.Context, securityFilterId string, body datadog.SecurityFilterUpdateRequest) (datadog.SecurityFilterResponse, *_nethttp.Response, error)
	UpdateSecurityMonitoringRule(ctx context.Context, ruleId string, body datadog.SecurityMonitoringRuleUpdatePayload) (datadog.SecurityMonitoringRuleResponse, *_nethttp.Response, error)
}

type ServiceAccountsAPI interface {
	CreateServiceAccountApplicationKey(ctx context.Context, serviceAccountId string, body datadog.ApplicationKeyCreateRequest) (datadog.ApplicationKeyResponse, *_nethttp.Response, error)
	DeleteServiceAccountApplicationKey(ctx context.Context, serviceAccountId string, appKeyId string) (*_nethttp.Response, error)
	GetServiceAccountApplicationKey(ctx context.Context, serviceAccountId string, appKeyId string) (datadog.PartialApplicationKeyResponse, *_nethttp.Response, error)
	ListServiceAccountApplicationKeys(ctx context.Context, serviceAccountId string, o ...datadog.ListServiceAccountApplicationKeysOptionalParameters) (datadog.ListApplicationKeysResponse, *_nethttp.Response, error)
	UpdateServiceAccountApplicationKey(ctx context.Context, serviceAccountId string, appKeyId string, body datadog.ApplicationKeyUpdateRequest) (datadog.PartialApplicationKeyResponse, *_nethttp.Response, error)
}

type UsersAPI interface {
	CreateServiceAccount(ctx context.Context, body datadog.ServiceAccountCreateRequest) (datadog.UserResponse, *_nethttp.Response, error)
	CreateUser(ctx context.Context, body datadog.UserCreateRequest) (datadog.UserResponse, *_nethttp.Response, error)
	DisableUser(ctx context.Context, userId string) (*_nethttp.Response, error)
	GetInvitation(ctx context.Context, userInvitationUuid string) (datadog.UserInvitationResponse, *_nethttp.Response, error)
	GetUser(ctx context.Context, userId string) (datadog.UserResponse, *_nethttp.Response, error)
	ListUserOrganizations(ctx context.Context, userId string) (datadog.UserResponse, *_nethttp.Response, error)
	ListUserPermissions(ctx context.Context, userId string) (datadog.PermissionsResponse, *_nethttp.Response, error)
	ListUsers(ctx context.Context, o ...datadog.ListUsersOptionalParameters) (datadog.UsersResponse, *_nethttp.Response, error)
	SendInvitations(ctx context.Context, body datadog.UserInvitationsRequest) (datadog.UserInvitationsResponse, *_nethttp.Response, error)
	UpdateUser(ctx context.Context, userId string, body datadog.UserUpdateRequest) (datadog.UserResponse, *_nethttp.Response, error)
}
