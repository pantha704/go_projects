package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main(){
	// fmt.Println("Hello World")

	var numi int16 = 32767  // maximum value that can be assigned to int16 
	fmt.Println(numi);

	numf, numu := float32(12),uint32(64) // type casting
	fmt.Println(numf,numu)

	fmt.Println(utf8.RuneCountInString("ao")) // import "unicode/utf8" and use utf8.RuneCountInString() to get the length of a string 
	// returns the length of "ao".

	fmt.Println(rune('A'))
	// returns ASCII value of 'A'

	var result, remainder, err = intDiv(4, 0)
	if err!= nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result, remainder)
	}

	fmt.Println("\nArrays,")
	Arrays()
}

func intDiv(num int, den int) (int, int, error){
	var err error
	if den == 0 {
		err = errors.New("Cannot divide by 0")
		return 0, 0, err
	}
	result := num / den
	remainder := num % den

	return result, remainder, err
}

func Arrays(){
	// init
	// var intArr [3]int
	intArr := [3]int{1,2,3}
	// var intArr = [3]int{1,2,3}
	// intArr := [3]int{1,2,3}
	fmt.Println(intArr)

	// intArr := [...]int{1,2,3,4}  array size becomes 4

}