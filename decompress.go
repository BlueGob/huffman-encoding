package main

import (
	"encoding/binary"
	"fmt"
)
type DecodeTree struct{
	name byte
	left_child *DecodeTree
	right_child *DecodeTree
}

func Decompress(content []byte){
	header_len := binary.BigEndian.Uint32(content[:4])
	i := 0
	m := decodeHeaders(content[4:header_len+4], &i)
	n := m
	file_content := content[4+header_len:]
	res := make([]byte, 0, 64*1000)
	for i := range file_content{
		for j := range 8{
			n = decode((file_content[i] >> (7 - j)) & 1, n)
			if n.name != 0{
				res = append(res, n.name)
				n = m
			}
		}
		if len(res) >= 64*1000{
			fmt.Print(string(res))
			res = res[:0]
		}
	}
	fmt.Print(string(res))
}
func decode(b byte, t *DecodeTree) *DecodeTree{
	if b == 1{
		return t.right_child
	}
	return t.left_child
}

func decodeHeaders(content []byte, i *int) *DecodeTree{
	code := content[*i]
	*i = (*i) + 1
	if int(code) == 1{
		name := content[*i]
		*i = (*i)+1
		return &DecodeTree{name: name}
	}	
	left_child := decodeHeaders(content, i)
	right_child := decodeHeaders(content, i)
	return &DecodeTree{
		left_child: left_child,
		right_child: right_child,
	}
}
