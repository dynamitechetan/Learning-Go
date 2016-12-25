package main

import (
	"crypto/md5"
	"fmt"
	"hash"
)

var expected = "e59ff97941044f85df5297e1c302d260"

func main() {
	fmt.Printf("<start md5.go>\n")
	original := "Hello World\n"
	var h hash.Hash = md5.New()
	h.Write([]byte(original))
	calculated := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Printf("%s: %x\n", original, h.Sum(nil))
	fmt.Printf("%s calculated\n", calculated)
	fmt.Printf("%s expected\n", expected)
	if expected != calculated {
		fmt.Printf("result did not match expected\n")
	} else {
		fmt.Printf("result matched expected\n")
	}
}