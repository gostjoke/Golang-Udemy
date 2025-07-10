package container

import "fmt"

func HashmapExample() {
	hashone := map[string]string{
		"key1":  "value1",
		"key2":  "value2",
		"key10": "value10",
	}
	fmt.Printf("Hashmap: %v\n", hashone)
	hashone["key3"] = "value3" // 添加新鍵值對
	fmt.Printf("Updated Hashmap: %v\n, length is %v\n", hashone, len(hashone))
	delete(hashone, "key3")
	for key, value := range hashone {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	fmt.Printf("After Deletion Hashmap: %v\n, length is %v\n", hashone, len(hashone))

	hashtwo := make(map[string]int) // 使用 make 創建空的 map
	fmt.Printf("Empty Hashmap: %v\n, length is %v\n", hashtwo, len(hashtwo))

	hashthree := map[string]map[string]string{
		"key1": {"key1": "value1"},
		"key2": {"key2": "value2"},
	}
	fmt.Printf("Nested Hashmap: %v\n, length is %v\n", hashthree, len(hashthree))

	hashfour := map[string][]int{
		"key1": {1, 2, 3},
		"key2": {4, 5, 6},
	}
	fmt.Printf("Hashmap with Slice Values: %v\n, length is %v\n", hashfour, len(hashfour))
}
