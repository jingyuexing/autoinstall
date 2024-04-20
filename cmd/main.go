package main

import (
	"autoinstall/check"
	"fmt"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		return
	}
	check.Analyze(path)
	for _, tech := range check.Analyze(path) {
		fmt.Printf("Program Language: %s ,Tech: %s\n", tech.Language, tech.Type)
	}
}
