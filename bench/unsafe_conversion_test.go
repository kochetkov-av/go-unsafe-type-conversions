package main

import (
	"bytes"
	"testing"
)

func TestUnsafeBytesToString(t *testing.T) {
	bs := []byte("Test")
	str := UnsafeBytesToString(bs)
	if str != "Test" {
		t.Fail()
	}

	// Test string mutation
	bs[0] = 't'
	if str != "test" {
		t.Fail()
	}
}

func TestUnsafeStringToBytes(t *testing.T) {
	if bytes.Compare(UnsafeStringToBytes("Test"), []byte("Test")) != 0 {
		t.Fail()
	}
}

func BenchmarkSafeBytesToString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SafeBytesToString([]byte{'T', 'e', 's', 't'})
	}
}

func BenchmarkSafeStringToBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SafeStringToBytes("Test")
	}
}

func BenchmarkUnsafeBytesToString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UnsafeBytesToString([]byte{'T', 'e', 's', 't'})
	}
}

func BenchmarkUnsafeStringToBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UnsafeStringToBytes("Test")
	}
}
