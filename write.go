package main

import (
	"os"
	"fmt"
)

type BitsWriter struct{
	buffer byte
	count  uint8
	output_buffer []byte
	threshold int
	file 		*os.File
}

func (b *BitsWriter) WriteBit(bit uint8) {
	if b.count == 8{ //byte is full
		b.output_buffer = append(b.output_buffer, b.buffer)
		b.count = 0
	}

	b.buffer <<= 1
	b.buffer |= bit
	b.count += 1
	if len(b.output_buffer) == b.threshold{
		b.Flush()
		b.output_buffer = b.output_buffer[:0]
	}
}
func (b BitsWriter) Flush(){
	n, err := b.file.Write(b.output_buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("writen %v byte\n", n)
}

func (b *BitsWriter) Write(code Code){
	for i:= code.len; i > 0; i--{
		b.WriteBit(uint8((code.bits >> (i - 1)) & 1))
	} 
}
