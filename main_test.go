package main

import "testing"

func TestMain(t *testing.T) {
  want := "Hello, world!"
  got := main()
  if got != want {
    t.Errorf("main() = %q, want %q", got, want)
  }
}
