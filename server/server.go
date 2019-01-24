package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/emakrashov/orca/storage"
)

// either "ok sth" or "error sth"
func parseCommand(storage *storage.Storage, line string) (string, bool) {
	terms := strings.Split(line, " ")
	if len(terms) == 0 {
		return "", true
	}
	switch terms[0] {
	case "get":
		if len(terms) >= 2 {
			arg := strings.TrimSpace(terms[1])
			res, ok := storage.GetValue(arg)
			if ok {
				return string(res), ok
			}
			return "Key not found", ok
		}
		return "get is missing an argument", false

		// case "set":
		// if len(terms) >= 3 {
		// key := strings.TrimSpace(terms[1])
		// args := strings.TrimSpace(terms[1])
		// rest = terms[2:len(terms)]
		// }

		// return "Command not found", true

	default:
		return "Command not found", true
	}

}

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
