package main

import (
	"fmt"
	"sort"
	"testing"
)

	var mapping = map[string]int{
			"Shailendra": 100,
			"Soorya":  150,
			"Aman": 500,
			"Vinayak": 300,
		}
	

func TestMapping(t *testing.T) {

	// mapping := map[string]int{}
	// mapping["Shailendra"] = 100
	// mapping["Soorya"] = 150
	// mapping["Aman"] = 500
	// mapping["Vinayak"] = 300

	keys := []string{}
	for k, v := range mapping {
		keys = append(keys, k)
		fmt.Println("values are: ",v)
	}

	sort.Strings(keys)
	for i, key := range keys {
		fmt.Printf("At index i: %d, Key: %s and Value: %v\n", i, key, mapping[key])
	}
	
}

func TestMapSortByValues(t *testing.T) {
	//create a slice of key-value pair
	type kv struct {
		Key string
		Value int
	}
	var kvSlice []kv
	for k, v := range mapping {
		kvSlice = append(kvSlice, kv{k,v})
	}

	sort.Slice(kvSlice, func(i, j int) bool { // equivalent to comparators in java
		return kvSlice[i].Value < kvSlice[j].Value
	})
	for _, kv := range kvSlice {
		fmt.Printf("Key: %v and Value: %v\n", kv.Key, kv.Value)
	}

}