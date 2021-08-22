package main 

import(
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}

/*
型変換

変数 v 、型 T があった場合、 T(v) は、変数 v を T 型へ変換します。

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
よりシンプルに記述できます:

i := 42
f := float64(i)
u := uint(f)

C言語とは異なり、Goでの型変換は明示的な変換が必要です

*/