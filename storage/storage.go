package storage

import (
	"os"
)

// Coords is offset and length of the value in the database file
type Coords struct {
	offset int64
	len    int
}

// CacheIndex stores in-memory map of keys to offset and length
type CacheIndex struct {
	size  int64
	store map[string]Coords
}

// Storage is main data structure
type Storage struct {
	file       *os.File
	cacheIndex CacheIndex
}

// CreateStorage creates an empty storage.
func CreateStorage(filename string) Storage {
	f, err := os.Create(filename)
	check(err)
	store := map[string]Coords{}
	cacheIndex := CacheIndex{size: 0, store: store}
	return Storage{file: f, cacheIndex: cacheIndex}
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
	cache := storage.cacheIndex.store
	size := storage.cacheIndex.size
	bytesWritten := addBlock(storage.file, value)
	cache[key] = Coords{offset: size, len: bytesWritten}
	size = int64(size + int64(bytesWritten))
	storage.cacheIndex = CacheIndex{size: size, store: cache}
}

// GetValue gets value from the store
func (storage *Storage) GetValue(key string) []byte {
	file := storage.file
	cache := storage.cacheIndex.store
	coord := cache[key]
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
