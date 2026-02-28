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
	
	code := HuffmanEncode(content)

	file, err := os.OpenFile("output.hfmn", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bitsWriter := BitsWriter{file: file, threshold: 64*1000}
	for i:= range content{
		bitsWriter.Write(code[string(content[i])])
	}	
	bitsWriter.Flush()
}
