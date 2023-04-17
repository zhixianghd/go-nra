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

type Error struct {
	code       int
	args       []any
	definition *ErrorDefinition
}

func (e *Error) Error() string {
	return fmt.Sprintf("nra.Error: code=(%d) args=(%+v)", e.code, e.args)
}

func (e *Error) GetCode() int {
	if e.definition != nil {
		return e.code
	}
	return CodeUndefined
}

func (e *Error) GetReason() string {
	if e.definition != nil {
		return fmt.Sprintf(e.definition.Reason, e.args...)
	}
	return fmt.Sprintf("undefined " + e.Error())
}

func (e *Error) HasNotice() bool {
	return e.definition != nil && e.definition.Notice != ""
}

func (e *Error) GetNotice() string {
	if e.HasNotice() {
		return e.definition.Notice
	}
	return ""
}

func (e *Error) HasRetry() bool {
	return e.definition != nil && e.definition.Retry
}

type ErrorCreator byte

func (c ErrorCreator) Text(text string, cause error) error {
	return errors.Join(errors.New(text), cause)
}

func (c ErrorCreator) Code(code int, cause error, args ...any) error {
	return errors.Join(&Error{
		code: code,
		args: args,
	}, cause)
}

var Errors = ErrorCreator(0)
