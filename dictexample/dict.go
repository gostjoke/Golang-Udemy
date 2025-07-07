package dictexample

import (
	"encoding/json"
	"fmt"
)

type Dictionary map[string]string

func (d Dictionary) Add(key, value string) {
	if _, exists := d[key]; exists {
		panic("key already exists")
	}
	d[key] = value
	pretty, _ := json.MarshalIndent(d, "", "  ") //
	// fmt.Println(pretty) 會打印byte
	fmt.Println(string(pretty))
}
