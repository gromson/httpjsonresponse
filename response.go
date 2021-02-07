package http_json_response

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type JsonSuccess struct {
	Result interface{}
	Log    Log
}

func NewSuccessResponse(result interface{}) JsonSuccess {
	return JsonSuccess{
		Result: result,
		Log:    logError,
	}
}

func (r JsonSuccess) Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	serialized, err := json.Marshal(r.Result)
	if err != nil {
		r.Log("error while trying to serialize a response", err, r.Result)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(serialized); err != nil {
		r.Log("error while trying to write json response body", err, string(serialized))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type Problem struct {
	Type   string      `json:"type"`
	Title  string      `json:"title"`
	Status int         `json:"status"`
	Detail interface{} `json:"details"`
	Log    Log         `json:"-"`
}

func (r *Problem) Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")

	serialized, err := json.Marshal(r)
	if err != nil {
		r.Log("error while trying to serialize a Problem response", err, r)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(r.Status)

	if _, err := w.Write(serialized); err != nil {
		r.Log("error while trying to write Problem response body", err, string(serialized))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewProblemResponse(title string, detail interface{}) *Problem {
	if errs, ok := detail.([]error); ok {
		detail = errorSliceToStringSlice(errs)
	}

	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.5.1",
		Title:  title,
		Status: http.StatusBadRequest,
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

func NewInternalError() *Problem {
	return &Problem{
		Type:   "https://tools.ietf.org/html/rfc7231#section-6.6.1",
		Title:  "Internal Server Error",
		Status: http.StatusInternalServerError,
		Detail: nil,
		Log:    logError,
	}
}

func errorSliceToStringSlice(data []error) []string {
	ss := make([]string, 0, len(data))

	for _, e := range data {
		ss = append(ss, e.Error())
	}

	return ss
}

type Log func(title string, err error, additional interface{})

func logError(title string, err error, additional interface{}) {
	log.WithError(err).WithField("additional", additional).Error(title)
}
