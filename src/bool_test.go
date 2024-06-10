package src_test

import (
	"testing"
)

func TestBool(t *testing.T) {

	// 変数宣言
	var b bool

	// 代入
	b = true
	b = false

	b = true || false

	if b != true {
		t.Fatal("b should be true")
	}

}
