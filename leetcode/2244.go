package main 


func main() {

	var tasks = []int{2,2,2,2,2}
	
	var freqMap = make(map[int]int)

	var round int

	for _, ele:= range tasks{
		_, status := freqMap[ele]

		if !status{
			freqMap[ele] = 1
		} else{
			freqMap[ele] += 1
		}
	}

	flag := false
	for _,value := range freqMap{

		if value == 1{
			flag = true
		} else{
			round += (value+2)/3
		}
	}

	if flag{
		println(-1)
	} else{
		println(round)
	}


}