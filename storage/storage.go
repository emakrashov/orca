package storage

import (
	"os"
	"sync"
	"sync/atomic"
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
	mutex *sync.Mutex
}

// CreateStorage creates an empty storage.
func CreateStorage(filename string) Storage {
	f, err := os.Create(filename)
	check(err)
	return Storage{file: f, size: 0, store: sync.Map{}, mutex: &sync.Mutex{}}
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
	fullEntity := encodeRecord(key, string(value))
	storage.mutex.Lock()
	bytesWritten := addBlock(storage.file, fullEntity)
	offset := int64(storage.size) + int64(len(key)) + 8
	atomic.AddInt64(&storage.size, int64(bytesWritten))
	storage.mutex.Unlock()
	storage.store.Store(key, Coords{offset: offset, len: len(value)})
}

// GetValue gets value from the store
func (storage *Storage) GetValue(key string) ([]byte, bool) {
	file := storage.file
	coordz, ok := storage.store.Load(key)
	if !ok {
		return make([]byte, 0), ok
	}
	coord := coordz.(Coords)
	storage.mutex.Lock()
	file.Seek(coord.offset, 0)
	buffer := make([]byte, coord.len)
	_, err := file.Read(buffer)
	storage.mutex.Unlock()
	check(err)
	return buffer, true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
