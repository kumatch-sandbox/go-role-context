package main

import (
	"fmt"
)

type Person struct {
	name string
}

func (p *Person) hello() string {
	return fmt.Sprintf("hello, I'm %s", p.name)
}

type CatRole struct {
	*Person
}

func (cat *CatRole) say() string {
	return "meow, meow"
}

func (cat *CatRole) hello() string {
	return cat.say()
}

func catContext(p *Person) {
	cat := &CatRole{p}

	fmt.Println(cat.hello())
}

func main() {
	p := &Person{name: "mike"}
	fmt.Println(p.hello())

	catContext(p)
}
