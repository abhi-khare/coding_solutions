package main

import "fmt"

type maxHeap struct {
	arr []int
}

func (h *maxHeap) Insert(key int) {
	// adding the element
	h.arr = append(h.arr, key)
	h.maxHeapify(len(h.arr) - 1)
}

func (h *maxHeap) maxHeapify(index int) {

	for h.arr[parent(index)] < h.arr[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *maxHeap) maxHeapifyDown(index int) {

	for leftChild(index) < len(h.arr) || rightChild(index) < len(h.arr) {
		if h.arr[leftChild(index)] > h.arr[rightChild(index)] {
			h.swap(leftChild(index), index)
		} else {
			h.swap(rightChild(index), index)
		}
	}
}

func (h *maxHeap) Extract() int {

	out := h.arr[0]
	if len(h.arr) == 0 {
		fmt.Println("No element present in the heap")
		return -1
	}
	len := len(h.arr) - 1
	h.arr[0] = h.arr[len]
	h.arr = h.arr[:len]

	h.maxHeapifyDown(0)

	return out
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func (h *maxHeap) swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}

func main() {

	buildHeap := []int{10, 20, 30, 40, 15}

	var heap maxHeap
	for _, ele := range buildHeap {
		heap.Insert(ele)
	}

	fmt.Println(heap.arr)

}
