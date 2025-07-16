package JsonMarshal

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

func MarshalExpample() {
	// Example of how to use the Person struct
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "ddd@gmail.com",
		Age:       30,
	}

	jsonData, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON Data:", string(jsonData))
	fmt.Printf("Type of jsonData: %T\n", jsonData)
	// Convert jsonData to string and then to []byte
	s := string(jsonData)
	bs := []byte(s)

	var p2 Person

	err2 := json.Unmarshal(bs, &p2)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Unmarshaled Person:", p2)
}
