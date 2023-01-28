
package main

import (
	"fmt"
	"strconv"
)
func RunLengthEncode (input string) (encodedOutput string) {
	rep :=1
	for i := range input {
	if i < len(input)-1 {
		if input[1] == input[i+1] {
			rep++
		} else {
			numOfRep := ""
			if rep > 1 {
				numOfRep = strconv.Itoa(rep)
			}
			encodedOutput += numOfRep + string(input[i])
			rep = 1

		}
	}else{
		numOfReps := ""
		if rep > 1 {
			numOfReps = strconv.Itoa(rep)

		}
		encodedOutput += numOfReps + string(input[i])
	}
}
	return 
}
func RunLengthDecode (input string) (encodedOutput string) {
	rep :=1
	for i := range input {
	if i < len(input)-1 {
		if input[1] == input[i+1] {
			rep++
		} else {
			numOfRep := ""
			if rep > 1 {
				numOfRep = strconv.Itoa(rep)
			}
			encodedOutput += numOfRep + string(input[i])
			rep = 1

		}
	}else{
		numOfReps := ""
		if rep > 1 {
			numOfReps = strconv.Itoa(rep)

		}
		encodedOutput += numOfReps + string(input[i])
	}
}
	return 
}

func main(){
	fmt.Println(RunLengthDecode("12W13WBB25WB"))
	fmt.Println(RunLengthEncode("WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"))
}
