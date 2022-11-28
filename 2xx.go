package httpjsonresponse

import "net/http"

func NewOkResponse(data interface{}) *Success {
	return &Success{
		Status: http.StatusOK,
		Data:   data,
		Log:    logError,
	}
}

func NewCreatedResponse(data interface{}) *Success {
	return &Success{
		Status: http.StatusCreated,
		Data:   data,
		Log:    logError,
	}
}

func NewAcceptedResponse(data interface{}) *Success {
	return &Success{
		Status: http.StatusAccepted,
		Data:   data,
		Log:    logError,
	}
}

func NewNonAuthoritativeInfoResponse(data interface{}) *Success {
	return &Success{
		Status: http.StatusNonAuthoritativeInfo,
		Data:   data,
		Log:    logError,
	}
}

func NewNoContentResponse() *Success {
	return &Success{
		Status: http.StatusNoContent,
		Data:   nil,
		Log:    logError,
	}
}

func NewResetContentResponse() *Success {
	return &Success{
		Status: http.StatusResetContent,
		Data:   nil,
		Log:    logError,
	}
}
