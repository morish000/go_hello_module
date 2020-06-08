package go-hello-module

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello World"
	if got := HelloWorld(); got != want {
		t.Error("HelloWrold() = %q, want = %q", got, want)
	}
}