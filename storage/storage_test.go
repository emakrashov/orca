package storage

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestStorage(t *testing.T) {
	storage := CreateStorage("/tmp/dat4")
	defer storage.CloseStorage()

	storage.AddValue("a", []byte("a-value"))
	storage.AddValue("b", []byte("abracadabra"))
	storage.AddValue("c", []byte("jupiter"))

	a := storage.ReadValue("a")
	if string(a) != "a-value" {
		t.Fatal("Incorrect value fetched", a)
	}

	a = storage.ReadValue("b")
	if string(a) != "abracadabra" {
		t.Fatal("Incorrect value fetched", a)
	}

	a = storage.ReadValue("c")
	if string(a) != "jupiter" {
		t.Fatal("Incorrect value fetched", a)
	}
}

func printFileContent(file string) {
	dat, err := ioutil.ReadFile(file)
	check(err)
	fmt.Print(string(dat))
}

// func TestConcurrentStorage(t *testing.T) {
// 	os.Remove("/tmp/dat4")
// 	file, err := os.Create("/tmp/dat4")
// 	check(err)
// 	defer file.Close()

// 	cache := map[string]coords{}
// 	cache["limit"] = coords{offset: 0}

// 	for w := 1; w <= 10; w++ {
// 		x := w
// 		go func() {
// 			key := fmt.Sprintf("k%d", x)
// 			value := fmt.Sprintf("[value%d]", x)
// 			addValue(*file, cache, key, []byte(value))
// 		}()
// 	}

// 	printFileContent("/tmp/dat4")

// 	// addValue(*file, cache, "a", []byte("a-value"))
// 	// addValue(*file, cache, "b", []byte("abracadabra"))
// 	// addValue(*file, cache, "c", []byte("jupiter"))

// }
