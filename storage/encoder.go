package storage

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

func decodeEntity(reader *bytes.Reader) string {
	// Read first 4 byte containing length of the entity
	buf := make([]byte, 4)
	if _, err := io.ReadFull(reader, buf); err != nil {
		log.Fatal(err)
	}
	entitySize := binary.LittleEndian.Uint32(buf)
	entityBuf := make([]byte, entitySize)
	if _, err := io.ReadFull(reader, entityBuf); err != nil {
		log.Fatal(err)
	}
	return string(entityBuf)
}

func decodeBlock(block []byte) (string, string) {
	reader := bytes.NewReader(block)
	key := decodeEntity(reader)
	value := decodeEntity(reader)
	return key, value
}

func encodeBlock(key string, value string) []byte {
	buf := new(bytes.Buffer)
	encodeEntity(buf, key)
	encodeEntity(buf, value)
	return buf.Bytes()
}

func encodeEntity(buf *bytes.Buffer, entity string) {
	num := int32(len(entity))
	err := binary.Write(buf, binary.LittleEndian, num)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	binary.Write(buf, binary.LittleEndian, []byte(entity))
}
