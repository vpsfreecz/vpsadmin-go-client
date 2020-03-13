package client

// Client represents a connection to an API server
type Client struct {
	// API URL
	Url string

	// Options for authentication method
	Authentication Authenticator

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
	// Resource Host_ip_address
	HostIpAddress *ResourceHostIpAddress
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
	// Resource Network
	Network *ResourceNetwork
	// Resource Network_interface
	NetworkInterface *ResourceNetworkInterface
	// Resource Node
	Node *ResourceNode
	// Resource Object_history
	ObjectHistory *ResourceObjectHistory
	// Resource Os_template
	OsTemplate *ResourceOsTemplate
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
	// Resource User_cluster_resource_package
	UserClusterResourcePackage *ResourceUserClusterResourcePackage
	// Resource User_namespace
	UserNamespace *ResourceUserNamespace
	// Resource User_namespace_map
	UserNamespaceMap *ResourceUserNamespaceMap
	// Resource User_session
	UserSession *ResourceUserSession
	// Resource Vps
	Vps *ResourceVps
	// Resource Vps_config
	VpsConfig *ResourceVpsConfig
	// Resource Monitored_event
	MonitoredEvent *ResourceMonitoredEvent
	// Resource User_request
	UserRequest *ResourceUserRequest
	// Resource Incoming_payment
	IncomingPayment *ResourceIncomingPayment
	// Resource Payment_stats
	PaymentStats *ResourcePaymentStats
	// Resource User_account
	UserAccount *ResourceUserAccount
	// Resource User_payment
	UserPayment *ResourceUserPayment
	// Resource Help_box
	HelpBox *ResourceHelpBox
	// Resource Outage
	Outage *ResourceOutage
	// Resource Outage_update
	OutageUpdate *ResourceOutageUpdate
	// Resource User_outage
	UserOutage *ResourceUserOutage
	// Resource Vps_outage
	VpsOutage *ResourceVpsOutage
	// Resource Vps_outage_mount
	VpsOutageMount *ResourceVpsOutageMount
	// Resource News_log
	NewsLog *ResourceNewsLog
	// Resource Action_state
	ActionState *ResourceActionState
}

// Create a new client for API at url
func New(url string) *Client {
	c := &Client{Url: url}

	c.Cluster = NewResourceCluster(c)
	c.ClusterResource = NewResourceClusterResource(c)
	c.ClusterResourcePackage = NewResourceClusterResourcePackage(c)
	c.Dataset = NewResourceDataset(c)
	c.DatasetPlan = NewResourceDatasetPlan(c)
	c.DnsResolver = NewResourceDnsResolver(c)
	c.Environment = NewResourceEnvironment(c)
	c.Export = NewResourceExport(c)
	c.HostIpAddress = NewResourceHostIpAddress(c)
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
	c.Network = NewResourceNetwork(c)
	c.NetworkInterface = NewResourceNetworkInterface(c)
	c.Node = NewResourceNode(c)
	c.ObjectHistory = NewResourceObjectHistory(c)
	c.OsTemplate = NewResourceOsTemplate(c)
	c.Pool = NewResourcePool(c)
	c.SessionToken = NewResourceSessionToken(c)
	c.SnapshotDownload = NewResourceSnapshotDownload(c)
	c.SystemConfig = NewResourceSystemConfig(c)
	c.Transaction = NewResourceTransaction(c)
	c.TransactionChain = NewResourceTransactionChain(c)
	c.User = NewResourceUser(c)
	c.UserClusterResourcePackage = NewResourceUserClusterResourcePackage(c)
	c.UserNamespace = NewResourceUserNamespace(c)
	c.UserNamespaceMap = NewResourceUserNamespaceMap(c)
	c.UserSession = NewResourceUserSession(c)
	c.Vps = NewResourceVps(c)
	c.VpsConfig = NewResourceVpsConfig(c)
	c.MonitoredEvent = NewResourceMonitoredEvent(c)
	c.UserRequest = NewResourceUserRequest(c)
	c.IncomingPayment = NewResourceIncomingPayment(c)
	c.PaymentStats = NewResourcePaymentStats(c)
	c.UserAccount = NewResourceUserAccount(c)
	c.UserPayment = NewResourceUserPayment(c)
	c.HelpBox = NewResourceHelpBox(c)
	c.Outage = NewResourceOutage(c)
	c.OutageUpdate = NewResourceOutageUpdate(c)
	c.UserOutage = NewResourceUserOutage(c)
	c.VpsOutage = NewResourceVpsOutage(c)
	c.VpsOutageMount = NewResourceVpsOutageMount(c)
	c.NewsLog = NewResourceNewsLog(c)
	c.ActionState = NewResourceActionState(c)

	return c
}
