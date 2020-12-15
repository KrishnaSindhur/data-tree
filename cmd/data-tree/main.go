package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/KrishnaSindhur/data-tree/pkg/constants"
	"github.com/KrishnaSindhur/data-tree/pkg/handler"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

const (
	help = `
		data-tree-service has all endpoints related to  data-tree information.
		
		The following sub-commands are supported
		  1. serve - Run the application as a server.
		  3. help - Prints this help.
		  4. version - Prints the version of the application.
		
		You can pass the following configuration through environment variables.
		The configuration, unless stated otherwise, are mandatory.
	`
)

const (
	contextPath    = "/data-tree/v1"
	DataInsertPath = contextPath + "/insert"
	DataQueryPath  = contextPath + "/query"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("No arguments provided.")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch strings.ToLower(cmd) {
	case "serve":
		serve()
	case "help":
		fmt.Print(help)
	case "version":
		fmt.Println(constants.Version)
	default:
		fmt.Printf("Unknown command %q", cmd)
		os.Exit(1)
	}
}

func serve() {

	r := InitializeRoutes()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", constants.Port),
		Handler: r,
	}

	go func() {
		log.Print("Starting server")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Panic().Err(err).Msg("Failed to start server")
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-stop

	log.Print("Shutting the server down...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Err(err).Msg("Server did not shutdown gracefully")
	} else {
		log.Info().Msg("Server stopped")
	}
}

func InitializeRoutes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc(DataInsertPath, handler.Add()).Methods(http.MethodPost)
	router.HandleFunc(DataQueryPath, handler.Get()).Methods(http.MethodGet)
	return router
}
