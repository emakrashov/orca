package storage

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestStorage(t *testing.T) {
	storage := CreateStorage("/tmp/dat4")
	defer storage.CloseStorage()

	storage.SetValue("a", []byte("a-value"))
	storage.SetValue("b", []byte("abracadabra"))
	storage.SetValue("c", []byte("jupiter"))

	a, _ := storage.GetValue("a")
	if string(a) != "a-value" {
		t.Fatal("Incorrect value fetched", string(a))
	}

	a, _ = storage.GetValue("b")
	if string(a) != "abracadabra" {
		t.Fatal("Incorrect value fetched", string(a))
	}

	a, _ = storage.GetValue("c")
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
	storage := CreateStorage("/tmp/dat10")
	var wg sync.WaitGroup
	c := 100
	assertMap := sync.Map{}

	wg.Add(c)
	for w := 1; w <= c; w++ {
		x := w
		go func() {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", x)
			value := []byte(RandStringRunes())
			assertMap.Store(key, value)
			storage.SetValue(key, value)
		}()
	}

	wg.Wait()

	var wg2 sync.WaitGroup
	wg2.Add(c)

	for w := 1; w <= c; w++ {
		x := w
		go func() {
			defer wg2.Done()
			key := fmt.Sprintf("key-%d", x)
			expectedValue, ok := assertMap.Load(key)
			expectedValueS := string(expectedValue.([]byte))
			actualValue, _ := storage.GetValue(key)
			if !ok || expectedValueS != string(actualValue) {
				t.Fatal(x, string(actualValue), "!=", expectedValueS)
			}
		}()
	}
	wg2.Wait()

}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes() string {
	n := rand.Intn(10) + 3
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
