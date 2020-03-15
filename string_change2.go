package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := "Hello. Have a nice day!"

	fmt.Println(a)

	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&a))
	*(*byte)(unsafe.Pointer(stringHeader.Data + 5)) = '!'

	fmt.Println(a)
}
