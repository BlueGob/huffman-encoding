package main

import (
	"encoding/binary"
	"bufio"
	"io"
	"fmt"
)
type DecodeTree struct{
	name byte
	left_child *DecodeTree
	right_child *DecodeTree
}

func Decompress(in io.Reader, out io.Writer) error{
	var header_len uint32 
	//reader header length
	r := bufio.NewReader(in)
	err := binary.Read(r, binary.BigEndian, &header_len)
	if err != nil{
		return err
	}
	//extract header data
	headers := make([]byte, header_len)
	_, err = io.ReadFull(r, headers)	
	if err != nil {
		return err
	}
	//decode headers	
	i := 0
	treeRoot, err := decodeHeaders(headers, &i)
	if err != nil {
		return err
	}
	tree := treeRoot

	w := bufio.NewWriter(out)
	buf := make([]byte, 4096)
	//decode text
	for {
		n, err := r.Read(buf)
		if n > 0{
			for i := range n {
				for j := range 8{
					tree = decode((buf[i] >> (7 - j)) & 1, tree)
					if tree.name != 0{
						w.WriteByte(tree.name)
						tree = treeRoot
					}
				}
			}
		}
		if err == io.EOF {
			w.Flush()
			break
		}
		if err != nil{
			return err
		}
	}
	return nil
}
func decode(b byte, t *DecodeTree) *DecodeTree{
	if b == 1{
		return t.right_child
	}
	return t.left_child
}

func decodeHeaders(content []byte, i *int) (*DecodeTree,error){
	if (*i) >= len(content){
		return nil, fmt.Errorf("Malformed header")
	}
	code := content[*i]
	*i = (*i) + 1
	if int(code) == 1{
		if (*i) >= len(content){
			return nil, fmt.Errorf("Malformed header")
		}
		name := content[*i]
		*i = (*i)+1
		return &DecodeTree{name: name}, nil
	}	
	left_child, err := decodeHeaders(content, i)
	if err != nil {
		return nil, err
	}
	right_child, err := decodeHeaders(content, i)
	if err != nil {
		return nil, err
	}
	return &DecodeTree{
		left_child: left_child,
		right_child: right_child,
	}, nil
}
