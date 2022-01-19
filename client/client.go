package client

// Client represents a connection to an API server
type Client struct {
	// API URL
	Url string

	// Options for authentication method
	Authentication Authenticator

	// Resource Action_state
	ActionState *ResourceActionState
	// Resource Api_server
	ApiServer *ResourceApiServer
	// Resource Cluster
	Cluster *ResourceCluster
	// Resource Cluster_resource
	ClusterResource *ResourceClusterResource
	// Resource Cluster_resource_package
	ClusterResourcePackage *ResourceClusterResourcePackage
	// Resource Dataset
	Dataset *ResourceDataset
	// Resource Dataset_plan
	DatasetPlan *ResourceDatasetPlan
	// Resource Dns_resolver
	DnsResolver *ResourceDnsResolver
	// Resource Environment
	Environment *ResourceEnvironment
	// Resource Export
	Export *ResourceExport
	// Resource Export_outage
	ExportOutage *ResourceExportOutage
	// Resource Help_box
	HelpBox *ResourceHelpBox
	// Resource Host_ip_address
	HostIpAddress *ResourceHostIpAddress
	// Resource Incoming_payment
	IncomingPayment *ResourceIncomingPayment
	// Resource Integrity_check
	IntegrityCheck *ResourceIntegrityCheck
	// Resource Integrity_fact
	IntegrityFact *ResourceIntegrityFact
	// Resource Integrity_object
	IntegrityObject *ResourceIntegrityObject
	// Resource Ip_address
	IpAddress *ResourceIpAddress
	// Resource Ip_traffic
	IpTraffic *ResourceIpTraffic
	// Resource Ip_traffic_monitor
	IpTrafficMonitor *ResourceIpTrafficMonitor
	// Resource Language
	Language *ResourceLanguage
	// Resource Location
	Location *ResourceLocation
	// Resource Location_network
	LocationNetwork *ResourceLocationNetwork
	// Resource Mail_log
	MailLog *ResourceMailLog
	// Resource Mail_recipient
	MailRecipient *ResourceMailRecipient
	// Resource Mail_template
	MailTemplate *ResourceMailTemplate
	// Resource Migration_plan
	MigrationPlan *ResourceMigrationPlan
	// Resource Monitored_event
	MonitoredEvent *ResourceMonitoredEvent
	// Resource Network
	Network *ResourceNetwork
	// Resource Network_interface
	NetworkInterface *ResourceNetworkInterface
	// Resource News_log
	NewsLog *ResourceNewsLog
	// Resource Node
	Node *ResourceNode
	// Resource Object_history
	ObjectHistory *ResourceObjectHistory
	// Resource Oom_report
	OomReport *ResourceOomReport
	// Resource Os_template
	OsTemplate *ResourceOsTemplate
	// Resource Outage
	Outage *ResourceOutage
	// Resource Outage_update
	OutageUpdate *ResourceOutageUpdate
	// Resource Payment_stats
	PaymentStats *ResourcePaymentStats
	// Resource Pool
	Pool *ResourcePool
	// Resource Session_token
	SessionToken *ResourceSessionToken
	// Resource Snapshot_download
	SnapshotDownload *ResourceSnapshotDownload
	// Resource System_config
	SystemConfig *ResourceSystemConfig
	// Resource Transaction
	Transaction *ResourceTransaction
	// Resource Transaction_chain
	TransactionChain *ResourceTransactionChain
	// Resource User
	User *ResourceUser
	// Resource User_account
	UserAccount *ResourceUserAccount
	// Resource User_cluster_resource_package
	UserClusterResourcePackage *ResourceUserClusterResourcePackage
	// Resource User_namespace
	UserNamespace *ResourceUserNamespace
	// Resource User_namespace_map
	UserNamespaceMap *ResourceUserNamespaceMap
	// Resource User_outage
	UserOutage *ResourceUserOutage
	// Resource User_payment
	UserPayment *ResourceUserPayment
	// Resource User_request
	UserRequest *ResourceUserRequest
	// Resource User_session
	UserSession *ResourceUserSession
	// Resource Vps
	Vps *ResourceVps
	// Resource Vps_config
	VpsConfig *ResourceVpsConfig
	// Resource Vps_outage
	VpsOutage *ResourceVpsOutage
	// Resource Vps_outage_mount
	VpsOutageMount *ResourceVpsOutageMount
}

