package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	fmt.Println(GetMD5Hash(input))
}

func GetMD5Hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
