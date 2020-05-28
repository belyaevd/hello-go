package main

import (
	"fmt"
	"math"
	"os"
)

func main(){
	var a float64
	var b float64
	var s float64
	var p float64
	var c float64
	print("Введите катет a")
	fmt.Fscan(os.Stdin, &a)
	print("Введите катет b")
	fmt.Fscan(os.Stdin, &b)
    c = math.Sqrt(math.Pow(a , 2) + math.Pow(b, 2))
    s = (a * b)/2
    p = a + b + c
	fmt.Printf("Площадь %.1f \n",s)
	fmt.Printf("Периметр %.1f \n",p)
	fmt.Printf("Гипотенуза %.1f ", c)
}
