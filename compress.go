package main

import (
	"os"
	"fmt"
	"encoding/binary"
)

type BitsWriter struct{
	buffer byte
	count  uint8
	output_buffer []byte
	threshold int
	file 		*os.File
}

func (b *BitsWriter) writeBit(bit uint8) {
	b.buffer <<= 1
	b.buffer |= bit
	b.count += 1
	if b.count == 8{ //byte is full
		b.output_buffer = append(b.output_buffer, b.buffer)
		b.count = 0
		b.buffer = 0
	}
	if len(b.output_buffer) == b.threshold{
		b.flush()
		b.output_buffer = b.output_buffer[:0]
	}
}
func (b BitsWriter) flush(){
	n, err := b.file.Write(b.output_buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("writen %v byte\n", n)
}

func (b *BitsWriter) write(code Code){
	for i:= code.len; i > 0; i--{
		b.writeBit(uint8((code.bits >> (i - 1)) & 1))
	} 
}
func header(node Node, h []byte)[]byte{
	if node.right_child == nil && node.left_child == nil{
		h = append(h, 1)
		h = append(h, node.name[0])
		return h
	}
	h = append(h, 0)
	h = header(*node.left_child, h)
	h = header(*node.right_child, h)
	return h
}
func Compress(node Node, content []byte, file *os.File){
	bitsWriter := BitsWriter{file: file, threshold: 64*1000}
	metadata := make([]byte, 0)
	metadata = header(node, metadata)
	for i := range metadata{fmt.Printf("%b \n", metadata[i])}
	header_len := make([]byte, 4)
	binary.BigEndian.PutUint32(header_len, uint32(len(metadata)))
	file.Write(header_len)
	file.Write(metadata)
	code := TreeToDict(node)
	for i:= range content{
		bitsWriter.write(code[content[i]])
	}	
	bitsWriter.flush()
}
