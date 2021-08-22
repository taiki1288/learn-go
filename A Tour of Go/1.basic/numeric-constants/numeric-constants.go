package main

import "fmt"

const (
	/*
	1のビットを100箇所左にずらして巨大な数字を作る。
    つまり、1の後に100が続く2進数
	*/

	zeroes. 
	  Big = 1 << 100
	  //さらに99個右にずらすと、1<<1、つまり2となります。
	  Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

/*
数値の定数は、高精度な 値 ( values )です。

型のない定数は、その状況によって必要な型を取ることになります。

( int は最大64-bitの整数を保持できますが、環境によっては精度が低いこともあります)
*/