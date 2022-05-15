package main

import (
	"fmt"
	"io"
	"os"
)

var nums = []int{}

func main() {
	filetointarray()
	fmt.Println(nums)
	inttochars()
}

func inttochars() {
	for i := range nums {
		fmt.Printf("%c", nums[i])
	}
}

func filetointarray() {
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
		nums = append(nums, line)
	}
}
