package src_test

import (
	"testing"
)

func TestType(t *testing.T) {

	// int型からmyInterger型を宣言
	type myInteger int

	var i myInteger = 123

	// int型の同じように演算が可能
	i++

	if i != 124 {
		t.Fatal("i should be 124")
	}

	// 新しい構造体を作成
	type myStruct struct {
		a int
		b int
	}

	myst := myStruct{a: 1, b: 2}

	if myst.a != 1 || myst.b != 2 {
		t.Fatal("error struct initialization")
	}

}
