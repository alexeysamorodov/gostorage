package main

import (
	"fmt"
	"log"

	"github.com/alexeysamorodov/gostorage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("hello", []byte("hello"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Uploaded file", file)

	foundFile, err := st.GetById(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Read file", foundFile)
}
