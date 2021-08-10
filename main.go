package main

import (
	"Huffman/tree"
	"fmt"
	"os"
	"strings"
)

func main() {
	text := strings.Join(os.Args[1:], " ")
	tree := tree.GenerateTree(text)
	b := tree.Encode(text)

	fmt.Println(b)
	strDecoded := tree.Decode(b)

	fmt.Println(strDecoded)
}
