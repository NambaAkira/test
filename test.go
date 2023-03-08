package main

import (
	"fmt"
)

type Int int

func (n *Int) Inc() {
	*n++
}

type Person struct {
	Name string
}

func (p Person) Greeting(msg string) {
	fmt.Println("name(関数内)", &(p.Name))
}

func (p *Person) Greeting2(msg string) {
	fmt.Println("name(関数内)", &(p.Name))
}

func (p Person) setName(name string) {
	p.Name = name
	fmt.Println("関数内1:", p.Name)
}

func (p *Person) setName2(name string) {
	p.Name = name
	fmt.Println("関数内2:", p.Name)

}

func main() {

	var n Int
	fmt.Println(n)
	n.Inc()
	fmt.Println(n)

	a := 2
	b := 3
	fmt.Println(sum.sum(a, b))

	/*
		user := Person{"yuta"}
		user.setName("yuuki")
		fmt.Println(user.Name)

		user2 := Person{"yuta"}
		user2.setName2("yuuki")
		fmt.Println(user2.Name)

		println("-----------")

		str := "hello world"
		fmt.Print(str, "\n")
		fmt.Print("おはようございます。\n")
	*/
}
