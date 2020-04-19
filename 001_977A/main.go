package main

import (
	"fmt"
)

func tanyaSub(n, s int) int {
	for i := 0; i < s; i++ {
		if n%10 == 0 {
			n = n / 10
		} else {
			n = n - 1
		}
	}

	return n
}

func main() {
	var n, s int
	_, err := fmt.Scanf("%d %d", &n, &s)

	fmt.Println(tanyaSub(n, s))

	if err != nil {
		fmt.Print(err)
	}
}
