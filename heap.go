package main

func pop(heap []Node) ([]Node,Node) {
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

func insert(heap []Node, n Node) []Node {
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
