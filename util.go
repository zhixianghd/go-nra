package nra

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func getEndpoint(context *gin.Context) string {
	endpoint := context.Request.RequestURI
	if GlobalConfig.BaseUri != "" && strings.HasPrefix(endpoint, GlobalConfig.BaseUri) {
		endpoint = endpoint[len(GlobalConfig.BaseUri):]
	}
	return endpoint
}

func getTraces(service string, endpoint string, error error) []*ErrorTraceDto {
	var traces []*ErrorTraceDto
	traces = append(traces, &ErrorTraceDto{
		Service:  service,
		Endpoint: endpoint,
		Details:  error.Error(),
	})
	return traces
}
