package main

import "fmt"

type person struct {
    name string
    age  int
}

func NewPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Kelvin", age: 30})

    fmt.Println(person{name: "Kendi"})

    fmt.Println(&person{name: "Ann", age: 40})

    fmt.Println(NewPerson("Alex"))

    s := person{name: "Munene", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
}