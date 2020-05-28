package main

import (
	"fmt"
	"os"
)

func main() {
	const dol float32 = 71
	var rub float32
	var res float32
	print("Введите рубли")
	fmt.Fscan(os.Stdin, &rub)
	res = rub / dol
	fmt.Printf("%.1f", res)
}
