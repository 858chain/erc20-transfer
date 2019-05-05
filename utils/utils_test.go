package utils

import (
	"testing"
)

func TestSplitLast(t *testing.T) {
	str := "foo-bar-bz"
	sep := "-"

	if LastSplit(str, sep) != "bz" {
		t.Errorf("LastSplit should be bz")
	}

	str = "foo-bar-bz"
	sep = ""
	if LastSplit(str, sep) != "" {
		t.Errorf("LastSplit should be ''")
	}

	str = ""
	sep = "foo - bar"
	if LastSplit(str, sep) != "" {
		t.Errorf("LastSplit should be ''")
	}

	str = "foobar"
	sep = "-"
	if LastSplit(str, sep) != "foobar" {
		t.Errorf("LastSplit should be fooabr")
	}
}
