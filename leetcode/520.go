package main

import "unicode"

func detectCapitalUse(word string) bool {

    case1, case2, case3 := "", "", "1"

    for iter:=0; iter<len(word); iter++{
        case1 += "1"
        case2 += "0"
		if iter != 0{
			case3 += "0"
		}
    }

    // parsing the word
    parseWord := ""
    for _, r := range word {
        if unicode.IsUpper(r) {
            parseWord += "1"
        } else {
            parseWord += "0"
        }
    }

    if parseWord == case1 || parseWord == case2 || parseWord == case3 {
		return true
	}

	return false
    
}


