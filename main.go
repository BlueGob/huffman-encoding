package main

import (
	"fmt"
	"os"
)

func main(){
	file_name := os.Args[1]
	content, err := os.ReadFile(file_name)
	
	if err != nil {
		panic(err)
	}
	letters_freq := make(map[byte]int)

	for i := range content{
		letters_freq[content[i]] += 1	
	}
	for k, v := range letters_freq{
		fmt.Printf("%v => %v \n", string(k), v)
	}
}
