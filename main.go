package main

import (
	"fmt"
	"runtime"
)

//########################################################
//###################  CAP. 2  ###########################
//########################################################
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

//Cap 2 - Aula 8
type hotdog int

//Cap 2 - Aula 8
var r hotdog

//Cap 2 - Aula 9
var t hotdog = 101

//########################################################
//###################  CAP. 4  ###########################
//########################################################
//Cap 4 - Aula 1
var aa bool

//Cap 4 - Aula 7
const pp = 10

//Cap 4 - Aula 7
const (
	ss = 10
	tt = 110
	uu = "Oi, seja bem vindo!"
)

//Cap 4 - Aula 7
var qq int
var rr float64
var vv string

//Cap 4 - Aula 8
const (
	ww = iota
	_  = iota
	xx = iota
	_  = iota
	yy = iota
)

//Cap 4 - Aula 8
const (
	zz  = iota
	aaa = iota
	bbb = iota * 5
)

//Cap 4 - Aula 8
const (
	ccc = iota + 100
	ddd
	eee
	fff
)

func main() {

	//########################################################
	//###################  CAP. 2  ###########################
	//########################################################

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
	o := "Aqui está chovendo"
	p := "Aqui está relampiano"
	q := fmt.Sprint(o, " ", p)
	fmt.Println(q)

	//Cap 2 - Aula 8
	s := 10
	fmt.Printf("%T\n", r)
	fmt.Printf("%T\n", s)

	//Cap 2 - Aula 9
	s = int(t)
	fmt.Printf("%v\n", s)

	//########################################################
	//###################  CAP. 4  ###########################
	//########################################################
	//Cap 4 - Aula 1
	fmt.Println(aa)
	bb := true
	fmt.Println(bb)
	bb = 10 < 100
	cc := 10 == 101
	dd := 10 > 100
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)

	//Cap 4 - Aula 3
	ee := "e"
	ff := "é"
	gg := "u9999"
	fmt.Printf("%v, %v, %v\n", ee, ff, gg)
	hh := []byte(ee)
	ii := []byte(ff)
	jj := []byte(gg)
	fmt.Printf("%v, %v, %v\n", hh, ii, jj)

	//Cap 4 - Aula 3
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	//Cap 4 - Aula 4
	var kk uint16
	kk = 65535
	fmt.Println(kk)
	kk++
	fmt.Println(kk)
	kk++
	fmt.Println(kk)

	//Cap 4 - Aula 5
	ll := "Hello, Playground"
	mm := []byte(ll)
	fmt.Println(ll)
	fmt.Printf("%v, %T\n", ll, ll)
	fmt.Printf("%v, %T\n", mm, mm)

	//Cap 4 - Aula 5
	for _, nn := range mm {
		fmt.Printf("%v - %T - %#U - %#x\n", nn, nn, nn, nn)
	}

	//Cap 4 - Aula 6
	oo := 100
	fmt.Printf("%d, %b, %#x\n", oo, oo, oo)

	//Cap 4 - Aula 7
	qq = pp
	fmt.Printf("%v, %T\n", qq, qq)

	//Cap 4 - Aula 7
	rr = pp
	fmt.Printf("%v, %T\n", rr, rr)

	//Cap 4 - Aula 7
	vv = uu
	fmt.Printf("%v, %T\n", vv, vv)

	//Cap 4 - Aula 8
	fmt.Println(ww, xx, yy)
	fmt.Println(zz, aaa, bbb)
	fmt.Println(ccc, ddd, eee, fff)

	//Cap 4 - Aula 9
	ggg := 1
	hhh := ggg << 2
	iii := ggg >> 1
	fmt.Printf("%b\n", ggg)
	fmt.Printf("%b\n", hhh)
	fmt.Printf("%b\n", iii)
}

//########################################################
//###################  CAP. 2  ###########################
//########################################################
//Cap 2 - Aula 4
func qualquercoisa(x int) {
	fmt.Println(b)
	fmt.Println(x)
}
