package fundemental

import (
	"fmt"
	"strconv"
)

// Generic helps generic types to be parameterized
type Generic interface {
}

type sampleStruct struct {
	a int
	b int
	c string
}

// TestPointers tests pointer operations in Go
func TestPointers() {
	a := "1"
	b := 456789123456
	c := "789"
	d := 123
	pa := &a
	pb := &b
	pc := &c

	fmt.Printf("a locates at %p\n", pa)
	fmt.Printf("b locates at %p\n", pb)
	fmt.Printf("c locates at %p\n", pc)
	fmt.Printf("d locates at %p\n", &d)
	fmt.Printf("b is %d bytes away from a\n", getPointerDistance(pa, pb))
}

// PrintMemoryLayout prints memory layout for an object
func PrintMemoryLayout() {

}

func getIntFromPointer(ptr Generic) int {
	fmt.Printf("incoming %d\n", ptr)
	r, _ := strconv.Atoi(fmt.Sprintf("%d", ptr))
	fmt.Printf("got result %d\n", r)
	return r
}

func getPointerDistance(ptrA Generic, ptrB Generic) int {
	return Abs(getIntFromPointer(ptrA) - getIntFromPointer(ptrB))
}
