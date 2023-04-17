package nra

type Config struct {
	Service           string
	BaseUri           string
	ExposeErrorReason bool
	ExposeErrorTraces bool
	ProtocolFields    *ProtocolFieldsConfig
}

type ProtocolFieldsConfig struct {
	Version  string
	Code     string
	Service  string
	Endpoint string
	Notice   string
	Retry    string
}

var GlobalConfig = &Config{
	Service:           "unnamed",
	BaseUri:           "",
	ExposeErrorReason: false,
	ExposeErrorTraces: false,
	ProtocolFields: &ProtocolFieldsConfig{
		Version:  "Nra-Version",
		Code:     "Nra-Code",
		Service:  "Nra-Service",
		Endpoint: "Nra-Endpoint",
		Notice:   "Nra-Notice",
		Retry:    "Nra-Retry",
	},
}
