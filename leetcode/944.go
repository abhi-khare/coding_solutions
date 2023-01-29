package main

func minDeletionSize(strs []string) int {
    
    delMap := make([]int, len(strs[0]))

    if len(strs) == 1{
        return 0
    }

    for iter:=1; iter<len(strs); iter++{

        for jter:=0; jter<len(strs[0]); jter++{

            if strs[iter][jter] < strs[iter-1][jter]{
                delMap[jter] = 1
            }
        }
    }

    delCount :=0

    for _, ele := range delMap{
        if ele == 1{
            delCount +=1
        }
    }

	return delCount
}

func main() {

	pattern := []string{"a", "e", "p"}

	count := minDeletionSize(pattern)

	print(count)
}