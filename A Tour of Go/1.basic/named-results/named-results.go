package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

/*Goでは戻り値となる変数に名前をつけることができます。
戻り値に名前をつけると、関数の最初で定義した変数名として扱われます。

この戻り値の名前は、戻り値の意味を示す名前とすることで、
関数のドキュメントとして表現するようにしましょう。

名前をつけた戻り値の変数を使うと、 return ステートメントに何も書かずに戻すことができます。
これを "naked" return と呼びます。
*/

