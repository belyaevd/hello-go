package main

import (
	"fmt"
	"os"
)

var a int64

func main() {
	println("Введите число")
	fmt.Fscan(os.Stdin, &a)
	even_number()
	division()
	println("Ряд фибоначчи ")
	fiba()
	mass()
}

func even_number() { //чётно ли число
	if a%2 == 0 {
		println("Число", a, "чётное ")
	} else {
		println("Число", a, "не чётное")
	}
}
func division() { //проверка дения на 3 без остатка
	if a%3 == 0 {
		println("Число", a, "делится на 3 без остатка ")
	} else {
		println("Число", a, "не делится на 3 без остатка")
	}
}
func fiba() {
	var (
		n  = 100
		f1 = 0
		f2 = 1
	)
	fmt.Printf("%d %d ", f1, f2)
	for i := 3; i <= n; i++ {
		fmt.Printf("%d ", f1+f2)
		b := f1
		f1 = f2
		f2 = b + f1
	}
	fmt.Printf("\n")
}

func mass() {
	var numbers [100]int
	var p = 2
	for k := 0; k < 100; k++ {
		numbers[k] = k
		print(numbers[k], " ")
	}
}
