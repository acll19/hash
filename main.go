package main

import (
	"flag"
	"fmt"
	"hash/fnv"
)

var length int

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: hash [string] [-l length]")
		return
	}

	input := args[0]
	h := fnv.New64a()
	h.Write([]byte(input))
	hash := fmt.Sprintf("%x", h.Sum64())

	if length > 0 && length < len(hash) {
		hash = hash[:length]
	}

	fmt.Printf("Hash: %s\n", hash)
}

func init() {
	flag.IntVar(&length, "l", 10, "Length of the hash")
}
