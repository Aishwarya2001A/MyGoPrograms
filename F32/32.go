package main

import (
	"fmt"
)

const hashTableSize = 10

type HashTableLinear struct {
	keys   [hashTableSize]int
	values [hashTableSize]interface{}
	size   int
}

func hashFunction(key int) int {
	return key % hashTableSize
}

func (ht *HashTableLinear) insert(key int, value interface{}) {
	index := hashFunction(key)

	for ht.keys[index] != 0 {
		index = (index + 1) % hashTableSize
	}

	ht.keys[index] = key
	ht.values[index] = value
	ht.size++
}

func (ht *HashTableLinear) get(key int) interface{} {
	index := hashFunction(key)

	for ht.keys[index] != key {
		index = (index + 1) % hashTableSize
	}

	return ht.values[index]
}

func main() {
	ht := HashTableLinear{}

	ht.insert(101, "Aish")
	ht.insert(205, "Swati")
	ht.insert(306, "Deepi")
	ht.insert(401, "Divi")
	ht.insert(502, "Chavi")

	fmt.Println("Name for key 205:", ht.get(205))
	fmt.Println("Name for key 401:", ht.get(401))
	fmt.Println("Name for key 306:", ht.get(306))

	hashTable := make(map[int]string)

	hashTable[1] = "hii"
	hashTable[2] = "hello"
	hashTable[3] = "tata"
	hashTable[4] = "bubye"

	fmt.Println("Value for key 1:", hashTable[1])
	fmt.Println("Value for key 3:", hashTable[3])

	if value, exists := hashTable[2]; exists {
		fmt.Println("Key 2 exists with value:", value)
	} else {
		fmt.Println("Key 2 does not exist in the hash table.")
	}

	delete(hashTable, 4)

	if _, exists := hashTable[4]; exists {
		fmt.Println("Key 4 still exists in the hash table.")
	} else {
		fmt.Println("Key 4 is successfully removed from the hash table.")
	}

}
