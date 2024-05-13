package util

import (
	"testing"
)

func CheckFatal(err error) {
	t := &testing.T{}
	if err != nil {
		t.Fatal(err)
	}
}
