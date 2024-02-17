package web

import (
	"context"
	"net/http"
)

func Respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	if err := Encode(w, statusCode, data); err != nil {
		return err
	}

	return nil
}
