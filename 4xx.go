package httpjsonresponse

import "net/http"

func NewBadRequestResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.1",
		Title:  "Bad Request",
		Status: http.StatusBadRequest,
		Detail: detail,
		Log:    logError,
	}
}

func NewUnauthorizedResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7235#section-3.1",
		Title:  "Unauthorized",
		Status: http.StatusUnauthorized,
		Detail: detail,
		Log:    logError,
	}
}

func NewForbiddenResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.3",
		Title:  "Forbidden",
		Status: http.StatusForbidden,
		Detail: detail,
		Log:    logError,
	}
}

func NewNotFoundResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.4",
		Title:  "Not Found",
		Status: http.StatusNotFound,
		Detail: detail,
		Log:    logError,
	}
}

func NewMethodNotAllowedResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.5",
		Title:  "Method Not Allowed",
		Status: http.StatusMethodNotAllowed,
		Detail: detail,
		Log:    logError,
	}
}

func NewNotAcceptableResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.6",
		Title:  "Not Acceptable",
		Status: http.StatusNotAcceptable,
		Detail: detail,
		Log:    logError,
	}
}

func NewRequestTimeoutResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.7",
		Title:  "Request Timeout",
		Status: http.StatusRequestTimeout,
		Detail: detail,
		Log:    logError,
	}
}

func NewConflictResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.8",
		Title:  "Conflict",
		Status: http.StatusConflict,
		Detail: detail,
		Log:    logError,
	}
}

func NewGoneResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.9",
		Title:  "Gone",
		Status: http.StatusGone,
		Detail: detail,
		Log:    logError,
	}
}

func NewLengthRequiredResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.10",
		Title:  "Length Required",
		Status: http.StatusLengthRequired,
		Detail: detail,
		Log:    logError,
	}
}

func NewRequestEntityTooLargeResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.11",
		Title:  "Request Entity Too Large",
		Status: http.StatusRequestEntityTooLarge,
		Detail: detail,
		Log:    logError,
	}
}

func NewRequestURITooLongResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.12",
		Title:  "Request URI Too Long",
		Status: http.StatusRequestURITooLong,
		Detail: detail,
		Log:    logError,
	}
}

func NewUnsupportedMediaTypeResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.13",
		Title:  "Unsupported Media Type",
		Status: http.StatusUnsupportedMediaType,
		Detail: detail,
		Log:    logError,
	}
}

func NewExpectationFailedResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.14",
		Title:  "Expectation Failed",
		Status: http.StatusExpectationFailed,
		Detail: detail,
		Log:    logError,
	}
}

func NewUpgradeRequiredResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.15",
		Title:  "Upgrade Required",
		Status: http.StatusUpgradeRequired,
		Detail: detail,
		Log:    logError,
	}
}
