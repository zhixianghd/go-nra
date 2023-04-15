package nra

import (
	"errors"
	"fmt"
)

type ErrorDefinition struct {
	Reason   string
	Notice   string
	Retry    bool
	Alarm    bool
	Loggable bool
}

type BizError struct {
	code       int
	args       []any
	definition *ErrorDefinition
}

func (e *BizError) Error() string {
	return fmt.Sprintf("nra error: code=(%d) args=(%+v)", e.code, e.args)
}

func (e *BizError) GetCode() int {
	if e.definition != nil {
		return e.code
	}
	return CodeUndefined
}

func (e *BizError) GetReason() string {
	if e.definition != nil {
		return fmt.Sprintf(e.definition.Reason, e.args...)
	}
	return fmt.Sprintf("undefined " + e.Error())
}

func (e *BizError) HasNotice() bool {
	return e.definition != nil && e.definition.Notice != ""
}

func (e *BizError) GetNotice() string {
	if e.HasNotice() {
		return e.definition.Notice
	}
	return ""
}

func (e *BizError) HasRetry() bool {
	return e.definition != nil && e.definition.Retry
}

type ErrorCreator byte

func (c ErrorCreator) Text(text string, cause error) error {
	return errors.Join(errors.New(text), cause)
}

func (c ErrorCreator) Code(code int, cause error, args ...any) error {
	return errors.Join(&BizError{
		code: code,
		args: args,
	}, cause)
}

var Errors = ErrorCreator(0)
