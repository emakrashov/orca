package storage

import "os"

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

func (storage Storage) CloseStorage() {
	storage.file.Close()
}

func addBlock(file *os.File, data []byte) int {
	bytesWritten, err := file.Write(data)
	check(err)
	return bytesWritten
}

func (storage Storage) AddValue(key string, value []byte) {
	cache := storage.cacheIndex.store
	offset := storage.cacheIndex.size
	bytesWritten := addBlock(storage.file, value)
	cache[key] = Coords{offset: offset, len: bytesWritten}
	offset = offset + int64(bytesWritten)
	storage.cacheIndex.size = offset
}

func (storage Storage) ReadValue(key string) []byte {
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
