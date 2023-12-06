package main

import "fmt"

type Heap struct {
	slice []int
}

var dummyData = []int{34, 27, 12, 35, 8, 23, 5, 2, 43, 78, 9}

func main() {
	container := &Heap{}
	fmt.Println("container: ", container)

	for _, value := range dummyData {
		container.Insert(value)
		fmt.Println("new container: ", container)
	}

	for i := 0; i < len(container.slice); i++ {
		container.Extract()
		fmt.Println("extract", container)
	}

}

func (h *Heap) Insert(key int) {
	h.slice = append(h.slice, key)
	h.MaxHeapUp(len(h.slice) - 1)
}

func (h *Heap) Extract() int {

	if len(h.slice) == 0 {
		fmt.Println("Empty slice")
		return -1
	}

	extracted := h.slice[0]
	l := len(h.slice) - 1

	h.slice[0] = h.slice[l]
	h.slice = h.slice[:l]

	h.MaxHeapDown(0)

	return extracted

}

func (h *Heap) MaxHeapUp(index int) {
	for h.slice[Parent(index)] < h.slice[index] {
		h.Swap(Parent(index), index)
		index = Parent(index)
	}
}

func (h *Heap) MaxHeapDown(index int) {
	lastIndex := len(h.slice) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.slice[l] > h.slice[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.slice[index] < h.slice[childToCompare] {
			h.Swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *Heap) Swap(index1, index2 int) {
	h.slice[index1], h.slice[index2] = h.slice[index2], h.slice[index1]
}

func Parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
