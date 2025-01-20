package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	Addr string
}

type application struct {
	cfg  config 
	logger slog.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.Addr,"Addr",":5000","Server Port Address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout,nil))

	
	app := &application{
		cfg : cfg,
		logger: *logger,
	}

	srv := &http.Server{
		Addr: cfg.Addr,
		Handler: app.routes(),
	}
	
	logger.Info("server is runing at ","Addr",cfg.Addr)
	err := srv.ListenAndServe()
	if err!=nil {
		logger.Error("Error running server")
	}
}
