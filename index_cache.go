package main

func createCache() cacheIndex {
	var store = map[string]coords{}
	var c = cacheIndex{size: 0, store: store}
	return c
}
