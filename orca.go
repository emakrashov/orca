package main

import (
	"github.com/emakrashov/orca/server"
	"github.com/emakrashov/orca/storage"
)

const helpMessage = `
Usage:
        orca [command] [flags]

Available commands:
				server                    - starts gRPC server
				create_store [store_name] - create new store
				delete_store
				import_store store_name store_path
`

// orca server
// orca create storage_name
func main() {
	storage := storage.CreateStorage("/tmp/data1234")
	defer storage.CloseStorage()
	server.LaunchGrpc(&storage)
}
