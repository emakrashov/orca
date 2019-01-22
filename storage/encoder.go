package storage

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func decodeEntity(reader io.Reader) (string, error) {
	// Read first 4 byte containing length of the entity
	buf := make([]byte, 4)
	if _, err := io.ReadFull(reader, buf); err != nil {
		return "", err
	}
	entitySize := binary.LittleEndian.Uint32(buf)
	entityBuf := make([]byte, entitySize)
	if _, err := io.ReadFull(reader, entityBuf); err != nil {
		return "", err
	}
	return string(entityBuf), nil
}

func decodeBlock(reader io.Reader) (string, string, error) {
	key, err := decodeEntity(reader)
	if err != nil {
		return "", "", err
	}
	value, err := decodeEntity(reader)
	if err != nil {
		return "", "", err
	}

	return key, value, nil
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
