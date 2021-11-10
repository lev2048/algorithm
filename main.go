package main

import "fmt"

type A struct {
	v int
}

func main() {
	x := &A{4}
	setNil(&x)
	fmt.Println(x)
}

func setNil(x **A) {
	z := *x
	*x.v = 3
	*x = nil
}
