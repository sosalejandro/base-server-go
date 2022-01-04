package main

import (
	"log"
	"net/http"
	"os"
	"thirty/homepage"
	"thirty/server"

	"github.com/jmoiron/sqlx"
)

var (
	serverAddress = os.Getenv("SERVER_ADDR")
	CertFile      = os.Getenv("CERT_FILE")
	KeyFile       = os.Getenv("KEY_FILE")
)

func main() {
	logger := log.New(os.Stdout, "backend: ", log.LstdFlags|log.Lshortfile)

	db, err := sqlx.Open("postgres", "postgres://postgres:postgres")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	// D.I.
	h := homepage.NewHandlers(logger, db)

	mux := initiateAndRegisterRoutes(h)

	srv := server.New(mux, serverAddress)

	err = srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}

func initiateAndRegisterRoutes(h *homepage.Handlers) *http.ServeMux {
	mux := http.NewServeMux()
	h.SetupRoutes(mux)
	return mux
}
