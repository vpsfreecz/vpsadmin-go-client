package client

// Type for resource Vps
type ResourceVps struct {
	// Pointer to client
	Client *Client

	// Resource Vps.State_log
	StateLog *ResourceVpsStateLog
	// Resource Vps.Config
	Config *ResourceVpsConfig
	// Resource Vps.Feature
	Feature *ResourceVpsFeature
	// Resource Vps.Mount
	Mount *ResourceVpsMount
	// Resource Vps.Maintenance_window
	MaintenanceWindow *ResourceVpsMaintenanceWindow
	// Resource Vps.Console_token
	ConsoleToken *ResourceVpsConsoleToken
	// Resource Vps.Status
	Status *ResourceVpsStatus
	// Action Vps#Index
	Index *ActionVpsIndex
	// Action Vps#Index
	List *ActionVpsIndex
	// Action Vps#Create
	Create *ActionVpsCreate
	// Action Vps#Create
	New *ActionVpsCreate
	// Action Vps#Show
	Show *ActionVpsShow
	// Action Vps#Show
	Find *ActionVpsShow
	// Action Vps#Update
	Update *ActionVpsUpdate
	// Action Vps#Delete
	Delete *ActionVpsDelete
	// Action Vps#Delete
	Destroy *ActionVpsDelete
	// Action Vps#Start
	Start *ActionVpsStart
	// Action Vps#Restart
	Restart *ActionVpsRestart
	// Action Vps#Stop
	Stop *ActionVpsStop
	// Action Vps#Passwd
	Passwd *ActionVpsPasswd
	// Action Vps#Reinstall
	Reinstall *ActionVpsReinstall
	// Action Vps#Migrate
	Migrate *ActionVpsMigrate
	// Action Vps#Clone
	Clone *ActionVpsClone
	// Action Vps#Swap_with
	SwapWith *ActionVpsSwapWith
	// Action Vps#Deploy_public_key
	DeployPublicKey *ActionVpsDeployPublicKey
	// Action Vps#Set_maintenance
	SetMaintenance *ActionVpsSetMaintenance
}

func NewResourceVps(client *Client) *ResourceVps {
	actionIndex := NewActionVpsIndex(client)
	actionCreate := NewActionVpsCreate(client)
	actionShow := NewActionVpsShow(client)
	actionUpdate := NewActionVpsUpdate(client)
	actionDelete := NewActionVpsDelete(client)
	actionStart := NewActionVpsStart(client)
	actionRestart := NewActionVpsRestart(client)
	actionStop := NewActionVpsStop(client)
	actionPasswd := NewActionVpsPasswd(client)
	actionReinstall := NewActionVpsReinstall(client)
	actionMigrate := NewActionVpsMigrate(client)
	actionClone := NewActionVpsClone(client)
	actionSwapWith := NewActionVpsSwapWith(client)
	actionDeployPublicKey := NewActionVpsDeployPublicKey(client)
	actionSetMaintenance := NewActionVpsSetMaintenance(client)

	return &ResourceVps{
		Client: client,
		StateLog: NewResourceVpsStateLog(client),
		Config: NewResourceVpsConfig(client),
		Feature: NewResourceVpsFeature(client),
		Mount: NewResourceVpsMount(client),
		MaintenanceWindow: NewResourceVpsMaintenanceWindow(client),
		ConsoleToken: NewResourceVpsConsoleToken(client),
		Status: NewResourceVpsStatus(client),
		Index: actionIndex,
		List: actionIndex,
		Create: actionCreate,
		New: actionCreate,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
		Start: actionStart,
		Restart: actionRestart,
		Stop: actionStop,
		Passwd: actionPasswd,
		Reinstall: actionReinstall,
		Migrate: actionMigrate,
		Clone: actionClone,
		SwapWith: actionSwapWith,
		DeployPublicKey: actionDeployPublicKey,
		SetMaintenance: actionSetMaintenance,
	}
}
