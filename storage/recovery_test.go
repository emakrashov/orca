package storage

import (
	"testing"
)

func TestRecovery(t *testing.T) {
	storage := CreateStorage("/tmp/dat4")

	storage.SetValue("a", []byte("abracadabra"))
	storage.SetValue("bfs", []byte("breadth-first search"))
	storage.SetValue("co", []byte("coca-cola"))

	storage.CloseStorage()

	s := Recover("/tmp/dat4")

	value := string(s.GetValue("a"))
	if value != "abracadabra" {
		t.Fatal("Incorrect value fetched", value)
	}

	value = string(s.GetValue("bfs"))
	if value != "breadth-first search" {
		t.Fatal("Incorrect value fetched", value)
	}

	value = string(s.GetValue("co"))
	if value != "coca-cola" {
		t.Fatal("Incorrect value fetched", value)
	}
}
