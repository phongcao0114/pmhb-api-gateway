package models

import "time"

type (
	GraphQLPostBody struct {
		Query         string                 `json:"query"`
		Variables     map[string]interface{} `json:"variables"`
		OperationName string                 `json:"operationName"`
	}

	// KbankResponseHeader struct contains response header
	KbankResponseHeader struct {
		ResponseAppID string                   `json:"response_app_id"`
		ResponseDate  time.Time                `json:"response_datetime"`
		StatusCode    string                   `json:"status_code"`
		Errors        ResponseErrorKbankHeader `json:"error"`
	}

	// ResponseErrorKbankHeader struct contains response error header
	ResponseErrorKbankHeader struct {
		ErrorCode string `json:"error_code"`
		ErrorDesc string `json:"error_desc"`
	}
)
