package middleware

import (
	"github.com/zhenyanesterkova/nepblog/internal/app/logger"
)

type MiddlewareStruct struct {
	Logger   logger.LogrusLogger
	respData *responseDataWriter
}

func NewMiddlewareStruct(log logger.LogrusLogger) MiddlewareStruct {
	responseData := &responseData{
		status: 0,
		size:   0,
	}

	lw := responseDataWriter{
		responseData: responseData,
	}

	return MiddlewareStruct{
		Logger:   log,
		respData: &lw,
	}
}