// Create a new client for API at url
func New(url string) *Client {
	c := &Client{Url: url}

	c.ActionState = NewResourceActionState(c)
	c.ApiServer = NewResourceApiServer(c)
	c.Cluster = NewResourceCluster(c)
	c.ClusterResource = NewResourceClusterResource(c)
	c.ClusterResourcePackage = NewResourceClusterResourcePackage(c)
	c.Dataset = NewResourceDataset(c)
	c.DatasetPlan = NewResourceDatasetPlan(c)
	c.DnsResolver = NewResourceDnsResolver(c)
	c.Environment = NewResourceEnvironment(c)
	c.Export = NewResourceExport(c)
	c.ExportOutage = NewResourceExportOutage(c)
	c.HelpBox = NewResourceHelpBox(c)
	c.HostIpAddress = NewResourceHostIpAddress(c)
	c.IncomingPayment = NewResourceIncomingPayment(c)
	c.IntegrityCheck = NewResourceIntegrityCheck(c)
	c.IntegrityFact = NewResourceIntegrityFact(c)
	c.IntegrityObject = NewResourceIntegrityObject(c)
	c.IpAddress = NewResourceIpAddress(c)
	c.IpTraffic = NewResourceIpTraffic(c)
	c.IpTrafficMonitor = NewResourceIpTrafficMonitor(c)
	c.Language = NewResourceLanguage(c)
	c.Location = NewResourceLocation(c)
	c.LocationNetwork = NewResourceLocationNetwork(c)
	c.MailLog = NewResourceMailLog(c)
	c.MailRecipient = NewResourceMailRecipient(c)
	c.MailTemplate = NewResourceMailTemplate(c)
	c.MigrationPlan = NewResourceMigrationPlan(c)
	c.MonitoredEvent = NewResourceMonitoredEvent(c)
	c.Network = NewResourceNetwork(c)
	c.NetworkInterface = NewResourceNetworkInterface(c)
	c.NewsLog = NewResourceNewsLog(c)
	c.Node = NewResourceNode(c)
	c.ObjectHistory = NewResourceObjectHistory(c)
	c.OomReport = NewResourceOomReport(c)
	c.OsTemplate = NewResourceOsTemplate(c)
	c.Outage = NewResourceOutage(c)
	c.OutageUpdate = NewResourceOutageUpdate(c)
	c.PaymentStats = NewResourcePaymentStats(c)
	c.Pool = NewResourcePool(c)
	c.SessionToken = NewResourceSessionToken(c)
	c.SnapshotDownload = NewResourceSnapshotDownload(c)
	c.SystemConfig = NewResourceSystemConfig(c)
	c.Transaction = NewResourceTransaction(c)
	c.TransactionChain = NewResourceTransactionChain(c)
	c.User = NewResourceUser(c)
	c.UserAccount = NewResourceUserAccount(c)
	c.UserClusterResourcePackage = NewResourceUserClusterResourcePackage(c)
	c.UserNamespace = NewResourceUserNamespace(c)
	c.UserNamespaceMap = NewResourceUserNamespaceMap(c)
	c.UserOutage = NewResourceUserOutage(c)
	c.UserPayment = NewResourceUserPayment(c)
	c.UserRequest = NewResourceUserRequest(c)
	c.UserSession = NewResourceUserSession(c)
	c.Vps = NewResourceVps(c)
	c.VpsConfig = NewResourceVpsConfig(c)
	c.VpsOutage = NewResourceVpsOutage(c)
	c.VpsOutageMount = NewResourceVpsOutageMount(c)

	return c
}
