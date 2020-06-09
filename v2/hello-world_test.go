package go_hello_module

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello World v2.1.1"
	if got := HelloWorld(); got != want {
		t.Errorf("HelloWrold() = %q, want = %q", got, want)
	}
}
