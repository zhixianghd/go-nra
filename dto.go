package nra

type ErrorRsp struct {
	Code   int              `json:"code,omitempty"`
	Source string           `json:"source,omitempty"`
	Reason string           `json:"reason,omitempty"`
	Traces []*ErrorTraceDto `json:"traces,omitempty"`
}

type ErrorTraceDto struct {
	Source  string `json:"source,omitempty"`
	Details string `json:"details,omitempty"`
}
