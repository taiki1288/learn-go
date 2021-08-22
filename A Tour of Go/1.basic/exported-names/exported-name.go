package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}

/*
Goでは、最初の文字が大文字で始まる名前は、
外部のパッケージから参照できるエクスポート（公開）された名前( exported name )です。 
例えば、 Pi は math パッケージでエクスポートされています。
*/