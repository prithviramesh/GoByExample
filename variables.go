package main

import (
	
	"fmt"
	"reflect"	
)
func main() {

	var a string = "initial";
	fmt.Println(a);

	f := false;
	fmt.Println(reflect.TypeOf(f));

}