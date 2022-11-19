package main

import (
	"github.com/evandrojr/string-interpolation/esi"
)

func main() {
	esi.Print("Print ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	esi.Print(" no line bread")
	esi.Println()
	esi.Println("Println ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	f := esi.Sprint("Sprint ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	esi.Print(f)
}
