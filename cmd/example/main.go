package main

import (
	"github.com/BlueGob/huffman-enconding"
	"os"
)

func main() {
	// Open source file
	input, _ := os.Open("../../testdata/t8.shakespeare.txt")
	defer input.Close()

	// 1. Build the tree
	tree, _ := huffman.HuffmanEncode(input)

	// 2. Reset file pointer to beginning
	input.Seek(0, 0)

	// 3. Create destination file
	output, _ := os.Create("document.huff")
	defer output.Close()

	// 4. Compress
	huffman.Compress(tree, input, output)
}
