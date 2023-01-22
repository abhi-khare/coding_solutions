package main 


func getAllValidParan(num_pair int) []string{

	if num_pair==1{
		baseArr := []string{"()"}
		return baseArr
	} else{
		
		prevAns := getAllValidParan(num_pair-1)

		var currAns = []string{}

		for iter:=0; iter <len(prevAns); iter++{

			str := prevAns[iter]

			redundantStr := ""
			for jter:=0; jter<num_pair-1; jter++{
				redundantStr += "()"
			}

			//print(redundantStr)
			validStr := []string{}

			if str != redundantStr{
				
				validStr = append(validStr, "()" + str)
				validStr = append(validStr,  str + "()")
				
			} else {
				validStr = append(validStr, "()" + str)
			}

			// check if string is symmetric
			
			validStr = append(validStr, "(" + str + ")")
			
			currAns = append(currAns,validStr...)
			
		}
		return currAns
	}
}