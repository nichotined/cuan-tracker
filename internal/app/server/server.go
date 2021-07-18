package server

import (
	"context"
	"cuan-tracker/internal/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func Start() {
	startServer(router())
}

func startServer(handler http.Handler) {
	stop := make(chan os.Signal, 2)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("server.startServer", "server started", nil)
	server := &http.Server{
		Addr:    "",
		Handler: handler,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-stop
	server.Shutdown(context.Background())
	logger.Info("server.startServer", "server stopped", nil)
}

func router() *mux.Router {
	router := mux.NewRouter()
	// router.HandleFunc("/ping")

	return router
}
