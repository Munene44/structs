package main

import "fmt"
 
type person struct {
	name string
	age int
}
 func NewPerson(name string) *person {
 	p := person{name: name}
 	p.age = 42
 	return &p
 }

 fun main() {
 	fmt.Println(person{"Kim", 30})
 	fmt.Println(person{"john", 40})
 	f.mt.Println(&person{"lin", age 40})
 }