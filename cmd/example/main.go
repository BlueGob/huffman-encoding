package main

import (
	"os"
	"bytes"
	"fmt"
	"github.com/BlueGob/huffman-enconding"
)

func main(){
	file_name := os.Args[1]	
	f, err := os.Open(file_name)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	huffman_tree, err := huffman.HuffmanEncode(f)

	file, err := os.OpenFile("out.hfmn", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = f.Seek(0,0)
	if err != nil{
		panic(err)
	}

	huffman.Compress(huffman_tree, f, file)

	file, err = os.Open("out.hfmn")
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	huffman.Decompress(file, &b)
	fmt.Println(string(b.String()))
}
