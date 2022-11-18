package main

import "fmt"

func main() {
	var a int
	var b int

	//Scan이랑 똑같은데 Printf처럼 서식으로 받는다 (C랑 거의 동일)
	n, err := fmt.Scanf("%d %d\n", &a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
