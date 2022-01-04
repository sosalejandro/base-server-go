package homepage

import (
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

const message = "Hello World from Go"

type Handlers struct {
	logger *log.Logger
	db     *sqlx.DB
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.db.ExecContext(r.Context(), "")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Since(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

func NewHandlers(logger *log.Logger, db *sqlx.DB) *Handlers {
	return &Handlers{
		logger: logger,
		db:     db,
	}
}
