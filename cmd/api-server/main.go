package main

import (
	"os"

	server "intelligent-investor/cmd/api-server/server"
)

func main() {
	cmd := server.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
