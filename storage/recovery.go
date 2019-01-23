package storage

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Recover is used to recover from unplanned program crash or exit.
// It rebuilds the in-memory index map.
func Recover(filename string) *Storage {
	file, err := os.Open(filename)
	check(err)
	storage := Storage{file: file, size: 0, store: sync.Map{}, mutex: &sync.Mutex{}}

	for {
		k, v, err := decodeBlock(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return &storage
		}
		offset := int64(storage.size) + int64(len(k)) + 8
		bytesWritten := int64(len(k)) + 8 + int64(len(v))
		storage.store.Store(k, Coords{offset: offset, len: int(len(v))})
		storage.size = int64(storage.size + int64(bytesWritten))
	}

	return &storage
}
