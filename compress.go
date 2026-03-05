package main

import (
	"encoding/binary"
	"io"
	"bufio"
)

type BitsWriter struct{
	buffer byte
	count  uint8
	writer *bufio.Writer
}

func (b BitsWriter) writeHeader(node Node){
	metadata := make([]byte, 0)
	metadata = header(node, metadata)
	header_len := make([]byte, 4)
	binary.BigEndian.PutUint32(header_len, uint32(len(metadata)))
	b.writer.Write(header_len)
	b.writer.Write(metadata)
}

func (b *BitsWriter) writeBit(bit uint8) {
	b.buffer <<= 1
	b.buffer |= bit
	b.count += 1
	if b.count == 8{ //byte is full
		b.writer.WriteByte(b.buffer)
		b.count = 0
		b.buffer = 0
	}
}

func (b *BitsWriter) write(code Code){
	for i:= code.len; i > 0; i--{
		b.writeBit(uint8((code.bits >> (i - 1)) & 1))
	} 
}

func header(node Node, h []byte)[]byte{
	if node.right_child == nil && node.left_child == nil{
		h = append(h, 1)
		h = append(h, node.name)
		return h
	}
	h = append(h, 0)
	h = header(*node.left_child, h)
	h = header(*node.right_child, h)
	return h
}

func Compress(node Node, in io.Reader, out io.Writer) error {
	buf := make([]byte, 4096)
	w := bufio.NewWriter(out)
	r := bufio.NewReader(in)
	bitsWriter := BitsWriter{writer: w}
	bitsWriter.writeHeader(node)
	code := TreeToDict(node)
	for {
		n, err := r.Read(buf)
		if n > 0{
			for i := range n {
				bitsWriter.write(code[buf[i]])
			}
		}
		if err == io.EOF {
			bitsWriter.writer.Flush()
			break
		}
		if err != nil{
			return err
		}
	}
	return nil
}
