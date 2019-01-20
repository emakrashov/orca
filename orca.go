package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/emakrashov/orca/storage"
	"rsc.io/quote"
)

func handleConn(conn net.Conn, storage storage.Storage) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = strings.TrimSpace(message)
		line := strings.Split(message, " ")
		if len(line) >= 2 {
			command, arg := line[0], line[1]
			if (command == "add") && (line[2] != "") {
				storage.SetValue(arg, []byte(line[2]))
			}
			if command == "get" {
				res := storage.GetValue(strings.TrimSpace(arg))
				conn.Write([]byte("" + string(res) + "\n"))
			}
		}
		if (len(line) == 1) && (line[0] == "exit") {
			fmt.Println("Exiting...")
			conn.Close()
			break
		}
	}
}

func launch() {
	storage := storage.CreateStorage("/tmp/dat2")
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
		go handleConn(conn, storage)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println(quote.Hello())

	launch()
}
