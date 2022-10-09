// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func multiply(a, b int) int {
// 	return a * b
// }

// // type을 변수랑 리턴 앞에 꼭 명시해야한다

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// // 여러 return 값을 가질 수 있다

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

// //파라미터 갯수를  모를 때 ...string으로 받을 수 있다

// func main() {
// 	// fmt.Println(multiply(2, 2))
// 	// totalLength, upperName := lenAndUpper("woojin")
// 	// fmt.Println(totalLength, upperName)
// 	repeatMe("nico", "woojin", "dal", "jack")
// }

// // Println 에서 P가 대문자인 이유: export 가능 . 소문자 -> export 안됨
// // Go 에서는 변수를 선언하고 사용하지 않으면 오류가 난다
// // 리턴값을 무시할때는 _ 를 이용한다 a , _ :=

// //makechange
