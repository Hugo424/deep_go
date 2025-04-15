package main

import (
	"fmt"
	"unsafe"
)

func main() {
	ToLittleEndian(0x0000FFFF)
}

func ToLittleEndian(number uint32) uint32 {
	fmt.Printf("Исходное число: %08X\n", number)
	size := unsafe.Sizeof(number)
	var result uint32
	pointer := unsafe.Pointer(&result)

	for i := range size {
		sdvig := number >> ((size - 1 - i) * 8)
		bytePtr := (*uint8)(unsafe.Pointer(&sdvig))
		ptr := (*uint8)(unsafe.Add(pointer, i))
		*ptr = *bytePtr
	}

	fmt.Printf("Конечное число: %08X\n", result)

	return number
}
