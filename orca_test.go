package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
)

func TestTcpClient(t *testing.T) {
	os.Remove("/tmp/dat3")
	f, err := os.Create("/tmp/dat3")
	check(err)
	server, client := net.Pipe()
	go func() {
		handleConn(server, *f)
		server.Close()
	}()

	client.Write([]byte("add a 5\n"))
	client.Write([]byte("add b 3\n"))
	client.Write([]byte("get a\n"))

	message, _ := bufio.NewReader(client).ReadString('\n')
	message = strings.TrimSpace(message)
	if message != "5" {
		t.Fatal("message: ", message)
	}
	fmt.Println("OK: message: ", message)
	client.Write([]byte("exit\n"))
	client.Close()
}

func TestBlocks(t *testing.T) {
	os.Remove("/tmp/dat3")
	f, err := os.Create("/tmp/dat3")
	check(err)
	defer f.Close()

	data := []byte("test:a")
	addBlock(*f, data)
}
