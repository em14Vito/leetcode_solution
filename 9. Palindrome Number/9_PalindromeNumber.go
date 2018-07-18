package main

import (
	"fmt"
	// "math"
)



func main() {
	
	one := "III"

	ans := romanToInt(one)

	fmt.Println(ans)
}

var dict palindromeList

type  palindrome struct{
	origin byte
	order int
	value int
}

type palindromeList []palindrome



func getStructs(value byte) palindrome{

	var ans palindrome
	
	for _, element := range dict {
		if element.origin == value {
			ans = element
		}
	}
	return ans
}

func romanToInt(s string) int {

	dict = initValue()

	sum := 0

	for index := 0; index < len(s); index++ {

		if(index + 1 < len(s)){
			first := getStructs(s[index])
			second := getStructs(s[index+1])
			if(second.order > first.order){
				sum = sum + second.value - first.value
				index++
				continue
			}
		}

		sum = getStructs(s[index]).value + sum 	
	}


	return sum
}


func initValue() []palindrome{

	palindromes := palindromeList{
		palindrome{'I',1,1},
		palindrome{'V',2,5},
		palindrome{'X',3,10},
		palindrome{'L',4,50},
		palindrome{'C',5,100},
		palindrome{'D',6,500},
		palindrome{'M',7,1000},

	}
	return palindromes
	
}