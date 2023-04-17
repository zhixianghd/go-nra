package nra

type Config struct {
	Service           string
	BaseUri           string
	ExposeErrorReason bool
	ExposeErrorTraces bool
	ProtocolFields    *ProtocolFieldsConfig
}

type ProtocolFieldsConfig struct {
	Version string
	Code    string
	Source  string
	Notice  string
	Retry   string
}

var GlobalConfig = &Config{
	Service:           "unnamed",
	BaseUri:           "",
	ExposeErrorReason: false,
	ExposeErrorTraces: false,
	ProtocolFields: &ProtocolFieldsConfig{
		Version: "Nra-Version",
		Code:    "Nra-Code",
		Source:  "Nra-Source",
		Notice:  "Nra-Notice",
		Retry:   "Nra-Retry",
	},
}
