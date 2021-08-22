package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main(){
	fmt.Println(add(42, 13))
}
 
/*
x int, y intを
x, y intへ省略できます。
*/