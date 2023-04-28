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

func NewTeapotResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://datatracker.ietf.org/doc/html/rfc9110#name-418-unused",
		Title:  "I'm a teapot",
		Status: http.StatusTeapot,
		Detail: detail,
		Log:    logError,
	}
}

func NewMisdirectedRequestResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://datatracker.ietf.org/doc/html/rfc9110#name-421-misdirected-request",
		Title:  "Misdirected Request",
		Status: http.StatusMisdirectedRequest,
		Detail: detail,
		Log:    logError,
	}
}

func NewUnprocessableContentResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://datatracker.ietf.org/doc/html/rfc9110#name-422-unprocessable-content",
		Title:  "Unprocessable Content",
		Status: http.StatusUnprocessableEntity,
		Detail: detail,
		Log:    logError,
	}
}

func NewLockedResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://datatracker.ietf.org/doc/html/rfc4918#section-11.3",
		Title:  "Locked",
		Status: http.StatusLocked,
		Detail: detail,
		Log:    logError,
	}
}

func NewFailedDependencyResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://datatracker.ietf.org/doc/html/rfc4918#section-11.4",
		Title:  "Failed Dependency",
		Status: http.StatusFailedDependency,
		Detail: detail,
		Log:    logError,
	}
}

func NewTooEarlyResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://datatracker.ietf.org/doc/html/rfc8470#section-5.2",
		Title:  "Too Early",
		Status: http.StatusTooEarly,
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

func NewPreconditionRequiredResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://datatracker.ietf.org/doc/html/rfc6585#section-3",
		Title:  "Precondition Required",
		Status: http.StatusPreconditionRequired,
		Detail: detail,
		Log:    logError,
	}
}

func NewTooManyRequestsResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://datatracker.ietf.org/doc/html/rfc6585#section-4",
		Title:  "Too Many Requests",
		Status: http.StatusTooManyRequests,
		Detail: detail,
		Log:    logError,
	}
}

func NewRequestHeaderFieldsTooLargeResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://datatracker.ietf.org/doc/html/rfc6585#section-5",
		Title:  "Too Many Requests",
		Status: http.StatusRequestHeaderFieldsTooLarge,
		Detail: detail,
		Log:    logError,
	}
}

func NewUnavailableForLegalReasonsResponse(detail string) *Problem {
	return &Problem{
		Type:   "https://www.rfc-editor.org/rfc/rfc7725.html#section-3",
		Title:  "Unavailable For Legal Reasons",
		Status: http.StatusUnavailableForLegalReasons,
		Detail: detail,
		Log:    logError,
	}
}
