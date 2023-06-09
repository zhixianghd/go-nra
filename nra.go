package nra

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
)

const (
	VersionMin     = 10000
	VersionCurrent = 10000
)
const (
	CodeNone      = 0
	CodeUndefined = 999999
)

type Definition[Req interface{}, Rsp interface{}] struct {
	Path    string
	Errors  map[int]*ErrorDefinition
	Handler func(req *Req) (*Rsp, error)
}

func (d *Definition[Req, Rsp]) Define(router gin.IRouter) {
	router.POST(d.Path, func(context *gin.Context) {
		processRequest(context, d)
	})
}

func processRequest[Req interface{}, Rsp interface{}](context *gin.Context, definition *Definition[Req, Rsp]) {
	req := new(Req)
	if err := context.ShouldBind(req); err != nil {
		// TODO: log bad request error
		responseOtherError(context, err, http.StatusBadRequest)
	} else {
		if rsp, err := definition.Handler(req); err != nil {
			var nraError *Error
			if ok := errors.As(err, &nraError); ok {
				if errorDefinition, ok := definition.Errors[nraError.code]; ok {
					nraError.definition = errorDefinition
					// TODO: log errorDefinition.loggable is true
					// TODO: alarm errorDefinition.alarm is true
					responseError(context, err, nraError)
				} else {
					// TODO: log undefined error
					responseError(context, err, nraError)
				}
			} else {
				// TODO: log unknown error
				responseOtherError(context, err, http.StatusInternalServerError)
			}
		} else {
			responseSuccess(context, rsp)
		}
	}
}

func responseSuccess[Rsp interface{}](context *gin.Context, rsp *Rsp) {
	context.Header(GlobalConfig.ProtocolFields.Version, strconv.Itoa(VersionCurrent))
	context.JSON(http.StatusOK, rsp)
}

func responseError(context *gin.Context, error error, nraError *Error) {
	service := GlobalConfig.Service
	endpoint := getEndpoint(context)
	code := nraError.GetCode()

	context.Header(GlobalConfig.ProtocolFields.Version, strconv.Itoa(VersionCurrent))
	context.Header(GlobalConfig.ProtocolFields.Code, strconv.Itoa(code))
	context.Header(GlobalConfig.ProtocolFields.Service, service)
	context.Header(GlobalConfig.ProtocolFields.Endpoint, endpoint)

	if nraError.HasNotice() {
		context.Header(GlobalConfig.ProtocolFields.Notice, url.QueryEscape(nraError.GetNotice()))
	}

	if nraError.HasRetry() {
		context.Header(GlobalConfig.ProtocolFields.Retry, "1")
	}

	var reason string
	if GlobalConfig.ExposeErrorReason {
		reason = nraError.GetReason()
	}

	var traces []*ErrorTraceDto = nil
	if GlobalConfig.ExposeErrorTraces {
		traces = getTraces(service, endpoint, error)
	}

	responseErrorBody(context, http.StatusOK, code, service, endpoint, reason, traces)
}

func responseOtherError(context *gin.Context, error error, status int) {
	service := GlobalConfig.Service
	endpoint := getEndpoint(context)

	var reason string
	if GlobalConfig.ExposeErrorReason {
		reason = error.Error()
	}

	var traces []*ErrorTraceDto = nil
	if GlobalConfig.ExposeErrorTraces {
		traces = getTraces(service, endpoint, error)
	}

	responseErrorBody(context, status, CodeNone, service, endpoint, reason, traces)
}

func responseErrorBody(context *gin.Context, status int, code int, service string, endpoint string, reason string, traces []*ErrorTraceDto) {
	if reason != "" || traces != nil {
		context.JSON(status, &ErrorRsp{
			Code:     code,
			Service:  service,
			Endpoint: endpoint,
			Reason:   reason,
			Traces:   traces,
		})
	} else {
		context.Status(status)
	}
}
