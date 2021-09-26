package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64;
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tempo do i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("Tipo do i2", reflect.TypeOf(i2))
	fmt.Println("O valor de i2 é", i2)

	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo de 49.99 é", reflect.TypeOf(49.99))

	bo := true
	fmt.Println("O tipo da variavel bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Olá meu nome é Victor"
	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1))
	fmt.Println(s1, "!")
	fmt.Println("O tamanho da string é", len(s1))

	s2 := `Olá
	meu
	nome
	é
	Víctor`
	fmt.Println("O tamanho da string é", len(s2))

	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}