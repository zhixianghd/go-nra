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
	Source  string
	Code    string
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
		Source:  "Nra-Source",
		Code:    "Nra-Code",
		Notice:  "Nra-Notice",
		Retry:   "Nra-Retry",
	},
}
