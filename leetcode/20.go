package main

import "fmt"

func isValid(str string) bool {

	var stack = []string{}

	for iter := 0; iter < len(str); iter++ {
		if str[iter] == '(' || str[iter] == '{' || str[iter] == '[' {

			stack = append(stack, string(str[iter]))

		} else {

			if len(stack) == 0 {
				return false
			} else if string(str[iter]) == ")" && stack[len(stack)-1] != "(" ||
				string(str[iter]) == "}" && stack[len(stack)-1] != "{" ||
				string(str[iter]) == "]" && stack[len(stack)-1] != "[" {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	str := "}}{{"

	out := isValid(str)

	fmt.Println(out)
}
