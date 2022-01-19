package client

// Type for resource Vps
type ResourceVps struct {
	// Pointer to client
	Client *Client

	// Resource Vps.Config
	Config *ResourceVpsConfig
	// Resource Vps.Console_token
	ConsoleToken *ResourceVpsConsoleToken
	// Resource Vps.Feature
	Feature *ResourceVpsFeature
	// Resource Vps.Maintenance_window
	MaintenanceWindow *ResourceVpsMaintenanceWindow
	// Resource Vps.Mount
	Mount *ResourceVpsMount
	// Resource Vps.State_log
	StateLog *ResourceVpsStateLog
	// Resource Vps.Status
	Status *ResourceVpsStatus
	// Action Vps#Boot
	Boot *ActionVpsBoot
	// Action Vps#Clone
	Clone *ActionVpsClone
	// Action Vps#Create
	Create *ActionVpsCreate
	// Action Vps#Create
	New *ActionVpsCreate
	// Action Vps#Delete
	Delete *ActionVpsDelete
	// Action Vps#Delete
	Destroy *ActionVpsDelete
	// Action Vps#Deploy_public_key
	DeployPublicKey *ActionVpsDeployPublicKey
	// Action Vps#Index
	Index *ActionVpsIndex
	// Action Vps#Index
	List *ActionVpsIndex
	// Action Vps#Migrate
	Migrate *ActionVpsMigrate
	// Action Vps#Passwd
	Passwd *ActionVpsPasswd
	// Action Vps#Reinstall
	Reinstall *ActionVpsReinstall
	// Action Vps#Replace
	Replace *ActionVpsReplace
	// Action Vps#Restart
	Restart *ActionVpsRestart
	// Action Vps#Set_maintenance
	SetMaintenance *ActionVpsSetMaintenance
	// Action Vps#Show
	Show *ActionVpsShow
	// Action Vps#Show
	Find *ActionVpsShow
	// Action Vps#Start
	Start *ActionVpsStart
	// Action Vps#Stop
	Stop *ActionVpsStop
	// Action Vps#Swap_with
	SwapWith *ActionVpsSwapWith
	// Action Vps#Update
	Update *ActionVpsUpdate
}

func NewResourceVps(client *Client) *ResourceVps {
	actionBoot := NewActionVpsBoot(client)
	actionClone := NewActionVpsClone(client)
	actionCreate := NewActionVpsCreate(client)
	actionDelete := NewActionVpsDelete(client)
	actionDeployPublicKey := NewActionVpsDeployPublicKey(client)
	actionIndex := NewActionVpsIndex(client)
	actionMigrate := NewActionVpsMigrate(client)
	actionPasswd := NewActionVpsPasswd(client)
	actionReinstall := NewActionVpsReinstall(client)
	actionReplace := NewActionVpsReplace(client)
	actionRestart := NewActionVpsRestart(client)
	actionSetMaintenance := NewActionVpsSetMaintenance(client)
	actionShow := NewActionVpsShow(client)
	actionStart := NewActionVpsStart(client)
	actionStop := NewActionVpsStop(client)
	actionSwapWith := NewActionVpsSwapWith(client)
	actionUpdate := NewActionVpsUpdate(client)

	return &ResourceVps{
		Client: client,
		Config: NewResourceVpsConfig(client),
		ConsoleToken: NewResourceVpsConsoleToken(client),
		Feature: NewResourceVpsFeature(client),
		MaintenanceWindow: NewResourceVpsMaintenanceWindow(client),
		Mount: NewResourceVpsMount(client),
		StateLog: NewResourceVpsStateLog(client),
		Status: NewResourceVpsStatus(client),
		Boot: actionBoot,
		Clone: actionClone,
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
		DeployPublicKey: actionDeployPublicKey,
		Index: actionIndex,
		List: actionIndex,
		Migrate: actionMigrate,
		Passwd: actionPasswd,
		Reinstall: actionReinstall,
		Replace: actionReplace,
		Restart: actionRestart,
		SetMaintenance: actionSetMaintenance,
		Show: actionShow,
		Find: actionShow,
		Start: actionStart,
		Stop: actionStop,
		SwapWith: actionSwapWith,
		Update: actionUpdate,
	}
}
