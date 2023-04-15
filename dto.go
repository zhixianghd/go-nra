package nra

type ErrorRsp struct {
	Source string           `json:"source,omitempty"`
	Code   int              `json:"code,omitempty"`
	Reason string           `json:"reason,omitempty"`
	Traces []*ErrorTraceDto `json:"traces,omitempty"`
}

type ErrorTraceDto struct {
	Source  string `json:"source,omitempty"`
	Details string `json:"details,omitempty"`
}
