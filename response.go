package http_json_response

import (
	"bytes"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var jsonNull = []byte("null")

type JsonSuccess struct {
	Status int
	Data   interface{}
	Log    Log
}

func (r *JsonSuccess) Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	serialized, err := json.Marshal(r.Data)
	if err != nil {
		r.Log("error while trying to serialize a response", err, r.Data)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(r.Status)

	if bytes.Equal(serialized, jsonNull) {
		return
	}

	if _, err := w.Write(serialized); err != nil {
		r.Log("error while trying to write json response body", err, string(serialized))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type Problem struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
	Log    Log    `json:"-"`
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

func NewSuccessResponse(status int, result interface{}) *JsonSuccess {
	return &JsonSuccess{
		Status: status,
		Data:   result,
		Log:    logError,
	}
}

func NewProblemResponse(typ, title string, status int, detail string) *Problem {
	return &Problem{
		Type:   typ,
		Title:  title,
		Status: status,
		Detail: detail,
		Log:    logError,
	}
}

type Log func(title string, err error, additional interface{})

func logError(title string, err error, additional interface{}) {
	log.WithError(err).WithField("additional", additional).Error(title)
}
