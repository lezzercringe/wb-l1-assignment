package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	keys   = []string{"key1", "key2", "key3", "key4"}
	values = []string{"value1", "value2", "value3", "value4"}
)

func writer(m *sync.Map) {
	for {
		key := keys[rand.Intn(len(keys))]
		value := values[rand.Intn(len(values))]

		m.Store(key, value)
	}
}

func main() {
	m := sync.Map{}

	go writer(&m)

	for {
		key := keys[rand.Intn(len(keys))]
		val, ok := m.Load(key)
		if !ok {
			fmt.Printf("Key %v - empty\n", key)
			continue
		}

		fmt.Printf("Key %v = %v\n", key, val)
	}
}
