package main

import (
	"fmt"
	"unsafe"
)

type Bytes400 struct {
	val [100]int32
}

type TestStruct struct {
	a [9]int64
	b byte
	c *Bytes400
	d int64
}

func main() {
	array := [10]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var sum int8

	// Unsafe array iteration
	sizeOfUint64 := unsafe.Sizeof(array[0])
	for i := uintptr(0); i < 10; i++ {
		sum += *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&array)) + sizeOfUint64*i))
	}
	fmt.Println(sum)

	// Size of struct and offsets of struct fields
	t := TestStruct{b: 42}
	fmt.Println(unsafe.Sizeof(t))
	fmt.Println(unsafe.Offsetof(t.a), unsafe.Offsetof(t.b), unsafe.Offsetof(t.c), unsafe.Offsetof(t.d))

	fmt.Println(unsafe.Sizeof(Bytes400{}))

	// Change struct field t.b value
	*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&t)) + unsafe.Offsetof(t.b)))++
	fmt.Println(t.b)
}
