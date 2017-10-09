package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}
	sum := 0

	for _, num := range nums{

		sum += num;
	}

	fmt.Println("Sum: ", sum);

	for i, num := range nums {

		if num == 3 {
			fmt.Println("Index of 3: ", i);
		}

	}

	kvs := map[string]string{"a":"apple", "b":"banana", "c":"cat"}

	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value);
	}

	for key := range kvs {
		fmt.Println("key: ", key);
	}

	for index, unicode := range "go" {
		fmt.Println(index, unicode);
	}


}