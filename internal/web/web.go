package web

import (
	"net/http"

	"github.com/charlieroth/alexandria/internal/data"
)

type Config struct {
	Host string
	Port string
}

func NewServer(config *Config, db *data.JsonMutexDB) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, config, db)
	var handler http.Handler = mux
	return handler
}
