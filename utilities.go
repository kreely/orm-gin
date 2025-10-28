package main

import (
	"log"
	"os"
	"fmt"
)

func printFiles() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
 
	for _, e := range entries {
		fmt.Println(e.Name())
	}
}
