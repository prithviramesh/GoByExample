package main

import "fmt"

func main(){
	s := make([]string, 3);

	fmt.Println("empty slice:", s)
	s[0] = "hello";
	s[1] = "Gucci Mane says BRRR";
	s[2] = "C";

	fmt.Println("set:", s);
	fmt.Println("get:", s[1]);

	fmt.Println("len:", len(s));

	s = append(s, "d");
	s = append(s, "e", "f", "rest of the alphabet");

	fmt.Println("appended:", s);

	c := make([]string, len(s));
	copy(c,s);

	fmt.Println("copy:", c);
	
	l := s[2:5]
	fmt.Println("Index 2 inclusive, 5 exclusive: ", l);

	l = s[:5]
	fmt.Println("Up to but not include index 5 ",l);

	l = s[2:]
	fmt.Println("Index 2 inclusive, to the rest of the slice ", l);

	t := []string{"x", "y", "z"};
	fmt.Println("declared new slice: ", t);

	twoD := make([][]int, 3);

	for i := 0; i < 3; i++{

		innerLen := i+1;

		twoD[i] = make([]int, innerLen);
		for j:= 0; j < innerLen; j++{

			twoD[i][j] = i+j;
		}
	}
	fmt.Println("2d example: ", twoD);



}