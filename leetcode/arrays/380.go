package main 

import "rand"
import "math"


type RandomizedSet struct {
    hashMap map[int]int
}


func Constructor() RandomizedSet {
    var hashMap RandomizedSet
	return hashMap
}


func (this *RandomizedSet) Insert(key int) bool {
    if val, ok := this.hashMap[key]; ok{
		if val == 1{
			return false
		} else{
			this.hashMap[key] = 1
			return true
		}
	}

	this.hashMap[key] = 1
	return true

}


func (this *RandomizedSet) Remove(key int) bool {
    if val, ok := this.hashMap[key]; !ok{
		return false
	} else {
		if val == 0{
			return false
		} else {
			this.hashMap[key] = 0
			return true
		}
	}
}


func (this *RandomizedSet) GetRandom() int {
    
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main(){


}