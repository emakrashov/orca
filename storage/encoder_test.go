package storage

import (
	"bytes"
	"testing"
)

func TestEncodeBlock(t *testing.T) {
	block := bytes.NewReader(encodeRecord("apple", "APPL"))
	key, value, _ := decodeBlock(block)
	if key != "apple" || value != "APPL" {
		t.Fatal("Incorrect value fetched", key, value)
	}

	block = bytes.NewReader(encodeRecord("coca-cola", "CO"))
	key, value, _ = decodeBlock(block)
	if key != "coca-cola" || value != "CO" {
		t.Fatal("Incorrect value fetched", key, value)
	}

	block = bytes.NewReader(encodeRecord("netflix", "NFLX"))
	key, value, _ = decodeBlock(block)
	if key != "netflix" || value != "NFLX" {
		t.Fatal("Incorrect value fetched", key, value)
	}
}
