package main

import (
	"os"
	"testing"
)

func TestLoadDotEnv(t *testing.T) {
	LoadDotEnv(".testenv")

	if os.Getenv("ABC") != "123" {
		t.Fatalf("123 != %s", os.Getenv("ABC"))
	}

	if os.Getenv("KEY") != "val" {
		t.Fatalf("val != %s", os.Getenv("KEY"))
	}

	if os.Getenv("PORT") != "42069" {
		t.Fatalf("42069 != %s", os.Getenv("PORT"))
	}
}
