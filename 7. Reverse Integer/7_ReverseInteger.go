package main

import (
	"fmt"
	"strconv"
	"math"
)


func main() {
	
	one := 123
	two := -123
	third := -120
	four := 1534236469

	reverse(four)

	fmt.Println(one+two+third)
}

func reverse(x int) int {
	
	flag := false
	if(x < 0 ){
		flag = true
		x = int(math.Abs(float64(x)))
	}

	value := []byte(strconv.Itoa(x))
	lens := len(value)
	var ans = make([]byte,lens,lens)
	ans[lens/2] = value[lens/2]
	
	for index := 0; index < lens / 2; index++ {
		ans[index] = value[lens - index - 1]
		ans[lens - index - 1 ] = value[index]
	}
	aByteToInt, _ := strconv.Atoi(string(ans))

	if(flag){
		aByteToInt = aByteToInt * -1
	}

	if(float64(aByteToInt) > (math.Pow(2,31) - 1) || (float64(aByteToInt) < (math.Pow(2,31) * -1))){
		return 0 
	}

	return aByteToInt
}