package main

import "fmt"
import "reflect"

func main(){

	m := make(map[string]int);
	//m is a map of string->int key->value pairs

	m["key1"] = 7;
	m["key2"] = 8;

	fmt.Println("current map: ", m);

	value1 := m["key1"];
	fmt.Println("Value for key1 is: ", value1);

	fmt.Println("Current length of map is: ", len(m));

	delete(m, "key2");
	fmt.Println("Length of map after delete is: ", len(m));

	//check if key2 is present
	_, present := m["key2"];
	fmt.Println("Is key2 present: ", present);
	fmt.Println("Type of the  second argument returned from map[] is: ", reflect.TypeOf(present));

	n := map[string]int {"foo": 1, "bar": 2};
	fmt.Println("newmap: ", n)
}