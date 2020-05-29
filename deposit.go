package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var rate float64
	print("Введите сумму вклада")
	fmt.Fscan(os.Stdin, &sum)
	print("Введите процент")
	fmt.Fscan(os.Stdin, &rate)
	fmt.Printf("Сумма вклада в  1 год: %.2f \n ", sum)
	var sumAfter = sum
	for i := 0; i < 6; i++ {
		sumAfter = sumAfter + (sumAfter*rate)/100
		fmt.Printf("Сумма вклада в ", i, " год: %.2f \n", sumAfter)

	}
}
