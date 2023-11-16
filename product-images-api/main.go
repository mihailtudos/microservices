package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mihailtudos/product-images-api/files"
	"github.com/mihailtudos/product-images-api/handlers"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var (
	addr     string
	logLevel string
	basePath string
)

func init() {
	flag.StringVar(&addr, "BIND_ADDRESS", ":9090", "HTTP binging address")
	flag.StringVar(&logLevel, "LOG_LEVEL", "debug", "Log output level for the server [debug, info, trace]")
	flag.StringVar(&basePath, "BASE_PATH", "./filestorage", "Base path to save images")
}

func main() {
	flag.Parse()

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	l.Info("starting server", slog.String("BIND_ADDRESS", addr), slog.String("LOG_LEVEL", logLevel), slog.String("BASE_PATH", basePath))

	stor, err := files.NewLocal(basePath, 1024*1000*5)

	if err != nil {
		l.Error(fmt.Sprintf("Unable to create storage %s", err.Error()))
		os.Exit(1)
	}

	fh := handlers.NewFiles(stor, *l)

	sm := mux.NewRouter()

	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.UploadRest)
	ph.HandleFunc("/images/{id:[0-9]+}", fh.UploadMultiPart)

	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))))

	sm.Use(fh.CorsMiddleware)

	s := http.Server{
		Addr:         addr,
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err = s.ListenAndServe()
	l.Error("Couldn't start server " + err.Error())
	os.Exit(1)
}
