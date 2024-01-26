package constant

const (
	EnvAppName    = "APP_NAME"
	EnvDeployment = "APP_DEPLOYMENT"

	EnvAppLogDir     = "APP_LOG_DIR"
	EnvAppMode       = "APP_MODE" //dev,fat,pro
	EnvAppRegion     = "APP_REGION"
	EnvAppZone       = "APP_ZONE"
	EnvAppHost       = "APP_HOST"
	EnvAppInstance   = "APP_INSTANCE" // application unique instance id.
	EnvAppConfAddr   = "APP_CONF_ADDR"
	EnvAppConfFormat = "APP_CONF_FORMAT"
)

const (
	DefaultDeployment = ""
	DefaultRegion     = ""
	DefaultZone       = ""
)

const (
	// KeyBalanceGroup ...
	KeyBalanceGroup = "__group"

	// DefaultBalanceGroup ...
	DefaultBalanceGroup = "default"
)
