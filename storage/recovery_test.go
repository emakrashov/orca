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

	value, _ := s.GetValue("a")
	if string(value) != "abracadabra" {
		t.Fatal("Incorrect value fetched", value)
	}

	value, _ = s.GetValue("bfs")
	if string(value) != "breadth-first search" {
		t.Fatal("Incorrect value fetched", value)
	}

	value, _ = s.GetValue("co")
	if string(value) != "coca-cola" {
		t.Fatal("Incorrect value fetched", value)
	}
}
