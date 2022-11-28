package httpjsonresponse

import "net/http"

func NewInternalServerErrorResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.6.1",
		Title:  "Internal Server Error",
		Status: http.StatusInternalServerError,
		Detail: detail,
		Log:    logError,
	}
}

func NewNotImplementedResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.6.2",
		Title:  "Not Implemented",
		Status: http.StatusNotImplemented,
		Detail: detail,
		Log:    logError,
	}
}

func NewBadGatewayResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.6.3",
		Title:  "Bad Gateway",
		Status: http.StatusBadGateway,
		Detail: detail,
		Log:    logError,
	}
}

func NewServiceUnavailableResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.6.4",
		Title:  "Service Unavailable",
		Status: http.StatusServiceUnavailable,
		Detail: detail,
		Log:    logError,
	}
}

func NewGatewayTimeoutResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.6.5",
		Title:  "Gateway Timeout",
		Status: http.StatusGatewayTimeout,
		Detail: detail,
		Log:    logError,
	}
}

func NewHTTPVersionNotSupportedResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.6.6",
		Title:  "HTTP Version Not Supported",
		Status: http.StatusHTTPVersionNotSupported,
		Detail: detail,
		Log:    logError,
	}
}
