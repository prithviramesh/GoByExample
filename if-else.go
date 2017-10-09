package main

import "fmt"

func main(){

	if 3%2 == 0{
		fmt.Println("3 is even");
	} else{
		fmt.Println("3 is odd");
	}

	if 12%4 == 0{
		fmt.Println("12 is divisible by 4");
	} else{
		fmt.Println("12 is not divisible by 4");
	}

	if num:= 8; num < 0 {
		fmt.Println(num, "is negative");
	} else if num < 10{
		fmt.Println(num, "has one digit");
	} else{
		fmt.Println(num, "has multiple digits");
	}

}