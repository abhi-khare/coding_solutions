package main

import "fmt"

type node struct {
	value int 
	ptr *node // pointer to a variable of type node
}

func insertion(ptr *node, node_value int) *node {

	var temp_node node 
	temp_node.value = node_value 
	temp_node.ptr = nil 

	if ptr == nil{
		ptr = &temp_node // first node is basically set to head ptr
	} else{
		iterator := ptr // otherwise, we iterate till the end of the linkedlist and then set the node

		for iterator.ptr!= nil{
			iterator = iterator.ptr
		}
		iterator.ptr = &temp_node
	}
	return ptr
}

func searching(ptr *node, key int) bool {

	var iterator *node = ptr 

	for iterator != nil{

		if iterator.value == key{
			println("Element found")
			return true
		}
		iterator = iterator.ptr
	}
	println("Element Not found")
	return false
}

func print_list(iterator *node){
	// iterating through the linkedlist
	for iterator!= nil {
		println(iterator.value)
		iterator = iterator.ptr
	}
}

func search_and_delete(ptr *node, key int) *node {

	iterator := ptr

	if iterator.value == key {
		ptr = ptr.ptr 
		return ptr
	}

	for iterator.ptr.ptr != nil  {

		if iterator.ptr.value != key{
			iterator = iterator.ptr
		} else {
			iterator.ptr = iterator.ptr.ptr
			return ptr
		}
	}

	if iterator.ptr.value == key {
		iterator.ptr = nil
		return ptr
	}

	return ptr
}

func main() {

	var num_ele int
	fmt.Println("Enter the number of elements:")
	fmt.Scan(&num_ele)

	var head *node = nil // initialising head of the linkedlist and setting it to nil (null)

	fmt.Println("Enter the elements")

	for iter:=0; iter < num_ele; iter++{
		var node_value int 
		fmt.Scan(&node_value)

		head = insertion(head, node_value)
	}

	print_list(head)

	var key int
	fmt.Println("Enter the elements to be searched")
	fmt.Scan(&key)
	searching(head, key)

	// deleting an element if present in the linkedlist
	var del_key int
	fmt.Println("Enter the elements to be deleted")
	fmt.Scan(&del_key)
	head = search_and_delete(head, del_key)

	print_list(head)

}