package main

import "fmt"

//	Nível 1 - 2 e 1 - 3
var a int = 42
var b string = "James Bond"
var c bool = true

//	Nível 1 - 4
type ninjaGolang int

//	Nível 1 - 4
var d ninjaGolang

func main() {
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
	//	Use var para declarar três variáveis.
	//	Elas devem ter package-level scope.
	//	Não atribua valores a estas variáveis.
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
	//    1. Demonstre o valor da variável "d"
	//    2. Demonstre o tipo da variável "d"
	//    3. Atribua 42 à variável "d" utilizando o operador "="
	//    4. Demonstre o valor da variável "d"
	// Para os aventureiros: https://golang.org/ref/spec#Types
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	d = 42
	fmt.Println(d)

}
