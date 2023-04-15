package nra

import (
	"errors"
	"github.com/gin-gonic/gin"
	"testing"
)

type TestReq struct {
	Test string `json:"test" binding:"required"`
}

type TestRsp struct {
	Test string `json:"test,omitempty"`
}

func TestNra(t *testing.T) {
	GlobalConfig.ExposeErrorReason = true
	GlobalConfig.ExposeErrorTraces = true

	testApi := &Definition[TestReq, TestRsp]{
		Path: "/test",
		Errors: map[int]*ErrorDefinition{
			1000: {
				Reason: "error reason: %s %s %#v",
				Notice: "eng test / 测试中文通知",
			},
		},
		Handler: func(req *TestReq) (*TestRsp, error) {
			err := errors.New("original error")
			if req.Test == "1" {
				return &TestRsp{Test: "test"}, nil
			} else if req.Test == "2" {
				return nil, Errors.Code(1001, nil, "arg1", "arg2", TestRsp{Test: "test"})
			} else if req.Test == "3" {
				return nil, Errors.Code(1000, err, "arg1", "arg2", TestRsp{Test: "test"})
			}
			return nil, Errors.Text("fallback", err)
		},
	}

	engine := gin.Default()
	testApi.Define(engine)
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
