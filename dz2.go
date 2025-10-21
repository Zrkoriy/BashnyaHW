package main

import "fmt"

func main() {
	var n int
	fmt.Print("Ввод: целое число меньше 12307: ")
	fmt.Scan(&n)

	if n >= 12307 {
		fmt.Println("число больше или равно 12307")
	} else {
		for n < 12307 {
			if n < 0 {
				n = n * -1
			} else if n%7 == 0 {
				n = n * 39
			} else if n%9 == 0 {
				n = n*13 + 1
			} else {
				n = (n + 2) * 3
			}
			if n%13 == 0 && n%9 == 0 {
				fmt.Println("service error")
				break
			} else {
				n = n + 1
			}
		}
	}
	fmt.Printf("Результат: %v\n", n)
}
