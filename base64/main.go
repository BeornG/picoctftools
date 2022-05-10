package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	fmt.Println("Decoded text:")
	decoder(input)
}

func decoder(input string) {
	decodedstring, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(decodedstring))
}
