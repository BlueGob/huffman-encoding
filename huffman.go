package main

type Node struct{
	name byte
	freq uint64
	left_child *Node
	right_child *Node
}
type Code struct{
	len uint64
	bits uint64
}

func build_huffman_tree(heap []Node) Node{
	for len(heap) != 1{
		var node1, node2 Node
		heap, node1 = Pop(heap)
		heap, node2 = Pop(heap)
		new_node := Node{name: node1.name+node2.name, freq: node1.freq+node2.freq,left_child: &node1, right_child: &node2 }
		heap = Insert(heap, new_node) 
	}
	return heap[0]
}
func generate_code(n Node, code uint64, code_map map[byte]Code, depth uint64){	
	if n.left_child == nil && n.right_child == nil{
		code_map[n.name] = Code{len: depth, bits: code}
	}else{
		generate_code(*n.left_child, code<<1, code_map, depth+1)
		generate_code(*n.right_child, (code<<1) | 1, code_map, depth+1)
	}
}

func HuffmanEncode(file_content []byte) Node{
	heap := make([]Node, 0, 26)
	freqs := letters_freqs(file_content)
	for k, v := range freqs{
		heap = Insert(heap, Node{name: k, freq: v, left_child: nil, right_child: nil})
	}			
	return build_huffman_tree(heap)
}

func TreeToDict(node Node) map[byte]Code{
	code := make(map[byte]Code)
	generate_code(node, 0, code, 0)
	return code
}

func letters_freqs(file_content []byte) map[byte]uint64{
	m := make(map[byte]uint64)
	for i := range file_content {
		m[file_content[i]] += 1
	}
	return m
}

