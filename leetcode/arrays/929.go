package main 

import "fmt"

func processEmailAddress(email string) string{
	processedEmail := ""
	ignore := false
	for idx, ele := range email{
		if ele == '+'{
			ignore = true
		} else if ele == '@'{
			processedEmail += email[idx:]
			break
		} else if !ignore && ele >= 'a' && ele <= 'z'{
			processedEmail +=  string(ele)
		}
	}

	return processedEmail
}

func numUniqueEmails(emails []string) int {
    duplicates := 0
	hashMap := make(map[string]bool)

	for _, ele := range emails{

		processedEmail := processEmailAddress(ele)

		if _, ok := hashMap[processedEmail]; ok{
			duplicates +=1
		} else{
			hashMap[processedEmail] = true
		}

	}

	

	return len(emails) - duplicates
}


func main(){
	emails := []string{"cdr.+b.+f.d+s.@leetcode.com","b@leetcode.com","c@leetcode.com"}

	ans:= numUniqueEmails(emails)

	fmt.Println(ans)
}