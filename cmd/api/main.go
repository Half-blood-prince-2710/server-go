package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type config struct {
	Addr string
	env  string
}

type application struct {
	cfg  config 
	logger *slog.Logger
}

func main() {
	// declare instance of config struct 
	var cfg config

	// read values for config from command-line flags
	flag.StringVar(&cfg.Addr,"Addr",":5000","Server Port Address")
	flag.StringVar(&cfg.env,"env","development","service Environment: development|staging|production Default: development")
	flag.Parse()


	// initializing the new strutured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout,nil))

	// declare instance of application struct and its configuration
	app := &application{
		cfg : cfg,
		logger: logger,
	}

	// declare the http server and its configurations
	srv := &http.Server{
		Addr: cfg.Addr,
		Handler: app.routes(),
		WriteTimeout: 30* time.Second,
		ReadTimeout: 30* time.Second,
		IdleTimeout: time.Minute,
		ErrorLog: slog.NewLogLogger(logger.Handler(),slog.LevelError),
	}


	// start the http server
	logger.Info("server is runing at ","Addr",cfg.Addr)
	err := srv.ListenAndServe()
	if err!=nil {
		logger.Error("Error running server")
		os.Exit(1)
	}
}
