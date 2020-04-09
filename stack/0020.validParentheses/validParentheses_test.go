/*************************************************************************
+Author   : ChenhuiJia
+Data     : 2020-04-09
+************************************************************************/

package main

import "testing"

func TestIsValid(t *testing.T) {
	s1 := "{[()]}"
	s2 := "{]"
	s3 := "()"
	s4 := "{}[]("
	s5 := "(){}[]"

	if !IsValid(s1) {
		t.Error("s1 is valid")
	}

	if IsValid(s2) {
		t.Error("s2 is invalid")
	}

	if !IsValid(s3) {
		t.Error("s3 is valid")
	}

	if IsValid(s4) {
		t.Error("s4 is invalid")
	}

	if !IsValid(s5) {
		t.Error("s5 is valid")
	}
}
