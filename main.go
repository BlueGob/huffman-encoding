package main

import (
	"os"
	"fmt"
)

func main(){
	file_name := os.Args[1]	

	content, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}
	
	huffman_tree := HuffmanEncode(content)
	code := TreeToDict(huffman_tree)
	for k, v :=range code{ //for debugging !
		fmt.Printf("byte = %v, freq %0*b \n", string(k), v.len, v.bits)
	}
	fmt.Println("len map = ", len(code))
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
