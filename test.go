package main

import (
	"fmt"
	"net/http"
)

type Int int

type Student struct {
	name string
	age  int
}

func (s *Student) updateStudentInfo(name string, age int) {
	s.name = name
	s.age = age
	fmt.Println("In_Func: Name = ", s.name, "age = ", s.age)
	fmt.Println("In_Func: Name_address = ", &s.name, "age_address = ", &s.age)

}

func (n *Int) Inc() {
	*n++
}

type Person struct {
	Name string
	Age  int
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

func strings(str *string) string {
	return *str
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(w)
	w.Write([]byte(r.RequestURI))
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)

	/*
		s := Student{name: "Alice", age: 18}
		fmt.Println("Before update:")
		fmt.Println("Name:", s.name)
		fmt.Println("Age:", s.age)
		fmt.Println("Name_add:", &s.name)
		fmt.Println("Age_add:", &s.age)

		s.updateStudentInfo("Bob", 20)

		fmt.Println("After update:")
		fmt.Println("Name:", s.name)
		fmt.Println("Age:", s.age)
		fmt.Println("Name_add:", &s.name)
		fmt.Println("Age_add:", &s.age)
	*/
	/*
		user := Person{"yuuta"}
		user.setName("yuuki")
		fmt.Println(user.Name)

		user2 := Person{"yuta"}
		user2.setName2("yuuki")
		fmt.Println(user2.Name)
	*/

	/*
		str1 := "yuuki"
		str := &str1

		str3 := "yuuta"
		fmt.Println(str3)

		fmt.Println(reflect.TypeOf(strings(str)))
		fmt.Println((strings(str)))
	*/

	/*

		var n Int
		fmt.Println(n)
		n.Inc()
		fmt.Println(n)

		a := 2
		b := 3
		fmt.Println(call.Sum(a, b))
		fmt.Println(call.Multi(a, b))

		p1 := mkpkg.Person{Name: "yuuki", Age: 20}
		p2 := mkpkg.Person{Name: "asuna", Age: 18}

		fmt.Println(p1.Name, "は", p1.Age, "歳です。")
		fmt.Println(p2.Name, "は", p2.Age, "歳です。")
		fmt.Println("二人の合計年齢は", mkpkg.Add(p1.Age, p2.Age), "です。")
	*/

	/*
		user := Person{"yutmpta"}
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
