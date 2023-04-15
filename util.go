package nra

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func getSource(context *gin.Context) string {
	uri := context.Request.RequestURI
	if GlobalConfig.BaseUri != "" && strings.HasPrefix(uri, GlobalConfig.BaseUri) {
		uri = uri[len(GlobalConfig.BaseUri):]
	}
	return GlobalConfig.Service + ":" + uri
}

func getTraces(source string, error error) []*ErrorTraceDto {
	var traces []*ErrorTraceDto
	traces = append(traces, &ErrorTraceDto{
		Source:  source,
		Details: error.Error(),
	})
	return traces
}
