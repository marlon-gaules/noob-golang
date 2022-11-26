package main

import "fmt"

var b = "olá e bom dia!" //Cap 2 - Aula 3

//Cap 2 - Aula 5
var f int = 15
var g = 13.8
var h string

func main() {
	//Cap 2 - Aula 1 e 2
	numerodebytes, erros := fmt.Println("Hello World!", "Oi galera", 500)
	fmt.Println(numerodebytes, erros)

	//Cap 2 - Aula 3
	x := 10
	z := 12.4
	y := "Olá bom dia!"
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("y: %v, %T", y, y)

	//Cap 2 - Aula 3
	x = 20
	x, a := 20, 34
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("a: %v, %T\n", a, a)

	//Cap 2 - Aula 3
	fmt.Println(b)

	//Cap 2 - Aula 3
	c := 44
	d := 65
	fmt.Println(c + d)
	fmt.Println(c == d)
	fmt.Println(c < d)

	//Cap 2 - Aula 4
	e := 99
	qualquercoisa(e)

	//Cap 2 - Aula 5
	f = 168
	g = 188.79
	h = "Beleza?"
	fmt.Println(f, g, h)
}

//Cap 2 - Aula 4
func qualquercoisa(x int) {
	fmt.Println(b)
	fmt.Println(x)
}
