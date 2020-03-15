package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func main() {
	a := "Hello. Current time is " + time.Now().String()

	fmt.Println(a)

	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&a))
	*(*byte)(unsafe.Pointer(stringHeader.Data + 5)) = '!'

	fmt.Println(a)
}
