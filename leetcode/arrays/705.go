package main

import (
	"fmt"
	"math"
)

type MyHashSet struct {
    arr []int16
}


func Constructor() MyHashSet {
    var hashSet MyHashSet
	hashSet.arr = make([]int16, 250000)
	return hashSet
}


func (this *MyHashSet) Add(key int)  {

	
    arrIdx := int(key/8)
	offset := key%8
	
	this.arr[arrIdx] |=  int16(math.Pow(2, float64(offset)))

}


func (this *MyHashSet) Remove(key int)  {
    arrIdx := int(key/8)
	offset := key%8

	// checking if the bit is already set
	isSet := this.arr[arrIdx] & int16(math.Pow(2, float64(offset)))
	if key == 48{
		fmt.Println("remove", arrIdx, offset, this.arr[arrIdx])
	}
	if isSet > 0{
		this.arr[arrIdx] ^=  int16(math.Pow(2, float64(offset)))
	}
}


func (this *MyHashSet) Contains(key int) bool {
    arrIdx := int(key/8)
	offset := key%8
	out := this.arr[arrIdx] &  int16(math.Pow(2, float64(offset)))

	if out > 0{
		return true
	}
	return false
}


func main(){

	obj := Constructor()

	operations := []string {"contains","remove","add","add","contains","remove","contains","contains","add","add","add","add","remove","add","add","add","add","add","add","add","add","add","add","contains","add","contains","add","add","contains","add","add","remove","add","add","add","add","add","contains","add","add","add","remove","contains","add","contains","add","add","add","add","add","contains","remove","remove","add","remove","contains","add","remove","add","add","add","add","contains","contains","add","remove","remove","remove","remove","add","add","contains","add","add","remove","add","add","add","add","add","add","add","add","remove","add","remove","remove","add","remove","add","remove","add","add","add","remove","remove","remove","add","contains","add"}
	value := []int{72,91,48,41,96,87,48,49,84,82,24,7,56,87,81,55,19,40,68,23,80,53,76,93,95,95,67,31,80,62,73,97,33,28,62,81,57,40,11,89,28,97,86,20,5,77,52,57,88,20,48,42,86,49,62,53,43,98,32,15,42,50,19,32,67,84,60,8,85,43,59,65,40,81,55,56,54,59,78,53,0,24,7,53,33,69,86,7,1,16,58,61,34,53,84,21,58,25,45,3}
	ans := []bool{false,false,false,false,false,false,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,true,false,false,false,false,false,false,false,false,true,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,false,true,false,false,false,false,false,false,true,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false}

	for iter:=0; iter< len(operations);iter++{
		if operations[iter] == "add"{
			obj.Add(value[iter])
		} else if operations[iter] == "remove"{
			obj.Remove(value[iter])
		} else{

			out:= obj.Contains(value[iter])

			if (out != ans[iter]){
				fmt.Println(iter, out, operations[iter], value[iter], ans[iter])
			}
		}
	}
}