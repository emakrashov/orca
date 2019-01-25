package server

import (
	"testing"
)

func TestTcpClient(t *testing.T) {
	// storage := storage.CreateStorage("/tmp/dat4")
	// defer storage.CloseStorage()

	// s, client := net.Pipe()
	// go func() {
	// 	handleConn(s, &storage)
	// 	s.Close()
	// }()

	// client.Write([]byte("set orca whale\n"))
	// client.Write([]byte("set a abracadabra\n"))
	// client.Write([]byte("get orca\n"))

	// message, _ := bufio.NewReader(client).ReadString('\n')
	// message = strings.TrimSpace(message)

	// if message != "whale" {
	// 	t.Fatal("message: ", message)
	// }
	// fmt.Println("OK: message: ", message)

	// client.Write([]byte("exit\n"))
	// client.Close()
}
