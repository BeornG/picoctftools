package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filetochar()
}

func filetochar() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var line int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(string(line))
	}
}
