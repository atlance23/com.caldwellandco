package server

// Import
import (
	"context"
	"net/http"
	"log"
	"fmt"
	"path/filepath"
	"time"
)

const PORT = ":8080"

var srv *http.Server

// HTTP Handlers
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Set the path to the dist folder
	distPath := filepath.Join("..", "..", "frontend", "dist")

	// Handle requests by serving files from the dist folder
	http.ServeFile(w, r, filepath.Join(distPath, r.URL.Path[1:]))
}

// Server Management Function
func ManageSrv() {
	for {
		var userInput string
		fmt.Println("Server started on port " + PORT[1:])
		fmt.Println("Server management commands: {S}top | {R}eload")
		fmt.Print(">>> ")
		fmt.Scan(&userInput)

		switch userInput {
			case "S":
				Stop()
				return
			case "R":
				Reload()
			case "s":
				Stop()
			case "r":
				Reload()
			default:
				fmt.Println("Invalid Input! Please try again: \n")
		}
	}
}

// Exported Server Functions
func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	srv = &http.Server{
		Addr: PORT,
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()
}

// Stop Server Method
func Stop() {
	if srv != nil {
		fmt.Println("Server shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Server Shutdown Failed:%+v", err)
		} else {
			fmt.Println("Server stopped gracefully.")
			fmt.Println("")
		}
		srv = nil
	}
}

// Reload Method
func Reload() {
	fmt.Println("Reloading server...")
	Stop()
	Start()
}