package main

import "fmt"

func main() {

	var freqMAp = make(map[int]int)

	fmt.Println("Enter 3 key-value pairs:")
	for iter := 0; iter < 3; iter++ {
		var key, value int
		fmt.Scan(&key, &value)
		freqMAp[key] = value
	}

	// printing all the key value pairs

	for key, value := range freqMAp {
		println("Key: ", key, "value: ", value)
	}

	// check if a key is present

	value, status := freqMAp[33]

	if status {
		println("Key", value, "is present")
	} else {
		println("Key", 33, "is not present, returned default value 0")
	}
}
