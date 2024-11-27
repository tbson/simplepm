package testutil

import "flag"

func IsTest() bool {
	return flag.Lookup("test.v") != nil
}
