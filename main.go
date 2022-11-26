package main

import "fmt"

var b = "olá e bom dia!"

func main() {
	//
	numerodebytes, erros := fmt.Println("Hello World!", "Oi galera", 500)
	fmt.Println(numerodebytes, erros)

	//
	x := 10
	z := 12.4
	y := "Olá bom dia!"
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("y: %v, %T", y, y)

	x = 20
	x, a := 20, 34
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("a: %v, %T\n", a, a)

	fmt.Println(b)

	c := 44
	d := 65
	fmt.Println(c + d)
	fmt.Println(c == d)
	fmt.Println(c < d)
}
