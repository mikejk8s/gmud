package main

import "fmt"

type person struct {
    name string
    age int
    class string
    }

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

    fmt.Println(person{"Bob", 42, "software engineer"})

    fmt.Println(person{"Alice", 32, "software engineer"})

    fmt.Println(person{"Carl", 33, "software engineer"})

    fmt.Println(person{"John", 29, "accountant"})

    fmt.Println(person{"Rick", 34, "manager"})

    s := person{name: "Bob", age: 50, class: "Manager"}
    fmt.Println(s.name, s.age, s.class)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
}