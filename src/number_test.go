package src_test

import (
	"testing"
)

func TestNumber(t *testing.T) {

	// 変数宣言
	var i int = 12345

	// int型からint64型の変換
	var i64 int64 = int64(i)

	if i64 != 12345 {
		t.Fatal("i64 should be 12345")
	}

}
