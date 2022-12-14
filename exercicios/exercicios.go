package main

import "fmt"

//########################################################
//################# CAP. 3 - Exercícios ##################
//########################################################
//	Nível 1 - 2 e 1 - 3
var a int = 42
var b string = "James Bond"
var c bool = true

//	Nível 1 - 4
type ninjaGolang int

//	Nível 1 - 4
var d ninjaGolang

//	Nível 1 - 5
var e int

//########################################################
//################# CAP. 5 - Exercícios ##################
//########################################################
//	Nível 2 - 3
const (
	m string  = "Hello world"     //tipada
	n float64 = 6.54              //tipada
	o         = "Bom dia galera!" //não tipada
	p         = 5                 //não tipada
)

func main() {
	//########################################################
	//################# CAP. 3 - Exercícios ##################
	//########################################################
	//	Cap. 3 – Exercícios
	//	Nível 1 - 1
	//	Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
	//		1. 42
	//    	2. "James Bond"
	//    	3. true
	//	Agora demonstre os valores nestas variáveis, com:
	//    	1. Uma única declaração print
	//    	2. Múltiplas declarações print
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//	Nível 1 - 2
	//	Use var para declarar três variáveis
	//	Elas devem ter package-level scope
	//	Não atribua valores a estas variáveis
	//	Utilize os seguintes identificadores e tipos para estas variáveis:
	//		1. Identificador "a" deverá ter tipo int
	//		2. Identificador "b" deverá ter tipo string
	//		3. Identificador "c" deverá ter tipo bool
	//	Na função main:
	//		1. Demonstre os valores de cada identificador
	//		2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	//	Nível 1 - 3
	//	Utilizando a solução do exercício anterior:
	//  Em package-level scope, atribua os seguintes valores às variáveis:
	//  	1. para "a" atribua 42
	//      2. para "b" atribua "James Bond"
	//      3. para "c" atribua true
	//  Na função main:
	//      1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável.
	//		Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
	//      2. Demonstre a variável "s".
	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)

	//	Nível 1 - 4
	//  Crie um tipo. O tipo subjacente deve ser int.
	// 	Crie uma variável para este tipo, com o identificador "d", utilizando a palavra-chave var.
	//	Na função main:
	//  	1. Demonstre o valor da variável "d"
	//   	2. Demonstre o tipo da variável "d"
	//   	3. Atribua 42 à variável "d" utilizando o operador "="
	//   	4. Demonstre o valor da variável "d"
	// Para os aventureiros: https://golang.org/ref/spec#Types
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	d = 42
	fmt.Println(d)

	//	Nível 1 - 5
	//	Utilizando a solução do exercício anterior:
	//  Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "e".
	//	O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
	//  Na função main:
	//  	1. Isto já deve estar feito:
	//    		1. Demonstre o valor da variável "d"
	//          2. Demonstre o tipo da variável "d"
	//          3. Atribua 42 à variável "d" utilizando o operador "="
	//          4. Demonstre o valor da variável "d"
	//      2. Agora faça tambem:
	//          1. Utilize conversão para transformar o tipo do valor da variável "d" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "d" a "e"
	//         	2. Demonstre o valor de "e"
	//          3. Demonstre o tipo de "e"
	e = int(d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)

	//########################################################
	//################# CAP. 5 - Exercícios ##################
	//########################################################
	//	Nível 2 - 1
	//	Escreva um programa que mostre um número em decimal, binário, e hexadecimal
	f := 101
	fmt.Printf("%d, %b, %#x\n", f, f, f)

	//	Nível 2 - 2
	//	Escreva expressões utilizando os operadores, e atribua seus valores a variáveis
	//	==
	//	=/
	//	<=
	//	<
	//	>=
	//	>
	//	Demonstre estes valores
	g := 5 == 10
	h := 21 != 21
	i := 3 <= 5
	j := 100 < 99
	k := 55 >= 55
	l := 65 > 63
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", g, h, i, j, k, l)

	//	Nível 2 - 3
	//	Crie constantes tipadas e não-tipadas
	//	Demonstre seus valores
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", n)
	fmt.Printf("%v\n", o)
	fmt.Printf("%v\n", p)

	//	Nível 2 - 4
	//	Crie um programa que:
    //	Atribua um valor int a uma variável
    //	Demonstre este valor em decimal, binário e hexadecimal
    //	Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    //	Demonstre esta outra variável em decimal, binário e hexadecimal
	m := 200
	fmt.Printf("%d, %b, %#x\n", m, m, m)
	n := m << 1
	fmt.Printf("%d, %b, %#x\n", n, n, n)
}
