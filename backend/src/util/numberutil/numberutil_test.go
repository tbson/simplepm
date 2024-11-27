package numberutil

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestStrToInt(t *testing.T) {
	if StrToInt("1", 0) != 1 {
		t.Error("StrToInt failed")
	}
	if StrToInt("", 0) != 0 {
		t.Error("StrToInt failed")
	}
	if StrToInt("a", 0) != 0 {
		t.Error("StrToInt failed")
	}
}
