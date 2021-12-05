package main

import "testing"

//
// go run test -v
//
func TestShowHello(t *testing.T) {
	got := ShowHello("mahdi")
	expected := "Hello mahdi"
	if got != expected {
		t.Errorf("[ShowHello] The result `%s` is not equal to expected `%s` ", got, expected)
	}
}

//
// go test -bench=.
//
func BenchmarkShowHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShowHello("mahdi")
	}
}
