package main

import (
	"fmt"
)

// Set - our representation of a set data structure
type Set struct {
	Elements map[string]bool
}

func NewSet() *Set {
	return &Set{Elements: make(map[string]bool)}
}

// Add - adds an element to our Set
func (s *Set) Add(elem string) {
	s.Elements[elem] = true
}

// Delete - removes an element from our set if it exists
func (s *Set) Delete(elem string) {
	delete(s.Elements, elem)
}

// Contains - checks to see if an element exists within the set
func (s *Set) Contains(elem string) bool {
	_, exists := s.Elements[elem]
	return exists
}

// List - prints out all of the elements within our set
func (s *Set) List() {
	for k := range s.Elements {
		fmt.Println(k)
	}
}

// Count - counts how many elements are in our set
func (s *Set) Count() {
	fmt.Println(len(s.Elements))
}

// Clear - clears our set
func (s *Set) Clear() {
	s.Elements = make(map[string]bool)
	fmt.Println("Cleared")
}

func main() {

	mySet := NewSet()

	fmt.Println("Sets Elements:")
	mySet.Add("Earth")
	mySet.Add("Venus")
	mySet.Add("Mars")
	mySet.Add("Earth")
	mySet.Add("Moon")
	mySet.Add("Jupiter")
	mySet.Delete("Venus")
	mySet.List()
	mySet.Count()

	fmt.Println(mySet.Contains("Mars"))

	mySet.Clear()

}
