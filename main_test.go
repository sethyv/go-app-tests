// main_test.go
package main

import "testing"

func TestMain(t *testing.T) {
	want := "Hello, world!"
	got := runMain()
	if got != want {
		t.Errorf("runMain() = %q, want %q", got, want)
	}
}

func runMain() string {
	return "Hello, world!"
}

func TestAdd(t *testing.T) {
	expected := 5
	result := Add(2, 3)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
