package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/charlieroth/alexandria/internal/data"
)

func encode[T any](w http.ResponseWriter, statusCode int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func addRoutes(mux *http.ServeMux, config *Config, db *data.JsonMutexDB) {
	mux.Handle("GET /liveness", handleLiveness())
	mux.Handle("POST /document", handlePostDocument(db))
	mux.Handle("GET /document/{id}", handleGetDocument(db))
	mux.Handle("/", http.NotFoundHandler())
}

func handleGetDocument(db *data.JsonMutexDB) http.HandlerFunc {
	type response struct {
		DocumentId      string `json:"id,omitempty"`
		DocumentContent string `json:"content,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		db.BigLock.Lock()
		document, err := db.GetDocument(id)
		db.BigLock.Unlock()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		encode(w, http.StatusOK, response{
			DocumentId:      document.Id,
			DocumentContent: string(document.Content),
		})
	}
}

func handlePostDocument(db *data.JsonMutexDB) http.HandlerFunc {
	type request struct {
		Content string `json:"content"`
	}

	type response struct {
		DocumentId string `json:"id,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		decoded, err := decode[request](r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		db.BigLock.Lock()
		document := data.Document{
			Id:      "asdf",
			Content: []byte(decoded.Content),
		}
		err = db.AddDocument(document)
		db.BigLock.Unlock()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		encode(w, http.StatusCreated, response{document.Id})
	}
}

func handleLiveness() http.HandlerFunc {
	type response struct {
		Status     string `json:"status,omitempty"`
		Host       string `json:"host,omitempty"`
		GOMAXPROCS int    `json:"GOMAXPROCS,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		host, err := os.Hostname()
		if err != nil {
			host = "unavailable"
		}

		res := response{
			Status:     "up",
			Host:       host,
			GOMAXPROCS: runtime.GOMAXPROCS(0),
		}

		if err := encode(w, http.StatusOK, res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
