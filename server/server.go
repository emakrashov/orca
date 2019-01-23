package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/emakrashov/orca/storage"
)

func handleConn(conn net.Conn, storage *storage.Storage) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = strings.TrimSpace(message)
		line := strings.Split(message, " ")
		if len(line) >= 2 {
			command, arg := line[0], line[1]
			if (command == "set") && (line[2] != "") {
				storage.SetValue(strings.TrimSpace(arg), []byte(line[2]))
			}
			if command == "get" {
				res, ok := storage.GetValue(strings.TrimSpace(arg))

				if ok {
					conn.Write([]byte("" + string(res) + "\n"))
				} else {
					conn.Write([]byte("not found"))
				}
			}
		}
		if (len(line) == 1) && (line[0] == "exit") {
			fmt.Println("Exiting...")
			conn.Close()
			break
		}
	}
}

// Launch the server
func Launch(databaseFile string) {
	storage := storage.CreateStorage(databaseFile)
	defer storage.CloseStorage()

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081")
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConn(conn, &storage)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
