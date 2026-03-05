package huffman
import (
	"io"
	"bufio"
)
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

func buildHuffmanTree(heap []Node) Node{
	for len(heap) != 1{
		var node1, node2 Node
		heap, node1 = pop(heap)
		heap, node2 = pop(heap)
		new_node := Node{name: node1.name+node2.name, freq: node1.freq+node2.freq,left_child: &node1, right_child: &node2 }
		heap = insert(heap, new_node) 
	}
	return heap[0]
}
func generateCode(n Node, code uint64, code_map map[byte]Code, depth uint64){	
	if n.left_child == nil && n.right_child == nil{
		code_map[n.name] = Code{len: depth, bits: code}
	}else{
		generateCode(*n.left_child, code<<1, code_map, depth+1)
		generateCode(*n.right_child, (code<<1) | 1, code_map, depth+1)
	}
}

func TreeToDict(node Node) map[byte]Code{
	code := make(map[byte]Code)
	generateCode(node, 0, code, 0)
	return code
}

func HuffmanEncode(f io.Reader) (Node, error){	
	heap  := make([]Node, 0, 26)
	buf   := make([]byte, 4096)
	freqs := make(map[byte]uint64)
	r := bufio.NewReader(f)
	for {
		n, err := r.Read(buf)
		if n > 0{
			//procss
			for i := range n {
				freqs[buf[i]] += 1
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil{
			return Node{}, err
		}
	}
	for k, v := range freqs{
		heap = insert(heap, Node{name: k, freq: v, left_child: nil, right_child: nil})
	}			
	return buildHuffmanTree(heap), nil
}
