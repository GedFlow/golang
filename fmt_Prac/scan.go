package main

import "fmt"

func main() {
	var a int
	var b int

	// 특이하게도 goloang은 scan으로 두 개의 값을 반환한다. n에는 입력받은 값의 갯수, err에는 에러체크 값
	n, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
