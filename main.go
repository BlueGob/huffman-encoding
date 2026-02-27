package main

import (
	"fmt"
	"os"
)
type Node struct{
	str string
	freq int
	left_child *Node
	right_child *Node
}
func pop_root(heap []Node) ([]Node,Node) {
	pos := 0
	var smallest_child int
	element := heap[pos]

	heap[pos] = heap[len(heap) - 1]
	heap = heap[:len(heap)-1] // remove last element 
	//rebalence heap
	for 2*pos + 1 < len(heap)  {
		left_child  := 2*pos + 1
		right_child := 2*pos + 2
		if right_child >= len(heap){
			smallest_child = left_child
		}else if heap[right_child].freq <= heap[left_child].freq {
			smallest_child = right_child
		}else if heap[right_child].freq > heap[left_child].freq {
			smallest_child = left_child 
		}
		if heap[smallest_child].freq < heap[pos].freq{
			temp := heap[pos]
			heap[pos] = heap[smallest_child]
			heap[smallest_child] = temp
			pos = smallest_child 
		}else{
			break
		}
	}
	return heap, element
}
func bubble_up(heap []Node, n Node) []Node {
	heap = append(heap, n)
	pos := len(heap) - 1	
	for pos != 0 {
		parent_id := (pos - 1) / 2
		if heap[parent_id].freq > heap[pos].freq {
			temp := heap[pos]
			heap[pos] = heap[parent_id]
			heap[parent_id] = temp
		}else{
			break //no more bubbling up
		}
		pos = parent_id
	}
	return heap
}

func construct_huffman_tree(heap []Node) Node{
	for len(heap) != 1{
		var node1, node2 Node
		heap, node1 = pop_root(heap)
		heap, node2 = pop_root(heap)
		new_node := Node{str: node1.str+node2.str, freq: node1.freq+node2.freq,left_child: &node1, right_child: &node2 }
		heap = bubble_up(heap, new_node) 
	}
	return heap[0]
}
func print_tree(n Node){
	fmt.Printf("node_name = %v | freq = %v \n", n.str, n.freq)
	if n.left_child != nil{
		print_tree(*n.left_child)
	}
	if n.right_child != nil{
		print_tree(*n.right_child)
	}
}
func generate_code(n Node, code string, code_map map[string]string){	
	if n.left_child != nil{
		generate_code(*n.left_child, code+"0", code_map)
	}else{
		code_map[n.str] = code
	}
	if n.right_child != nil{
		generate_code(*n.right_child, code+"1", code_map)
	}else{
		code_map[n.str] = code
	}
}

func main(){
	file_name := os.Args[1]
	
	heap := make([]Node, 0, 26)

	content, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}
	letters_freq := make(map[string]int)
	
	// calculate letters freqs
	for i := range content{
		// if string(content[i]) == "\n"{
		// 	continue
		// }
		letters_freq[string(content[i])] += 1	
	}
	for k, v := range letters_freq{
		fmt.Printf("%v => %v \n", string(k), v)
	}	
	//init the heap
	for k, v := range letters_freq{
		heap = bubble_up(heap, Node{str: k,freq: v, left_child: nil, right_child: nil})
	}

	// for i:=range heap{
	// 	fmt.Printf("name = %v, freq = %v \n", heap[i].str, heap[i].freq)
	// }
	root := construct_huffman_tree(heap)
//	print_tree(root)
	code := make(map[string]string)
	generate_code(root, "", code)
	fmt.Println(len(code))		
	for k, v:=range code{
		fmt.Printf("letter = %v, code = %v \n", k, v)
	}
	//calculate file size
	var total uint64
	for k ,v := range code{
		total += uint64(len(v) * letters_freq[k])
	}
	fmt.Printf("new size = %vMB", total/(8*1000*1000))

}
