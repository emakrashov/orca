package storage

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestStorage(t *testing.T) {
	storage := CreateStorage("/tmp/dat4")
	defer storage.CloseStorage()

	storage.SetValue("a", []byte("a-value"))
	storage.SetValue("b", []byte("abracadabra"))
	storage.SetValue("c", []byte("jupiter"))

	a := storage.GetValue("a")
	if string(a) != "a-value" {
		t.Fatal("Incorrect value fetched", string(a))
	}

	a = storage.GetValue("b")
	if string(a) != "abracadabra" {
		t.Fatal("Incorrect value fetched", string(a), storage)
	}

	a = storage.GetValue("c")
	if string(a) != "jupiter" {
		t.Fatal("Incorrect value fetched", string(a))
	}
}

func printFileContent(file string) {
	dat, err := ioutil.ReadFile(file)
	check(err)
	fmt.Print(string(dat))
}

func TestConcurrentStorage(t *testing.T) {
	storage := CreateStorage("/tmp/dat4")

	for w := 1; w <= 1000; w++ {
		x := w
		go func() {
			key := fmt.Sprintf("k%d", x)
			value := []byte(fmt.Sprintf("[value%d]", x))
			storage.SetValue(key, value)
		}()
	}

	fmt.Scanln()
	time.Sleep(500000000)
	printFileContent("/tmp/dat4")
	println("end")
}
