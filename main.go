package main

import (
	"os"
)

func main(){
	file_name := os.Args[1]	

	content, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}
	
	huffman_tree := HuffmanEncode(content)
	file, err := os.OpenFile("output.hfmn", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Compress(huffman_tree, content, file)

	file_content, err := os.ReadFile("output.hfmn")

	if err != nil {
		panic(err)
	}
	Decompress(file_content)
}
