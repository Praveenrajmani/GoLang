package main

import "fmt"

func main() {
	
	mynum := 2
 	//result := Is_prime(mynum)
	if Is_prime(mynum) == 1 {
 		fmt.Println("Prime")
	} else if Is_prime(mynum) == 0 {
		fmt.Println("Not Prime")	
	}
}






func Is_prime(number int) int {

	var count int = 0

	for i:=1; i<=number; i++ {

		if number%i == 0 {
			count++
		}
	}

	if count == 2 {
	return 1
	} else {
	return 0
	}	
}


