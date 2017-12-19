package main

import "fmt"

const (
	duplicateVowel bool = true
	removeVowel    bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	fmt.Println("vim-go")
}
