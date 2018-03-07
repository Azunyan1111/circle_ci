package main

import "testing"

func TestHello2(t *testing.T) {
	s := Hello()
	if s != "Hello World"{
		t.Fatal("Error: Hello func test code. return miss")
	}
}

func TestHello3(t *testing.T) {
	s := Hello()
	if s != "Hello World2"{
		t.Fatal("Error: Hello func test code. return miss")
	}
}
