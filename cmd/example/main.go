package main

import (
	"os"
	"bytes"
	"fmt"
)

func main(){
	file_name := os.Args[1]	
	//generate tree from reading file
	f, err := os.Open(file_name)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	huffman_tree := HuffmanEncode(f)

	file, err := os.OpenFile("out.hfmn", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = f.Seek(0,0)
	if err != nil{
		panic(err)
	}

	Compress(huffman_tree, f, file)

	file, err = os.Open("out.hfmn")
	// file2, err1 := os.OpenFile("decompressed",os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// if err1 != nil {
	// 	panic(err1)
	// }
	var b bytes.Buffer
	Decompress(file, &b)
	fmt.Println(string(b.String()))
}
