package main

import (
	"fmt"

	"github.com/evandrojr/string-interpolation/esi"
)

func main() {
	esi.Print("Print ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	fmt.Print(" no line bread")
	fmt.Println()
	esi.Println("Println ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	f := esi.Sprint("Sprint ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	fmt.Print(f)
}
