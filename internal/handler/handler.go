package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/terryluciano/templ-test/internal/views/pages"
)

const maxJSONBodyBytes int64 = 1 << 20

type httpError struct {
	Status  int
	Message string
	Err     error
}

func (e *httpError) Error() string {
	if e.Err == nil {
		return e.Message
	}
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxJSONBodyBytes)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// Decode first JSON value.
	if err := dec.Decode(dst); err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError

		switch {
		case errors.Is(err, io.EOF):
			return &httpError{Status: http.StatusBadRequest, Message: "request body must not be empty", Err: err}
		case errors.As(err, &syntaxErr):
			return &httpError{Status: http.StatusBadRequest, Message: "malformed json", Err: err}
		case errors.As(err, &unmarshalTypeErr):
			return &httpError{Status: http.StatusBadRequest, Message: "invalid json type for a field", Err: err}
		case errors.Is(err, http.ErrBodyReadAfterClose):
			return &httpError{Status: http.StatusBadRequest, Message: "invalid request body", Err: err}
		default:
			// This includes "unknown field ..." from DisallowUnknownFields(),
			// and "http: request body too large" from MaxBytesReader.
			return &httpError{Status: http.StatusBadRequest, Message: "invalid json", Err: err}
		}
	}

	// Ensure there's only a single JSON object.
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return &httpError{Status: http.StatusBadRequest, Message: "request body must contain a single json object", Err: err}
	}

	return nil

}

func respondJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	name := "John Doe"

	ip := r.RemoteAddr

	pages.Home(name, ip).Render(r.Context(), w)
}
