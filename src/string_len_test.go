package src_test

import (
	"testing"
	"unicode/utf8"
)

func TestStringLen(t *testing.T) {

	var en string = "golang"

	var ja string = "Go言語"

	// lenは文字列のバイト数を返すので文字列の長さを返すわけではない
	enLen := len(en)
	jaLen := len(ja)

	if enLen != 6 {
		t.Fatal("enLen should be 6")
	}

	if jaLen != 8 {
		t.Fatal("jaLen should be 8")
	}

}

func TestStringCount(t *testing.T) {

	var ja string = "Go言語"

	// 文字列の長さを取得
	got := utf8.RuneCountInString(ja)

	want := 4

	if got != want {
		t.Fatal("got should be 6")
	}

}
