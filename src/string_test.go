package src_test

import (
	"testing"
)

func TestString(t *testing.T) {

	var str string

	str = "あ"

	str = str + "い"

	str += "う"

	if str != "あいう" {
		t.Fatal("str should be あいう")
	}

}
