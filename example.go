//go:build ignore

package main

import (
	"fmt"

	"github.com/hymkor/go-dutf"
)

func main() {
	sourceString := "your string to be encoded"
	bytes := dutf.EncodeString(sourceString)
	decodedString, _ := dutf.DecodeString(bytes)
	fmt.Println(decodedString)
}
