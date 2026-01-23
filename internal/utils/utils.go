package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Payload map[string]interface{}

func WriteJSON(w http.ResponseWriter, statusCode int, p Payload) error {

	// This approach (MarshalIndent) lets you modify the JSON (add \n) before sending.  
	// 1st "" -> prefix, 2nd "" -> indent

	js, err := json.MarshalIndent(p,"","")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)

	return nil
}

func ReadIDFromParams(r *http.Request) (int64, error) {
	idParams := chi.URLParam(r, "id")
	if idParams == "" {
		return 0, 
		errors.New("id parameter not found")
	}

	id, err := strconv.ParseInt(idParams, 10, 64)
	if err != nil {
		return 0,
		errors.New("wrong id parameter type")
	}

	return id, nil
}