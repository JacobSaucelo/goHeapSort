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

}

func (h *Heap) Insert(key int) {
	h.slice = append(h.slice, key)
	h.MaxHeapUp(len(h.slice) - 1)
}

func (h *Heap) MaxHeapUp(index int) {
	for h.slice[Parent(index)] < h.slice[index] {
		h.Swap(Parent(index), index)
		index = Parent(index)
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

// https://youtu.be/3DYIgTC4T1o?list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6&t=744
