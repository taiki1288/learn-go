package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

/*
この例では、 add 関数は、 int 型の２つのパラメータを取ります。
変数名の 後ろ に型名を書くことに注意してください。
*/