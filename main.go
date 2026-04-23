package main

import "fmt"

func main() {
	a := 10

	p := &a
	pp := &p

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(**pp)
	fmt.Println(pp)
	fmt.Printf("%T\n", pp)
}
