package main

import "fmt"

const tableSize = 16

type HashTable struct {
	data [tableSize]map[string]int
}

func (ht *HashTable) hash(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}
	return sum % tableSize
}

func (ht *HashTable) Set(key string, value int) {
	index := ht.hash(key)
	if ht.data[index] == nil {
		ht.data[index] = make(map[string]int)
	}
	ht.data[index][key] = value
}

func (ht *HashTable) Get(key string) (int, bool) {
	index := ht.hash(key)
	if ht.data[index] == nil {
		return 0, false
	}
	value, exists := ht.data[index][key]
	return value, exists
}

func main() {
	ht := &HashTable{}
	ht.Set("key1", 1)
	ht.Set("key2", 2)
	ht.Set("key3", 3)
	fmt.Println(ht.Get("key1"))
	fmt.Println(ht.Get("key2"))
	fmt.Println(ht.Get("key3"))
	fmt.Println(ht.Get("key4"))
}

