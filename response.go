package httpjsonresponse

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Success struct {
	Status int
	Data   interface{}
	Log    Log
}

func (r *Success) Respond(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	serialized, err := json.Marshal(r.Data)
	if err != nil {
		r.Log("error while trying to serialize a response", err, r.Data)
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(r.Status)

	if bytes.Equal(serialized, []byte("null")) {
		return
	}

	if _, err := writer.Write(serialized); err != nil {
		r.Log("error while trying to write json response body", err, string(serialized))
		writer.WriteHeader(http.StatusInternalServerError)

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

func (r *Problem) Respond(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/problem+json; charset=utf-8")

	serialized, err := json.Marshal(r)
	if err != nil {
		r.Log("error while trying to serialize a Problem response", err, r)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(r.Status)

	if _, err := writer.Write(serialized); err != nil {
		r.Log("error while trying to write Problem response body", err, string(serialized))
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func NewSuccessResponse(status int, result interface{}) *Success {
	return &Success{
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
	log.Fatalf(`%s: %s (%v)`, title, err, additional)
}
