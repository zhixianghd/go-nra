package nra

type ErrorRsp struct {
	Code     int              `json:"code,omitempty"`
	Service  string           `json:"service,omitempty"`
	Endpoint string           `json:"endpoint,omitempty"`
	Reason   string           `json:"reason,omitempty"`
	Traces   []*ErrorTraceDto `json:"traces,omitempty"`
}

type ErrorTraceDto struct {
	Service  string `json:"service,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	Details  string `json:"details,omitempty"`
}
