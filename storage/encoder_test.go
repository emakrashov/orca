package storage

import "testing"

func TestEncodeBlock(t *testing.T) {
	block := encodeBlock("apple", "APPL")
	key, value := decodeBlock(block)
	if key != "apple" || value != "APPL" {
		t.Fatal("Incorrect value fetched", key, value)
	}

	block = encodeBlock("coca-cola", "CO")
	key, value = decodeBlock(block)
	if key != "coca-cola" || value != "CO" {
		t.Fatal("Incorrect value fetched", key, value)
	}

	block = encodeBlock("netflix", "NFLX")
	key, value = decodeBlock(block)
	if key != "netflix" || value != "NFLX" {
		t.Fatal("Incorrect value fetched", key, value)
	}
}
