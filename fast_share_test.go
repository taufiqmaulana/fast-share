package main

import (
	"testing"
)

func Test_GenerateRandomString(t *testing.T) {
	a := GenerateRandomString(5)
	b := GenerateRandomString(5)
	if len(a) != 5 {
		t.Error("Generated string should be 5")
	}

	if len(a) != len(b) {
		t.Error("Generated string should have same length for each iteration")
	}

	if a == b {
		t.Error("Generated string should have different content for each iteration")
	}
}

func Test_GeGetLocalIP(t *testing.T) {
	ip := GetLocalIP()
	if ip == nil {
		t.Error("IP should not be nil")
	}
}

func Test_PrintQr(t *testing.T) {
	PrintQr("test")
}
