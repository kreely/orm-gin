package main

import (
	"log"
	"os"
	"fmt"
)

func printFiles() {
	log.Println("About to call os.ReadDir(\"./\")")
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
 
	for _, e := range entries {
		fmt.Println(e.Name())
	}
}
