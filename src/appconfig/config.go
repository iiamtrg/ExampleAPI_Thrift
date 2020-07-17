package appconfig

type AppConfig struct {
	EtcdServerEndpoints []string

	SourceMappingNewSID      string
	SourceMappingNewHost     string
	SourceMappingNewPort     string
	SourceMappingNewProtocol string
}

// Config is global AppConfig instance
var Config *AppConfig
