package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"

	"github.com/emakrashov/orca/storage"
)

func TestTcpClient(t *testing.T) {
	storage := storage.CreateStorage("/tmp/dat4")
	defer storage.CloseStorage()

	server, client := net.Pipe()
	go func() {
		handleConn(server, storage)
		server.Close()
	}()

	client.Write([]byte("add orca whale\n"))
	client.Write([]byte("add a abracadabra\n"))
	client.Write([]byte("get orca\n"))

	message, _ := bufio.NewReader(client).ReadString('\n')
	message = strings.TrimSpace(message)
	if message != "whale" {
		t.Fatal("message: ", message)
	}
	fmt.Println("OK: message: ", message)
	client.Write([]byte("exit\n"))
	client.Close()
}

// func TestBlocks(t *testing.T) {
// 	os.Remove("/tmp/dat3")
// 	f, err := os.Create("/tmp/dat3")
// 	check(err)
// 	defer f.Close()

// 	data := []byte("test:a")
// 	addBlock(*f, data)
// }
