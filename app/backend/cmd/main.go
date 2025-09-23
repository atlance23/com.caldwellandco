package main

import (
	"github.com/atlance23/com.caldwellandco/app/backend/internal/server"
)

// Entry Point
func main() {
	go server.Start()
	server.ManageSrv()
}