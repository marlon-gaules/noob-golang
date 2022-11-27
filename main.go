package main

import (
	"fmt"
)

//Cap 2 - Aula 3
var b = "olá e bom dia!"

//Cap 2 - Aula 5
var f int = 15
var g = 13.8
var h string

//Cap 2 - Aula 6
var i int
var j float64
var k string
var l bool

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

	//Cap 2 - Aula 6
	i = 1024
	fmt.Printf("%v, %T\n", i, i)
	i = 2048
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", l, l)

	//Cap 2 - Aula 7
	m := "Oi bebe\ncomo vai você?\tespero \"que\" tudo bem!"
	fmt.Println(m)

	//Cap 2 - Aula 7
	n := `oidomsmskcs
	
	wwndkwndknwdknwdknwkdnwkdnwkd
	dknwkndkwndkndknwkw
	
	`
	fmt.Println(n)

	//Cap 2 - Aula 7
	o := "Aqui e está chovendo"
	p := "Aqui está relampiano"
	q := fmt.Sprint(o, " ", p)
	fmt.Println(q)
}

//Cap 2 - Aula 4
func qualquercoisa(x int) {
	fmt.Println(b)
	fmt.Println(x)
}
