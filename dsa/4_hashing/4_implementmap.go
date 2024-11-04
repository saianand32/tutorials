package main

import (
	"fmt"
	"hash/fnv"
)

// Entity represents a key-value pair
type Entity struct {
	key   string
	value string
}

// HashMapFinal represents a simple hash map structure
type HashMapFinal struct {
	list [][]Entity // Array of lists for handling collisions
	size int        // Current size of the map
	lf   float32    // Load factor threshold for resizing
}

// NewHashMapFinal initializes a new hash map
func NewHashMapFinal() *HashMapFinal {
	list := make([][]Entity, 10) // Start with a list of 10 empty buckets
	return &HashMapFinal{list: list, size: 0, lf: 0.5}
}

// hash computes the hash index for a given key
func (h *HashMapFinal) hash(key string) int {
	hasher := fnv.New32a()                           // Create a new FNV-1a hasher
	hasher.Write([]byte(key))                        // Write the key to the hasher
	return int(hasher.Sum32() % uint32(len(h.list))) // Calculate hash index
}

// Put adds or updates a key-value pair in the hash map
func (h *HashMapFinal) Put(key, value string) {
	hashIndex := h.hash(key) // Compute hash index for the key
	entities := h.list[hashIndex]

	// Check if the key already exists
	for i, entity := range entities {
		if entity.key == key {
			h.list[hashIndex][i].value = value // Update value if found
			return
		}
	}

	// Check if rehashing is needed based on the load factor
	if float32(h.size)/float32(len(h.list)) > h.lf {
		h.rehash()
	}

	// Add new key-value pair
	h.list[hashIndex] = append(h.list[hashIndex], Entity{key: key, value: value})
	h.size++
}

// rehash doubles the size of the hash map and rehashes all entries
func (h *HashMapFinal) rehash() {
	fmt.Println("We are now rehashing!")
	old := h.list
	h.list = make([][]Entity, len(old)*2) // Create a new larger list
	h.size = 0                            // Reset size

	// Reinsert old entries into the new list
	for _, entities := range old {
		for _, entity := range entities {
			h.Put(entity.key, entity.value) // Use Put to handle rehashing
		}
	}
}

// Get retrieves the value associated with a key
func (h *HashMapFinal) Get(key string) string {
	hashIndex := h.hash(key) // Compute hash index
	entities := h.list[hashIndex]
	for _, entity := range entities {
		if entity.key == key {
			return entity.value // Return value if key is found
		}
	}
	return "" // Return empty string if key is not found
}

// Remove deletes the entry associated with a key
func (h *HashMapFinal) Remove(key string) {
	hashIndex := h.hash(key) // Compute hash index
	entities := h.list[hashIndex]

	// Find the target entity to remove
	for i, entity := range entities {
		if entity.key == key {
			h.list[hashIndex] = append(entities[:i], entities[i+1:]...) // Remove entity
			h.size--                                                    // Decrement size
			return
		}
	}
}

// ContainsKey checks if a key exists in the map
func (h *HashMapFinal) ContainsKey(key string) bool {
	return h.Get(key) != "" // Returns true if key exists
}

// String creates a string representation of the hash map
func (h *HashMapFinal) String() string {
	result := "{"
	for _, entities := range h.list {
		for _, entity := range entities {
			result += fmt.Sprintf("%s = %s, ", entity.key, entity.value)
		}
	}
	result += "}"
	return result
}

func main() {
	hashMap := NewHashMapFinal()     // Create a new hash map
	hashMap.Put("key1", "value1")    // Add a key-value pair
	hashMap.Put("key2", "value2")    // Add another pair
	fmt.Println(hashMap)             // Display the current state of the hash map
	fmt.Println(hashMap.Get("key1")) // Retrieve value for 'key1'
	hashMap.Remove("key1")           // Remove entry for 'key1'
	fmt.Println(hashMap)             // Display the state after removal
}
