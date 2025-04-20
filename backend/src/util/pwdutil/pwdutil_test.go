package pwdutil

import (
	"testing"
)

func TestMakeThenCheckPwd(t *testing.T) {
	pwd := "password"
	encodedHash := MakePwd(pwd)
	err := CheckPwd(pwd, encodedHash)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
