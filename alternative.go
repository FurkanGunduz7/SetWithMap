package main

import "fmt"

func main() {

	s := map[int]bool{5: true, 2: true}
	_, ok := s[6] // check for existence
	fmt.Println(ok)
	s[8] = true // add element
	fmt.Println(s)
	delete(s, 2) // remove element
	fmt.Println(s)
	for k, _ := range s {
		fmt.Println(k)
	}
}
