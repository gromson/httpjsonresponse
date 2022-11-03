package http_json_response

import "net/http"

func NewOkResponse(data interface{}) *JsonSuccess {
	return &JsonSuccess{
		Status: http.StatusOK,
		Data:   data,
		Log:    logError,
	}
}

func NewCreatedResponse(data interface{}) *JsonSuccess {
	return &JsonSuccess{
		Status: http.StatusCreated,
		Data:   data,
		Log:    logError,
	}
}

func NewAcceptedResponse(data interface{}) *JsonSuccess {
	return &JsonSuccess{
		Status: http.StatusAccepted,
		Data:   data,
		Log:    logError,
	}
}

func NewNonAuthoritativeInfoResponse(data interface{}) *JsonSuccess {
	return &JsonSuccess{
		Status: http.StatusNonAuthoritativeInfo,
		Data:   data,
		Log:    logError,
	}
}

func NewNoContentResponse() *JsonSuccess {
	return &JsonSuccess{
		Status: http.StatusNoContent,
		Data:   nil,
		Log:    logError,
	}
}

func NewResetContentResponse() *JsonSuccess {
	return &JsonSuccess{
		Status: http.StatusResetContent,
		Data:   nil,
		Log:    logError,
	}
}
