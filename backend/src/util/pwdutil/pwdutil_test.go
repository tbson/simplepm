package pwdutil

import (
	"testing"
)

func TestMakeThenCheckPwd(t *testing.T) {
	pwd := "password"
	encodedHash := MakePwd(pwd)
	result, err := CheckPwd(pwd, encodedHash)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !result {
		t.Errorf("Expected true, got false")
	}
}
