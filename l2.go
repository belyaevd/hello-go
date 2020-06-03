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

func mass() { //не получилось )))
	var p int
	var n = 100

	mass := make([]int, n)
	var flag bool //"зачеркивали" ли число для данного p

	n = n - 1 //теперь n - это количесвтво чисел в массиве
	for i := 0; i < n; i++ {
		mass[i] = i + 2
		for i := 0; i < n; i++ {
			p = mass[i]
			flag = false
			for j := i + 1; j < n; j++ {
				if mass[j]%p != 0 {
					for k := j; k < n-1; k++ {
						mass[k] = mass[k+1]
					}
					flag = true
					n-- //уменьшаем, потому что чисел на одно стало меньше
					j-- //уменьшаем, для того чтобы снова проверить на кратность j-е число. Оно же теперь стало другим
				}
			}
			if flag == false {
				break
			}
		}
		for i := 0; i < n; i++ {

			fmt.Printf("%d ", mass[i])
		}
	}
}
