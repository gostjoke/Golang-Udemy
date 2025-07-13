package interfacespolymorphstringer

import (
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) WriteOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first + "\n"))
	return err
}
func WriterExample() {
	p := person{first: "John"}
	// if file exists, return error
	_, err := os.Stat("interfacespolymorphstringer\\writer.txt")
	if !os.IsNotExist(err) {
		log.Fatalf("Error checking file: %s", err)
		return
	}

	f, err := os.Create("interfacespolymorphstringer\\writer.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	p.WriteOut(f)

	defer f.Close()
}
