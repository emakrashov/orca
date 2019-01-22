package storage

import (
	"os"
	"sync"
)

// Coords is offset and length of the value in the database file
type Coords struct {
	offset int64
	len    int
}

// Storage is main data structure
type Storage struct {
	file  *os.File
	size  int64
	store sync.Map
}

// CreateStorage creates an empty storage.
func CreateStorage(filename string) Storage {
	f, err := os.Create(filename)
	check(err)
	return Storage{file: f, size: 0, store: sync.Map{}}
}

// CloseStorage closes the corresponding file
func (storage *Storage) CloseStorage() {
	println("Close file")
	storage.file.Close()
}

func addBlock(file *os.File, data []byte) int {
	bytesWritten, err := file.Write(data)
	check(err)
	return bytesWritten
}

// SetValue sets the key with the given value in the store
func (storage *Storage) SetValue(key string, value []byte) {
	// string() ?
	fullEntity := encodeBlock(key, string(value))
	bytesWritten := addBlock(storage.file, fullEntity)
	offset := int64(storage.size) + int64(len(key)) + 8
	storage.store.Store(key, Coords{offset: offset, len: len(value)})
	storage.size = int64(storage.size + int64(bytesWritten))
}

// GetValue gets value from the store
func (storage *Storage) GetValue(key string) []byte {
	file := storage.file
	coordz, _ := storage.store.Load(key)
	coord := coordz.(Coords)
	file.Seek(coord.offset, 0)
	buffer := make([]byte, coord.len)
	_, err := file.Read(buffer)
	check(err)
	return buffer
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
